# Learnosity API GoLang Auth Client 

Very light POC to auth with the sample Learnosity question over local HTTP or AWS Lambda.

Built for the Serverless framework using a Lambda API router.

### Requirements
Working local src root with GoLang modules enabled.

Working copy of a Learnosity auth API client somewhere else if you want to validate the hash output for the auth part.


### Steps to run
1. `cp .env.example .env`
2. `make run`


You can override sessionID, userID, and timestamp in order to generate predictable outputs for the hash function. 

This allows you to modify your local code while still trying to generate the same timestamp over and over.

This is necessary because the current UTC minute is used in the hash so it that keeps changing or is generated dynamically you will always be trying to hit a moving target with your hash output / calculation.
