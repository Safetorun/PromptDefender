import json

from aws_lambda_powertools import Logger
from aws_lambda_powertools import Tracer
from aws_lambda_powertools.utilities.parser import event_parser
from aws_lambda_powertools.utilities.parser.envelopes import ApiGatewayEnvelope
from aws_lambda_powertools.utilities.typing import LambdaContext
from openai import BaseModel
from prompt_defender import build_xml_scanner
from typing import Optional

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


@tracer.capture_lambda_handler
@log_result_information
@cachable_result
@event_parser(model=WallRequest, envelope=ApiGatewayEnvelope)
def lambda_handler(event: WallRequest, _: LambdaContext):
    logger.info("Received event", event=event.json())

    response = WallResponse()
    if event.xml_tag:
        xml_scanner = build_xml_scanner()
        xml_scanner.validate_prompt(event.prompt, xml_tag=event.xml_tag)

    return {"statusCode": 200, "body": response.json()}
