import unittest
from unittest import mock
import redis
from favicon import Icon
from extract_favicon import extract_favicon

class MockResponse:
    def __init__(self, content=""):
        self.content = content
        self.status_code = 200

    def content():
        return self.content

# Mocks for favicon.get() 
def mock_responses(*args, **kwargs):
    if args[0] == "example.com":
        return [Icon("https://example.com/favicon.ico", 16, 16, 'ico')]
    if args[0] == "https://example.com/favicon.ico":
        return MockResponse(content=b'FAKE_FAVICON_CONTENT')


class TestExtractFaviconAndStoreInRedis(unittest.TestCase):
    def setUp(self):
        self.redis_client = redis.Redis(host="localhost", port=6379, db=0)
        self.redis_client.flushdb()

    def tearDown(self):
        self.redis_client.flushdb()

    def test_extract_favicon(self):

        # Call the function to extract the favicon and store it in Redis
        with mock.patch('extract_favicon.favicon.get', return_value=[Icon("https://example.com/favicon.ico", 16, 16, 'ico')]):
            with mock.patch('extract_favicon.requests.get', return_value=MockResponse(content=b'FAKE_FAVICON_CONTENT')):
                extract_favicon("example.com")

        # Verify that the favicon was stored in Redis
        favicon_content = self.redis_client.get("example.com")
        self.assertEqual(favicon_content, b"FAKE_FAVICON_CONTENT")

if __name__ == "__main__":
    unittest.main()

