---
title: Wall
excerpt: Description of the Wall module, and how it works  
category: 652f7b9ea6bf5e000bb8dc94
---

Wall is intended to be used as the first line of defence against attackers. It's primary capabilities are 
in detecting jailbreak attacks based on bad words, and detecting PII (if needed).

## PII Detection 

PII Detection is done using the PII package, and there is a single implementation for this - which is the AWS PII Detector (in aws_pii).

## Bad words 

Bad words is directly inside the 'Wall' module and contains information about the bad words that are used to detect jailbreak attacks.

To generate these words, we make use of the 