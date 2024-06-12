import json
import os
import time
from typing import Callable

import boto3
from aws_lambda_powertools import Logger
from aws_lambda_powertools.middleware_factory import lambda_handler_decorator
from aws_lambda_powertools.utilities.typing import LambdaContext

from cache.cache import retrieve_item_if_exists, store_item

logger = Logger(service="PromptDefender-Keep")


def __retrieve_item_if_exist(key):
    cache_table_name = os.getenv('CACHE_TABLE_NAME')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(cache_table_name)

    response = table.get_item(Key={"Id": key})

    if 'Item' in response:
        return json.loads(response['Item']['Value'])
    else:
        return None


def __store_item(key, item):
    cache_table_name = os.getenv('CACHE_TABLE_NAME')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(cache_table_name)

    table.put_item(Item={"Id": key, "Value": json.dumps(item)})


@lambda_handler_decorator
def cachable_result(
        handler: Callable[[dict, LambdaContext], dict],
        event: dict,
        context: LambdaContext,
) -> dict:
    start_time = time.time()

    return_data = retrieve_item_if_exists(event['body'], retrieve_function=__retrieve_item_if_exist)

    if return_data is not None:
        logger.info("Retrieved from cache", return_data=return_data)
        return {"statusCode": 200, 'body': return_data}
    logger.info("Not found in cache, generating new data")

    result = handler(event, context)

    logger.info("Storing in cache {}", result)

    store_item(event['body'], result['body'], store_function=__store_item)

    queue_message = to_request_log(event)
    queue_message["Response"] = result["body"]
    queue_message["Time"] = int((time.time() - start_time) * 1000)

    log_summary_message(queue_message)

    return result


def to_request_log(request: dict) -> dict:
    return {
        "UserId": request["requestContext"]["identity"]["apiKeyId"],
        "Domain": request["requestContext"]["domainName"],
        "Headers": request["headers"],
        "Method": request["requestContext"]["httpMethod"],
        "QueryParams": request["queryStringParameters"],
        "HttpMethod": request["requestContext"]["httpMethod"],
        "StartedDateTime": request["requestContext"]["requestTime"],
        "HttpResponseHeaders": {
            "content-type": "application/json",
        },
        "HttpResponse": 200,
        "Endpoint": request["path"],
        "Request": request["body"],
    }


def log_summary_message(message: dict):
    try:
        json_message = json.dumps(message)
        logger.info("Summary message ---- %s", json_message)
    except json.JSONDecodeError as err:
        logger.error("JSON marshaling error: %s", err)


@lambda_handler_decorator
def log_result_information(
        handler: Callable[[dict, LambdaContext], dict],
        event: dict,
        context: LambdaContext,
) -> dict:
    start_time = time.time()

    result = handler(event, context)

    queue_message = to_request_log(event)
    queue_message["Response"] = result["body"]
    queue_message["Time"] = int((time.time() - start_time) * 1000)

    return result
