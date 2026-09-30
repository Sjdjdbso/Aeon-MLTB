import unittest

from bot.core.config_manager import Config
from bot.core.telegram_manager import TgClient


class TestProxyHandling(unittest.IsolatedAsyncioTestCase):
    def setUp(self):
        Config.TG_PROXY = {}

    def test_empty_tg_proxy_normalization(self):
        Config.set("TG_PROXY", {})
        self.assertEqual(Config.TG_PROXY, {})

    def test_missing_scheme_tg_proxy_normalization(self):
        Config.set("TG_PROXY", {"hostname": "127.0.0.1", "port": 1080})
        self.assertEqual(Config.TG_PROXY, {})

    def test_valid_tg_proxy_normalization(self):
        valid_proxy = {
            "scheme": "socks5",
            "hostname": "127.0.0.1",
            "port": 1080,
        }
        Config.set("TG_PROXY", valid_proxy)
        self.assertEqual(Config.TG_PROXY, valid_proxy)

    def test_string_tg_proxy_normalization(self):
        Config.set(
            "TG_PROXY",
            "{'scheme': 'socks5', 'hostname': '127.0.0.1', 'port': 1080}",
        )
        self.assertEqual(
            Config.TG_PROXY,
            {"scheme": "socks5", "hostname": "127.0.0.1", "port": 1080},
        )

    async def test_get_proxy_returns_none_for_empty(self):
        Config.TG_PROXY = {}
        proxy = await TgClient._get_proxy()
        self.assertIsNone(proxy)

    async def test_get_proxy_returns_dict_for_valid(self):
        valid_proxy = {
            "scheme": "socks5",
            "hostname": "127.0.0.1",
            "port": 1080,
        }
        Config.TG_PROXY = valid_proxy
        proxy = await TgClient._get_proxy()
        self.assertEqual(proxy, valid_proxy)


if __name__ == "__main__":
    unittest.main()
