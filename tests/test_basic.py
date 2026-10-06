import unittest


def rewrite_pixeldrain_url(url: str) -> str:
    if "pixeldrain.com/u/" in url:
        return url.replace("pixeldrain.com/u/", "pixeldrain.com/api/file/")
    return url


class TestBasic(unittest.TestCase):
    def test_pixeldrain_url_rewrite(self):
        original = "https://pixeldrain.com/u/abc12345"
        expected = "https://pixeldrain.com/api/file/abc12345"
        self.assertEqual(rewrite_pixeldrain_url(original), expected)

    def test_pixeldrain_url_non_pixeldrain(self):
        original = "https://example.com/file.txt"
        self.assertEqual(rewrite_pixeldrain_url(original), original)


if __name__ == "__main__":
    unittest.main()
