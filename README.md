[![Codacy 
Badge](https://app.codacy.com/project/badge/Grade/080ff8f6c80d434484249b8dbb3a5ef0)](https://app.codacy.com/gh/Safetorun/PromptShield/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

[![Deploy](https://github.com/Safetorun/PromptDefender/actions/workflows/deploy.yml/badge.svg)](https://github.com/Safetorun/PromptDefender/actions/workflows/deploy.yml)

[Documentation](https://promptshield.readme.io/)

A prompt defence is a multi-layer defence that can be used to protect your applications against prompt injection attacks. You can use this with any LLM APIs (whether Bard, LlaMa, ChatGPT - or any other LLM) These types of attack are complex, and are difficult to solve with a single layer of defence - as such, a prompt shield is made up of multiple 'rings' of defence.

**Ring 1 - Moat**

Ring 1 is the first layer of defence, and is intended to sanitise input before it moves through the layers of defence. This will typically look at prompt input, and ensure that it meets certain rules. For example:

- Does it contain keywords that are known for jail-breaking attacks
- Does the information reveal PII which should not be provided to your LLM (e.g. email addresses, phone numbers, etc)
- Is this prompt from a user / ip address (or any other identifier you want to provide) which is probing or attacking your system? [Coming soon]

**Ring 2 - Walls**

Ring 2 is a layer of defence which uses AI in order to detect attempted attacks or jailbreaks.

**Ring 3 - Keep**

Ring 3 is a layer of defence on the prompt itself - it effectively wraps your prompt in an effective 'prompt defence' which provides instructions to the LLM as part of the prompt on what should happen, and what it should avoid doing (e.g. reminders not to leak a secret key)

**Ring 4 - Drawbridge**

Ring 4 is a final protection which looks at the returned value prior to it being provided to a client or using it for a follow-up action; this can contain defences such as:

- Avoid returning data containing a XSS or script tags
- Avoid returning information which has proprietary or secret information in it
