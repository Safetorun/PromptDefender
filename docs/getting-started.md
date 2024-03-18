---
title: Building your Prompt Defence
excerpt: This guide provides step-by-step instructions on building your Prompt Defence, starting with the creation of your 'keep', followed by the construction of a 'wall'. It also introduces the upcoming 'Drawbridge' feature for enhanced security.
category: 65f04fd9bf46c7003666a5fb
---

A prompt defence is a multi-layer defence that can be used to protect your applications against prompt injection
attacks. You can use this with any LLM APIs (whether Bard, LlaMa, ChatGPT - or any other LLM) These types of attack are
complex, and are difficult to solve with a single layer of defence - as such, a prompt shield is made up of multiple '
rings' of defence.

**Wall**

Wall is the first layer of defence, and is intended to sanitise input before it moves through the layers of defence.
This will typically look at prompt input, and ensure that it meets certain rules. For example:

- Does it contain keywords that are known for jail-breaking attacks
- Does the information reveal PII which should not be provided to your LLM (e.g. email addresses, phone numbers, etc)
- Is this prompt from a user / ip address (or any other identifier you want to provide) which is probing or attacking
  your system? [Coming soon]

**Keep**

Keep is a layer of defence on the prompt itself - it effectively wraps your prompt in an effective 'prompt defence'
which provides instructions to the LLM as part of the prompt on what should happen, and what it should avoid doing (e.g.
reminders not to leak a secret key)

**Drawbridge**

Ring 3 is a final protection which looks at the returned value prior to it being provided to a client or using it for a
follow-up action; this can contain defences such as:

- Avoid returning data containing a XSS or script tags
- Avoid returning information which has proprietary or secret information in it

# Start by building your prompt defence

The first step in building out your Prompt Defence is to build your 'keep' - your keep involves changing the prompt
itself to give a clue to your LLM about how to use it safely.

Take your existing prompt and pass it as the 'prompt' into the keep endpoint.

We also recommend randomising your tag.

You can always head to \<<https://defender.safetorun.com/keep> if you want to use the UI to generate your prompt
defence.

Once you've generated a shielded prompt - take a note of that, and the returned tag

[block:tutorial-tile]
{
"backgroundColor": "#018FF4",
"emoji": "🦉",
"id": "65ce49ae8f905e002002b1eb",
"link": "https://promptshield.readme.io/v1.0/recipes/building-your-prompt-defence",
"slug": "building-your-prompt-defence",
"title": "Building your prompt defence"
}
[/block]

# Now, build a Wall

**PromptDefender - Wall is executed before a request to your LLM. **

As an example - if you're taking user input and appending it to your prompt (now shielded) - then sending that to chat
GPT or some other LLM - call this endpoint just before you call the LLM, and pass in the user's input as the 'prompt'
parameter.

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

# Finish it up with a Drawbridge (Coming soon)

With Prompt Defender - Drawbridge, you can detect leaks in your prompts, detect PII about to be escaped and strip any
potentially damaging XSS from being release