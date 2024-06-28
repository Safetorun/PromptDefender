import os
from typing import Optional, Dict

import boto3
from aws_lambda_powertools import Logger
from aws_lambda_powertools import Tracer
from aws_lambda_powertools.utilities.parser import parse
from aws_lambda_powertools.utilities.parser.envelopes import ApiGatewayEnvelope
from aws_lambda_powertools.utilities.typing import LambdaContext
from prompt_defender import build_xml_scanner
from prompt_defender_aws_defences.wall import AwsPIIScannerWallExecutor, SagemakerWallExecutor
from pydantic import BaseModel

from badwords import find_lowest_score
from shared import log_result_information, cachable_result

logger = Logger(service="PromptDefender-Wall")
tracer = Tracer()


class WallRequest(BaseModel):
    prompt: str
    user_id: Optional[str] = None
    session_id: Optional[str] = None
    scan_pii: Optional[bool] = None
    xml_tag: Optional[str] = None
    check_badwords: Optional[bool] = None
    fast_check: Optional[bool] = None


class WallResponse(BaseModel):
    contains_pii: Optional[bool] = None
    potential_jailbreak: Optional[bool] = None
    potential_xml_escaping: Optional[bool] = None
    suspicious_user: Optional[bool] = None
    suspicious_session: Optional[bool] = None
    modified_prompt: Optional[str] = None


def check_for_suspicous_session(session_id: str, api_key_id: str) -> bool:
    cache_table_name = os.getenv('USERS_TABLE')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(cache_table_name)
    value = table.get_item(Key={'UserOrSessionId': session_id, 'ApiKeyId': api_key_id})
    return 'Item' in value


@tracer.capture_lambda_handler
@log_result_information
@cachable_result
def lambda_handler(event: Dict, _: LambdaContext):
    api_key_id = event['requestContext']['identity']['apiKeyId']
    event = parse(event, model=WallRequest, envelope=ApiGatewayEnvelope)
    logger.info('Received event', event=event.json())

    response = WallResponse()
    response.potential_jailbreak = False

    if event.xml_tag:
        logger.info('Going to scan for XML escaping')
        xml_scanner = build_xml_scanner()
        xml_escape_result = xml_scanner.is_user_input_safe(event.prompt, xml_tag=event.xml_tag)
        response.potential_xml_escaping = xml_escape_result.unacceptable_prompt

    if event.session_id:
        logger.info('Going to check for suspicious session')
        response.suspicious_session = check_for_suspicous_session(event.session_id, api_key_id)

    if event.user_id:
        logger.info('Going to check for suspicious user')
        response.suspicious_user = check_for_suspicous_session(event.user_id, api_key_id)

    if event.scan_pii:
        logger.info('Going to scan for PII entities')
        scanner = AwsPIIScannerWallExecutor()
        result = scanner.is_user_input_safe(event.prompt)
        response.contains_pii = result.unacceptable_prompt
    else:
        logger.info('Skipping PII scan')

    if event.check_badwords:
        logger.info("Going to check for bad words")
        lowest_score = find_lowest_score(event.prompt)
        logger.info('Lowest score:', lowest_score)
        response.potential_jailbreak = lowest_score < 0.20

    if not event.fast_check and not response.potential_jailbreak:
        logger.info("Going to check for potential jailbreak")
        long_check = SagemakerWallExecutor(os.getenv("SAGEMAKER_ENDPOINT_JAILBREAK"))
        long_check_result = long_check.is_user_input_safe(event.prompt)
        response.potential_jailbreak = long_check_result.unacceptable_prompt
    else:
        logger.info("Skipping potential jailbreak check")

    return {"statusCode": 200, "body": response.json()}
