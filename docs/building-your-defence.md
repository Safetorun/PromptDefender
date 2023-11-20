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

Check out a [guide on building your keep](/building-your-keep.md) for more information.
Or you find more information about prompt injection and what prompt defence looks like here: [prompt defence](https://medium.com/p/eadd2b993e45)

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


# Walls

The next thing to do is to build your walls. For the prompt inside your application, ask yourself if you want to use AI to
detect attacks? The advantage of this is that a lot of the work is done for you, and you can use the same AI to protect
against a range of attacks and jailbreaks. The key drawbacks at the moment are that it can be expensive, and it can be
difficult to understand what is happening inside the AI and therefore if it is working effectively. It can also have a negative
impact on performance as it adds a time consuming API call to your application.