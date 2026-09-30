import logging
import unittest
from unittest.mock import patch


class TestCustomFormatter(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # Setup environment and mocks to allow importing update safely
        cls.env_patcher = patch.dict(
            "os.environ", {"BOT_TOKEN": "dummy:token", "UPSTREAM_REPO": ""}
        )

        # Patch subprocess.run to return a mock with returncode=0 to avoid log_error on update failure
        cls.run_patcher = patch("subprocess.run")
        cls.mock_run = cls.run_patcher.start()
        cls.mock_run.return_value.returncode = 0

        cls.exit_patcher = patch("sys.exit")
        cls.remove_patcher = patch("os.remove")
        cls.path_exists_patcher = patch("os.path.exists", return_value=False)
        cls.mongo_patcher = patch("pymongo.mongo_client.MongoClient")

        # Also patch log_info and log_error to suppress output during testing
        cls.log_info_patcher = patch("logging.info")
        cls.log_error_patcher = patch("logging.error")

        cls.env_patcher.start()
        cls.exit_patcher.start()
        cls.remove_patcher.start()
        cls.path_exists_patcher.start()
        cls.mongo_patcher.start()
        cls.log_info_patcher.start()
        cls.log_error_patcher.start()

        import update

        cls.update_module = update

    @classmethod
    def tearDownClass(cls):
        cls.env_patcher.stop()
        cls.run_patcher.stop()
        cls.exit_patcher.stop()
        cls.remove_patcher.stop()
        cls.path_exists_patcher.stop()
        cls.mongo_patcher.stop()
        cls.log_info_patcher.stop()
        cls.log_error_patcher.stop()

    def test_format_info(self):
        formatter = self.update_module.CustomFormatter(
            fmt="%(levelname)s - %(message)s"
        )
        record = logging.LogRecord(
            name="test",
            level=logging.INFO,
            pathname="",
            lineno=0,
            msg="Test message",
            args=(),
            exc_info=None,
        )
        self.assertEqual(formatter.format(record), "I - Test message")

    def test_format_warning(self):
        formatter = self.update_module.CustomFormatter(
            fmt="%(levelname)s - %(message)s"
        )
        record = logging.LogRecord(
            name="test",
            level=logging.WARNING,
            pathname="",
            lineno=0,
            msg="Warning message",
            args=(),
            exc_info=None,
        )
        self.assertEqual(formatter.format(record), "W - Warning message")

    def test_format_error(self):
        formatter = self.update_module.CustomFormatter(
            fmt="%(levelname)s - %(message)s"
        )
        record = logging.LogRecord(
            name="test",
            level=logging.ERROR,
            pathname="",
            lineno=0,
            msg="Error message",
            args=(),
            exc_info=None,
        )
        self.assertEqual(formatter.format(record), "E - Error message")

    def test_format_critical(self):
        formatter = self.update_module.CustomFormatter(
            fmt="%(levelname)s - %(message)s"
        )
        record = logging.LogRecord(
            name="test",
            level=logging.CRITICAL,
            pathname="",
            lineno=0,
            msg="Critical message",
            args=(),
            exc_info=None,
        )
        self.assertEqual(formatter.format(record), "C - Critical message")


if __name__ == "__main__":
    unittest.main()
