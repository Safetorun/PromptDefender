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

The next thing to do is to build your moat. This adds the first layer of defence, looking at if PII is being sent in the
request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to bypass 
your keep defence), it can also detect if the user or session is marked as suspicious. 

# Walls

The next thing to do is to build your walls. For the prompt inside your application, ask yourself if you want to use AI to
detect attacks? The advantage of this is that a lot of the work is done for you, and you can use the same AI to protect
against a range of attacks and jailbreaks. The key drawbacks at the moment are that it can be expensive, and it can be
difficult to understand what is happening inside the AI and therefore if it is working effectively. It can also have a negative
impact on performance as it adds a time consuming API call to your application.