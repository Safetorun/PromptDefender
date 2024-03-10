---
title: Building your Wall
excerpt: Wall will look at if PII is being sent in the request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to bypass your keep defence), it can also detect if the user or session is marked as suspicious. 
category: 652be292f7eae600244211de
---

The next thing to do is to build your wall. This adds the first layer of defence, looking at if PII is being sent in the
request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to bypass
your keep defence), it can also detect if the user or session is marked as suspicious.

## PII Detection

PII Detection involves looking at user input and determining if it contains information which could be used to identify a user.
This could be email addresses, phone numbers, names, addresses, etc. If PII Detection is enabled, then the prompt will be
checked for any personal information.

In order to enable PII Detection, you will need to add the following to your request:

```json
{
    "pii": true
}
```

## Jailbreak detection

It is recommended to turn on jail-break detection; for now the only option available in 'Basic' jailbreak detection. This will look for keywords which are known to be used in jailbreak attacks. This might check for things like "ignore previous";
for each of these commands - Prompt Defender will look for common techniques such a base64 encoding that someone might use to
disguise a jailbreak attack.

In order to enable jailbreak detection, you will need to add the following to your request:

```json
{
    "jailbreak": 0
}
```

## User and session detection

You can also enable user and session detection. This will look at the user and session information provided in the request, and
will check to see if this user or session has been flagged as suspicious. This is useful if you have a user who is attempting
to attack your system, and you want to block them from using your LLMs.

## Script detection

You can also enable script detection. This will look at the prompt and will check to see if it contains any script tags or
other HTML tags which could be used to inject scripts into your application. This is useful as there are a lot of cases
where script tags and html tags are not needed, and so this can be used to block them.

## XML Escaping

If you are you using a keep defence (above), you can also enable XML escaping. This will look at the prompt and
will check to see if it contains any XML tag escaping which will likely indicate an attacker is trying to bypass your prompt defence
