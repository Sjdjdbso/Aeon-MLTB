import unittest

from bot.core.config_manager import Config


class TestConfigManager(unittest.TestCase):
    def test_default_upload_normalization(self):
        Config.set("DEFAULT_UPLOAD", "pixeldrain")
        self.assertEqual(Config.DEFAULT_UPLOAD, "pixeldrain")

        Config.set("DEFAULT_UPLOAD", "invalid_option")
        self.assertEqual(Config.DEFAULT_UPLOAD, "gd")

    def test_tg_proxy_normalization(self):
        Config.set("TG_PROXY", {})
        self.assertEqual(Config.TG_PROXY, {})


if __name__ == "__main__":
    unittest.main()
