import json
from datetime import datetime, UTC


def main(event=None, context=None):
    """
    Main business logic. Can be called from Lambda or Fargate.
    Handles ALB events for /hello and /goodbye routes.
    """
    path = event.get("path") or event.get("rawPath", "")
    query = event.get("queryStringParameters", {}) or {}
    name = query.get("name", "World")
    now = datetime.now(UTC).isoformat()
    status = 200
    if path.startswith("/hello"):
        message = f"Hello, {name}"
    elif path.startswith("/goodbye"):
        message = f"Goodbye, {name}"
    else:
        status = 404
        message = "Not found"
    body = json.dumps({"timestamp": now, "status": status, "message": message})
    return {
        "statusCode": status,
        "body": body,
        "headers": {"Content-Type": "application/json"},
    }
