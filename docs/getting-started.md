---
title: Building your Prompt Defence
excerpt: This guide provides step-by-step instructions on building your Prompt Defence, starting with the creation of your 'keep', followed by the construction of a 'wall'. It also introduces the upcoming 'Drawbridge' feature for enhanced security.  
category: 65f04fd9bf46c7003666a5fb
---

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