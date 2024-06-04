import json
import os

import boto3
from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.parser import event_parser
from aws_lambda_powertools.utilities.typing import LambdaContext
from langchain_openai.chat_models import ChatOpenAI
from prompt_defender_llm_defences import KeepExecutorLlm
from pydantic import BaseModel
from aws_lambda_powertools import Tracer
from cache.cache import retrieve_item_if_exists
from settings.ssm_retriever import get_secret

logger = Logger(service="APP")
tracer = Tracer()

os.environ["OPENAI_API_KEY"] = get_secret(os.getenv("OPENAI_SECRET_NAME"))


def __retrieve_item_if_exists__(key):
    cache_table_name = os.getenv('CACHE_TABLE_NAME')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(cache_table_name)

    response = table.get_item(Key={"Id": key})

    if 'Item' in response:
        return json.loads(response['Item'])
    else:
        return None


class KeepRequest(BaseModel):
    prompt: str
    randomise_xml_tag: bool


class KeepResponse(BaseModel):
    shielded_prompt: str
    xml_tag: str
    canary: str


@tracer.capture_lambda_handler
@event_parser(model=KeepRequest)
def lambda_handler(event: KeepRequest, _: LambdaContext) -> KeepResponse:
    logger.debug("Received event", event=event.dict())

    return_data = retrieve_item_if_exists(event.json(), retrieve_function=__retrieve_item_if_exists__)

    if return_data is not None:
        logger.debug("Retrieved from cache", return_data=return_data)
        return KeepResponse(**return_data)

    logger.debug("Not found in cache, generating new data")

    llm = ChatOpenAI(model="gpt-4o")
    executor = KeepExecutorLlm(llm=llm)

    safe_prompt = executor.generate_prompt(event.prompt)

    return_data = {
        "shielded_prompt": safe_prompt.safe_prompt,
        "xml_tag": safe_prompt.xml_tag,
        "canary": safe_prompt.canary
    }

    return KeepResponse(**return_data)
