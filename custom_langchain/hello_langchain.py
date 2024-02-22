import json
import os

import requests
from langchain.prompts import StringPromptTemplate
from langchain.prompts.pipeline import PipelinePromptTemplate
from langchain.prompts.prompt import PromptTemplate
from langchain_core.prompt_values import PromptValue, StringPromptValue


class UserInputValidator(StringPromptTemplate):
    _api_key: str
    _block_pii: bool = False
    _xml_to_escape: str = None

    def __init__(self, **kwargs):
        super().__init__(**kwargs)

        if "PROMPT_DEFENDER_API_KEY" in kwargs:
            self._api_key = kwargs["PROMPT_DEFENDER_API_KEY"]
        elif "PROMPT_DEFENDER_API_KEY" in os.environ:
            self._api_key = os.environ["PROMPT_DEFENDER_API_KEY"]
        else:
            raise ValueError(
                "PROMPT_DEFENDER_API_KEY not found. Either pass as an argument or set as an environment variable.")

        self._block_pii = False
        self._xml_to_escape = None

        if "xml_tag" in kwargs:
            self._xml_to_escape = kwargs["xml_tag"]

        if "block_pii" in kwargs:
            self._block_pii = kwargs["block_pii"]

    def __retrieve_data__(self, prompt: str) -> bool:
        url = "https://prompt.safetorun.com/moat"
        headers = {
            "accept": "application/json",
            "content-type": "application/json",
            "x-api-key": self._api_key
        }

        body_builder = {"prompt": prompt}
        if self._xml_to_escape:
            body_builder["xml_tag"] = self._xml_to_escape

        if self._block_pii:
            body_builder["block_pii"] = True

        data = requests.post(url, headers=headers, data=json.dumps(body_builder)).json()
        contains_something_dodgy = ((data.get("contains_pii", False)) or
                                    (data.get("potential_xml_escaping", False)) or
                                    (data.get("potential_jailbreak", False)) or
                                    (data.get("suspicious_user", False)) or
                                    (data.get("suspicious_session", False))
                                    )

        return contains_something_dodgy

    def format_prompt(self, **kwargs: any) -> PromptValue:
        """Create Prompt Value."""
        for k in self.input_variables:
            if self.__retrieve_data__(kwargs[k]):
                raise ValueError("Invalid input")
        return StringPromptValue(text=kwargs[self.input_variables[0]])

    def format(self, **kwargs: any) -> str:
        for k in self.input_variables:
            if self.__retrieve_data__(kwargs[k]):
                raise ValueError("Failed prompt defender check")

        return kwargs[self.input_variables[0]]


if __name__ == "__main__":
    validator = UserInputValidator(input_variables=["text_input"])

    full_template = """{user_input}"""
    full_prompt = PromptTemplate.from_template(full_template)

    input_prompts = [
        ("user_input", validator),
    ]

    pipeline_prompt = PipelinePromptTemplate(
        final_prompt=full_prompt, pipeline_prompts=input_prompts
    )

    print(
        pipeline_prompt.format(
            text_input="test",
        )
    )
