import json
import pytest
from src.app import main


def assert_response(result, expected_message, expected_status):
    assert result["statusCode"] == expected_status
    body = json.loads(result["body"])
    assert body["message"] == expected_message
    assert body["status"] == expected_status
    assert "timestamp" in body


@pytest.mark.parametrize(
    "event,expected_message,expected_status",
    [
        (
            {"path": "/hello", "queryStringParameters": {"name": "Alice"}},
            "Hello, Alice",
            200,
        ),
        ({"path": "/hello", "queryStringParameters": {}}, "Hello, World", 200),
        (
            {"path": "/goodbye", "queryStringParameters": {"name": "Bob"}},
            "Goodbye, Bob",
            200,
        ),
        ({"path": "/goodbye", "queryStringParameters": {}}, "Goodbye, World", 200),
        (
            {"path": "/unknown", "queryStringParameters": {"name": "Eve"}},
            "Not found",
            404,
        ),
    ],
)
def test_main(event, expected_message, expected_status):
    result = main(event)
    assert_response(result, expected_message, expected_status)
