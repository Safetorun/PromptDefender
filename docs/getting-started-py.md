---
title: Getting started - python
excerpt: To get started with python, first install the python packages using pip
category: 65f04fd9bf46c7003666a5fb
---

# Getting started - Python

To get started with python, first install the python packages using pip:

```shell
pip install prompt-defender
```

# API key

To use the remote API in wall or drawbridge, you'll first need an API key. You can get this at:

[https://defender.safetorun.com](https://defender.safetorun.com)

# Example python use


Then you can use the package in your code - here's an example using wall:

```python
from wall import should_block_prompt
from wall_builder import create_wall

wall = create_wall(
    # These options first require you to have a Prompt Defender account which you can sign up for at
    # https://defender.safetorun.com. Once you have an account you can get an API key  to use with the wall.
    remote_jailbreak_check=True,
    api_key="test_key",
    user_id="test_user",
    session_id="test_session",
    allow_pii=False,

    # When you create a prompt, with Prompt Defender - Keep, you will get
    # an XML tag that wraps user input. Pass this tag to the remote endpoint
    # in order to check for potential XML escaping which is likely because
    # someone is trying to attack your system
    xml_tag="tag",

    # The following are used for prompt validation - if you are only
    # expecting a certain number of values, or a certain length of prompt
    # you can use these to enforce that.
    max_prompt_length=100,
    allowed_prompt_values=["hello", "world"]
)

validation_response = wall.validate_prompt("hello")

if validation_response.contains_pii:
    print("Prompt contains PII")
elif validation_response.suspicious_user:  # etc etc etc
    print("Prompt is suspicious")
elif should_block_prompt(validation_response):
    print("Prompt should be blocked")
else:
    print("Prompt is OK")
```


# Next steps

Check out the individual guides on wall and drawbridge, or alternatively check out the recipes:


[block:tutorial-tile]
{
  "backgroundColor": "#018FF4",
  "emoji": "🦉",
  "id": "65ce4ca087d73c004bcc343e",
  "link": "https://promptshield.readme.io/recipes/python-sdk",
  "slug": "build-your-wall",
  "title": "Build your Wall"
}
[/block]