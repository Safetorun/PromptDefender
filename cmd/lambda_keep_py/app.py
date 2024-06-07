import json
import os
from typing import Optional

import boto3
from aws_lambda_powertools import Logger
from aws_lambda_powertools import Tracer
from aws_lambda_powertools.utilities.parser import event_parser
from aws_lambda_powertools.utilities.parser.envelopes import ApiGatewayEnvelope
from aws_lambda_powertools.utilities.typing import LambdaContext
from langchain_openai.chat_models import ChatOpenAI
from prompt_defender_llm_defences import KeepExecutorLlm
from pydantic import BaseModel

from cache.cache import store_item
from shared import cachable_result
from settings.ssm_retriever import get_secret

logger = Logger(service="PromptDefender-Keep")
tracer = Tracer()

os.environ["OPENAI_API_KEY"] = get_secret(os.getenv("OPENAI_SECRET_NAME"))


class KeepRequest(BaseModel):
    prompt: str
    randomise_xml_tag: Optional[bool] = False


class KeepResponse(BaseModel):
    shielded_prompt: str
    xml_tag: str
    canary: Optional[str]


@tracer.capture_lambda_handler
@cachable_result
@event_parser(model=KeepRequest, envelope=ApiGatewayEnvelope)
def lambda_handler(event: KeepRequest, _: LambdaContext):
    logger.info("Received event", event=event.dict())

    llm = ChatOpenAI(model="gpt-4o")
    executor = KeepExecutorLlm(llm=llm)

    safe_prompt = executor.generate_prompt(event.prompt, event.randomise_xml_tag)

    return_data = {
        "shielded_prompt": safe_prompt.safe_prompt,
        "xml_tag": safe_prompt.xml_tag,
        "canary": safe_prompt.canary
    }

    return {"statusCode": 200, "body": KeepResponse(**return_data).json()}


if __name__ == "__main__":
    print(lambda_handler(KeepRequest(prompt="hello", randomise_xml_tag=False), None))
