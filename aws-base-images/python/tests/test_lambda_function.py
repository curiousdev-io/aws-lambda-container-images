import json
import pytest
from src.lambda_function import handler


class MockContext:
    function_name = "test_function"
    memory_limit_in_mb = 128
    invoked_function_arn = (
        "arn:aws:lambda:us-east-1:123456789012:function:test_function"
    )
    aws_request_id = "test-request-id"


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
def test_handler(event, expected_message, expected_status):
    context = MockContext()
    result = handler(event, context)
    assert_response(result, expected_message, expected_status)
