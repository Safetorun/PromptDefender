[![Codacy
Badge](https://app.codacy.com/project/badge/Grade/080ff8f6c80d434484249b8dbb3a5ef0)](https://app.codacy.com/gh/Safetorun/PromptShield/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

[![Deploy](https://github.com/Safetorun/PromptDefender/actions/workflows/deploy.yml/badge.svg)](https://github.com/Safetorun/PromptDefender/actions/workflows/deploy.yml)

[Documentation](https://promptshield.readme.io/)

## Try out the hosted Hosted version

To use "Keep", go to: [PromptDefender Keep](https://defender.safetorun.com)

To use the APIs - check out our [Developer Portal](https://promptshield.readme.io)

## What is Prompt Defender?

A prompt defence is a multi-layer defence that can be used to protect your applications against prompt injection
attacks. You can use this with any LLM APIs (whether Bard, LlaMa, ChatGPT - or any other LLM) These types of attack are
complex, and are difficult to solve with a single layer of defence - as such, a prompt shield is made up of multiple '
rings' of defence.

**Ring 1 - Wall**

Ring 1 is the first layer of defence, and is intended to sanitise input before it moves through the layers of defence.
This will typically look at prompt input, and ensure that it meets certain rules. For example:

- Does it contain keywords that are known for jail-breaking attacks
- Does the information reveal PII which should not be provided to your LLM (e.g. email addresses, phone numbers, etc)
- Is this prompt from a user / ip address (or any other identifier you want to provide) which is probing or attacking
  your system? [Coming soon]

**Ring 2 - Keep**

Ring 2 is a layer of defence on the prompt itself - it effectively wraps your prompt in an effective 'prompt defence'
which provides instructions to the LLM as part of the prompt on what should happen, and what it should avoid doing (e.g.
reminders not to leak a secret key)

**Ring 3 - Drawbridge [Coming soon] **

Ring 3 is a final protection which looks at the returned value prior to it being provided to a client or using it for a
follow-up action; this can contain defences such as:

- Avoid returning data containing a XSS or script tags
- Avoid returning information which has proprietary or secret information in it

## Running integration tests

To run the integration tests, run the following command:

```bash
make integration_test
```

To debug in intellij, run the tests in `run_integration_cucumber_tests.go` with the following environment variables set:

```bash
URL
DEFENDER_API_KEY
```

You can get these after a `make deploy` with the following commands:

```bash
	export URL=`cd terraform && terraform output -json | dasel select -p json '.api_url.value' | tr -d '"'`
	export DEFENDER_API_KEY=`cd terraform && terraform output -json | dasel select -p json '.api_key_value.value' | tr -d '"'`
```

## Response times

### Tests
There are a k6 load tests in the test/load directory. 

Inside each test files are the response time to check for test adherence


### Deployment

* First, deploy the terraform-base-infrastructure which contains the huggingface/debert dataset. To do this, run:

`make deploy-base-infrastructure`

* Now, deploy the main infrastructure. To do this, run:

`make deploy`

Get the URL and API key from the terraform output and set them as environment variables:

```bash
export URL=`cd terraform && terraform output -json | dasel select -p json '.api_url.value' | tr -d '"'`
export DEFENDER_API_KEY=`cd terraform && terraform output -json | dasel select
```

* Now, run the integration tests if you want to check the setup:

`make integration_test`
