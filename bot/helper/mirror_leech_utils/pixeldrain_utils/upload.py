from io import BufferedReader
from logging import getLogger
from os import path as ospath
from os import walk as oswalk
from pathlib import Path

from aiofiles.os import path as aiopath
from aiohttp import ClientSession
from tenacity import (
    RetryError,
)

from bot.core.config_manager import Config
from bot.helper.ext_utils.bot_utils import SetInterval

LOGGER = getLogger(__name__)


class ProgressFileReader(BufferedReader):
    def __init__(self, filename, read_callback=None):
        super().__init__(open(filename, "rb"))
        self.__read_callback = read_callback
        self.length = Path(filename).stat().st_size

    def read(self, size=None):
        size = size or (self.length - self.tell())
        if self.__read_callback:
            self.__read_callback(self.tell())
        return super().read(size)


class PixeldrainUpload:
    def __init__(self, listener, path):
        self.listener = listener
        self._path = path
        self._updater = None
        self._is_errored = False
        self.total_files = 0
        self.total_folders = 0
        self.is_uploading = False
        self.total_bytes = 0
        self._processed_bytes = 0

        self.api_url = "https://pixeldrain.com/api"
        self.token = user_dict.get("PIXELDRAIN_API_KEY") or Config.PIXELDRAIN_API

        self.total_bytes = (
            sum(f.stat().st_size for f in Path(self._path).rglob("*") if f.is_file())
            if ospath.isdir(self._path)
            else ospath.getsize(self._path)
        )

        self.update_interval = 2

    @property
    def speed(self):
        if self._updater:
            return self._updater.speed
        return 0

    @property
    def processed_bytes(self):
        return self._processed_bytes

    def __progress_callback(self, current):
        self._processed_bytes += current
        if self._updater:
            self._updater.update(current)

    async def progress(self):
        pass

    async def upload_aiohttp(self, url, file_path):
        import base64

        headers = {}
        if self.token:
            auth = base64.b64encode(f":{self.token}".encode()).decode()
            headers["Authorization"] = f"Basic {auth}"

        async with ClientSession() as session:
            file_name = ospath.basename(file_path)
            with ProgressFileReader(file_path, self.__progress_callback) as f:
                async with session.put(
                    f"{self.api_url}/file/{file_name}", data=f, headers=headers
                ) as resp:
                    return await resp.json()

    async def upload_file(self, file_path):
        self.is_uploading = True
        try:
            res = await self.upload_aiohttp(self.api_url, file_path)
            if res and res.get("success"):
                return res.get("id")
            raise Exception(res.get("message", "Upload failed"))
        except Exception as err:
            LOGGER.error(f"Pixeldrain Upload failed for {file_path}: {err}")
            return None
        finally:
            self.is_uploading = False

    async def _upload_dir(self, input_directory):
        import base64

        uploaded_ids = []
        for root, _dirs, files in oswalk(input_directory):
            for file in files:
                if self.listener.is_cancelled:
                    return None
                file_path = ospath.join(root, file)
                fid = await self.upload_file(file_path)
                if fid:
                    uploaded_ids.append(fid)
                    self.total_files += 1

        if not uploaded_ids:
            return None

        # Create a list/album in Pixeldrain if multiple files
        if len(uploaded_ids) > 1:
            headers = {}
            if self.token:
                auth = base64.b64encode(f":{self.token}".encode()).decode()
                headers["Authorization"] = f"Basic {auth}"
            data = {
                "title": ospath.basename(input_directory),
                "files": [{"id": i} for i in uploaded_ids],
            }
            async with (
                ClientSession() as session,
                session.post(
                    f"{self.api_url}/list", json=data, headers=headers
                ) as resp,
            ):
                res = await resp.json()
                if res and res.get("success"):
                    return f"l/{res.get('id')}"
        elif len(uploaded_ids) == 1:
            return f"u/{uploaded_ids[0]}"
        return None

    async def upload(self):
        try:
            LOGGER.info(f"Pixeldrain Uploading: {self._path}")
            self._updater = SetInterval(self.update_interval, self.progress)

            # Run the async process directly
            await self._upload_process()

        except Exception as err:
            if isinstance(err, RetryError):
                LOGGER.info(f"Total Attempts: {err.last_attempt.attempt_number}")
                err = err.last_attempt.exception()
            err = str(err).replace(">", "").replace("<", "")
            LOGGER.error(err)
            await self.listener.on_upload_error(err)
            self._is_errored = True
        finally:
            if self._updater:
                self._updater.cancel()

    async def _upload_process(self):
        if await aiopath.isfile(self._path):
            file_id = await self.upload_file(self._path)
            if file_id:
                link = f"https://pixeldrain.com/u/{file_id}"
                mime_type = "File"
                self.total_files = 1
            else:
                raise ValueError("Failed to upload file to Pixeldrain")
        elif await aiopath.isdir(self._path):
            list_id = await self._upload_dir(self._path)
            if list_id:
                link = f"https://pixeldrain.com/{list_id}"
                mime_type = "Folder"
            else:
                raise ValueError("Failed to upload folder to Pixeldrain")
        else:
            raise ValueError("Invalid file path!")

        if self.listener.is_cancelled:
            return

        LOGGER.info(f"Uploaded To Pixeldrain: {self.listener.name}")
        await self.listener.on_upload_complete(
            link,
            self.total_files,
            self.total_folders,
            mime_type,
            dir_id="",
        )

    async def cancel_task(self):
        self.listener.is_cancelled = True
        if self.is_uploading:
            LOGGER.info(f"Cancelling Pixeldrain Upload: {self.listener.name}")
            await self.listener.on_upload_error(
                "Pixeldrain upload has been cancelled!"
            )
