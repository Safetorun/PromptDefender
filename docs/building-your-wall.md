---
title: Building your Wall
excerpt: Wall will look at if PII is being sent in the request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to bypass your keep defence), it can also detect if the user or session is marked as suspicious.
category: 65f7e5d65b0c910060007711
---

The next thing to do is to build your wall. This adds the first layer of defence, looking at if PII is being sent in the
request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to
bypass
your keep defence), it can also detect if the user or session is marked as suspicious.

## Recipe

Walk through this recipe for more information on how to build your wall with the REST API:

[block:tutorial-tile]
{
"backgroundColor": "#018FF4",
"emoji": "🦉",
"id": "65ce4ca087d73c004bcc343e",
"link": "https://promptshield.readme.io/v1.0/recipes/build-your-moat",
"slug": "build-your-moat",
"title": "Build your Wall"
}
[/block]

or use the following example to build your wall with the Python SDK:

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

## Jailbreak detection

Prompt Defender uses a combination of methods to detect potential jailbreaks. This includes looking at the prompt and
checking for "bad" words, which could indicate that the user is attempting to jailbreak - as well as other machine
learning models which look at the prompt and determine if it is likely to be a jailbreak attempt.

The only thing that is required for your request is the prompt itself.

```python 
from wall import create_wall

wall = create_wall(
    remote_jailbreak_check=True,
    api_key="test_key"
)

response = wall.validate_prompt("What is the capital of France?")
```
```json
{
  "prompt": "What is the capital of France?"
}
```

Your response will include:

```python
response.potential_jailbreak == False
```
```json
{
  "potential_jailbreak": false
}
```

## PII Detection

PII Detection involves looking at user input and determining if it contains information which could be used to identify
a user. This could be email addresses, phone numbers, names, addresses, etc. If PII Detection is enabled, then the
led, then the  
prompt will be checked for any personal information.

In order to enable PII Detection, you will need to add the following to your request:

```python
from wall import create_wall

wall = create_wall(
    # These options first require you to have a Prompt Defender account which you can sign up for at
    # https://defender.safetorun.com. Once you have an account you can get an API key  to use with the wall.
    remote_jailbreak_check=True,
    api_key="test_key",
    allow_pii=False
)
response = wall.validate_prompt("What is the capital of France?")
```
```json
{
  "prompt": "What is the capital of France?",
  "scan_pii": true
}
```


Your response will include:

```python
response.contains_pii == False
```
```json
{
  "contains_pii": true
}
```

## User and session detection

You can also enable user and session detection. This will look at the user and session information provided in the
request, and will check to see if this user or session has been flagged as suspicious. This is useful if you have a user
who is attempting to attack your system, and you want to block them from using your LLMs.

To use this, you will need to add the following to your request:


```python
from wall import create_wall

wall = create_wall(
    # These options first require you to have a Prompt Defender account which you can sign up for at
    # https://defender.safetorun.com. Once you have an account you can get an API key  to use with the wall.
    remote_jailbreak_check=True,
    api_key="test_key",
    user_id="1234",
    session_id="5678"
)
response = wall.validate_prompt("What is the capital of France?")
```
```json
{
  "prompt": "What is the capital of France?",
  "user_id": "1234",
  "session_id": "5678"
}
```

Both are optional - you can use either, neither or both. Your response will include:

```python
response.suspicious_user == False
response.suspicious_session == False
```
```json
{
  "suspicious_session": true,
  "suspicious_user": true
}
```

## XML Escaping

If you are you using a keep defence, you can also enable XML escaping. This will look at the prompt and will check to
see if it contains any XML tag escaping which will likely indicate an attacker is trying to bypass your prompt defence

Find the XML tag that is used in your prompt defence - for example, if you are using the keep defence, your prompt will
contain XML tags that encapsulated the user's input.

For example; if your prompt is:

```
Take the user's input below and answer their questions about cats 

<user_input>%s</user_input>
```

Then send the following to your request: (in python, this is executed locally on the device and does not require an API)

```python
from wall import create_wall

wall = create_wall(
    xml_tag="user_input"
)
response = wall.validate_prompt("What is the capital of France?")
```
```json
{
  "xml_tag": "user_input"
}
```

Your response will then contain:

```python
response.potential_xml_escaping == False
```
```json
{
  "potential_xml_escaping": true
}
```