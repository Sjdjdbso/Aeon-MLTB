import unittest

from bot.helper.ext_utils.links_utils import rewrite_pixeldrain


class TestPixeldrainRewriter(unittest.TestCase):
    def test_pixeldrain_url_rewriting(self):
        url = "https://pixeldrain.com/u/xyz12345"
        expected = "https://pixeldrain.com/api/file/xyz12345"
        self.assertEqual(rewrite_pixeldrain(url), expected)

    def test_non_pixeldrain_url(self):
        url = "https://example.com/file.zip"
        self.assertEqual(rewrite_pixeldrain(url), url)

    def test_already_api_url(self):
        url = "https://pixeldrain.com/api/file/xyz12345"
        self.assertEqual(rewrite_pixeldrain(url), url)

    def test_non_string_input(self):
        self.assertIsNone(rewrite_pixeldrain(None))


if __name__ == "__main__":
    unittest.main()
