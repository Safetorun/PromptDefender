import json
import os

import boto3


def __retrieve_item_if_exists__(key):
    cache_table_name = os.getenv('CACHE_TABLE_NAME')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(cache_table_name)
    response = table.get_item(Key=key)

    if 'Item' in response:
        return json.loads(response['Item'])
    else:
        return None


def lambda_handler(context, event):
    open_api_key = os.getenv("open_ai_api_key")

    return_data = {}
    if __retrieve_item_if_exists__(event['body']):
        return_data = __retrieve_item_if_exists__(event['body'])
    else:
        return_data = {"hello": "world"}

    return {'body': json.dumps(return_data), "statusCode": 200, "headers": {"Content-Type": "application/json"}}
