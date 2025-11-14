from typing import Dict

from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.typing import LambdaContext
from app import main

logger = Logger()


@logger.inject_lambda_context
def handler(event: Dict, context: LambdaContext) -> Dict[str, str]:
    response = main(event, context)
    logger.info(response)
    return response
