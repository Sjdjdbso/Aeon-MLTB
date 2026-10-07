import unittest

from bot.core.config_manager import Config
from bot.core.telegram_manager import TgClient


class TestProxyHandling(unittest.TestCase):
    def test_empty_proxy_returns_none(self):
        Config.TG_PROXY = {}
        self.assertIsNone(TgClient._get_proxy())

    def test_proxy_without_scheme_returns_none(self):
        Config.TG_PROXY = {"hostname": "127.0.0.1", "port": 1080}
        self.assertIsNone(TgClient._get_proxy())

    def test_valid_proxy_returns_dict(self):
        proxy = {"scheme": "socks5", "hostname": "127.0.0.1", "port": 1080}
        Config.TG_PROXY = proxy
        self.assertEqual(TgClient._get_proxy(), proxy)


if __name__ == "__main__":
    unittest.main()
