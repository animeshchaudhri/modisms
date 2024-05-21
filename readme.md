# Security Vulnerability: Lack of Rate Limiting on OTP Requests

## Overview

This repository demonstrates a significant security vulnerability in the OTP request functionality of the website https://phirekbaarmodisarkar.bjp.org. The issue arises due to the absence of rate limiting on the OTP requests, making the system susceptible to abuse and potential denial-of-service attacks.

## Proof of Concept

This Python script demonstrates how the lack of rate limiting can be exploited. The script sends multiple OTP requests concurrently, overwhelming the server.

## Requirements

- Python 3.x
- `requests` library
- `requests_toolbelt` library

## Installation

Install the required libraries using pip:

```bash
pip install requests requests_toolbelt
```

## Usage

```bash
git clone https://github.com/animeshchaudhri/modisms.git
cd modisms

```

## run the code

```bash
python new.py
```
