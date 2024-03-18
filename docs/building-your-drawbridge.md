---
title: Building your Drawbridge
excerpt:
category: 65f7e5d65b0c910060007711
---

Drawbridge is a part of the Prompt Defender project that is responsible for validating the response of an LLM execution.
That is - after you have executed an LLM, you can use Drawbridge to check the response for any potential security
issues.

> 📘 Drawbridge is currently only available from within the python SDK

Drawbridge is used to validate the response of an LLM execution. It has two main functionalities:

* Checking for a canary in the response.

* Cleaning the response by removing scripts if allow_unsafe_scripts is set to False.

## Recipe

[block:tutorial-tile]
{
"backgroundColor": "#018FF4",
"emoji": "🦉",
"id": "65f7e7ce345b2c0025ebb6f9",
"link": "https://promptshield.readme.io/v1.0/recipes/building-your-drawbridge",
"slug": "building-your-drawbridge",
"title": "Building your Drawbridge"
}
[/block]

## Example

Here is example usage:

```python
from drawbridge import build_drawbridge

# Create a Drawbridge instance
drawbridge = build_drawbridge(canary="test_canary")

# Validate and clean a response
response = "<script>alert('Hello!');</script>test_canary"
response_ok, cleaned_response = drawbridge.validate_response_and_clean(response)

print(f"Response OK: {response_ok}")
print(f"Cleaned Response: {cleaned_response}")
```

In this example, we first import the `build_drawbridge` function from the drawbridge module. We then use this function
to
create a Drawbridge instance, specifying a canary string that we want to check for in the response.

Next, we have a response string that we want to validate and clean. We pass this response to the
validate_response_and_clean method of our Drawbridge instance. This method returns two values:

* response_ok: This is a boolean value that indicates whether the canary was found in the response.
* cleaned_response: This is the cleaned version of the response. If allow_unsafe_scripts is False (which is the
  default),
  any scripts in the response will be removed.