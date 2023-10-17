---
title: Building your defence
excerpt: Now that we know the names, let's start building up our layers of defence against attackers. 
category: 652be292f7eae600244211de
---

# Keep

We will first build your keep - this is the activity you will perform least, but is one of the most effective ways to
protect your application. It involves designing your base prompt so that it is able to defend itself against attacks. 

This is effective as it does not add any additional API calls to your application before after the prompt is called, and
instead it something you can perform once and then forget about.

# Moat

The next thing to do is to build your moat. For the prompt inside your application, ask yourself if you want Personally identifiable
information PII to be detected? 

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