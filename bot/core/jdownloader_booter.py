from bot import LOGGER
from bot.helper.ext_utils.bot_utils import new_task
from myjd import MyJdApi


class JDownloader(MyJdApi):
    def __init__(self):
        super().__init__()
        self._username = ""
        self._password = ""
        self._device_name = ""
        self.is_connected = False
        self.error = "JDownloader is disabled to fit memory constraints."

    @new_task
    async def boot(self):
        """Initializes JDownloader process. Disabled for memory optimization."""
        LOGGER.info("JDownloader is disabled to fit memory constraints (512MB RAM target).")
        self.is_connected = False
        self.error = "JDownloader is disabled."


jdownloader = JDownloader()
