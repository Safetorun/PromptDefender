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

# Wall

The next thing to do is to build your wall. This adds the first layer of defence, looking at if PII is being sent in the
request, if the prompt itself contains "bad" words (indicating a jailbreak), or contains XML escaping (in an attempt to bypass 
your keep defence), it can also detect if the user or session is marked as suspicious. 
