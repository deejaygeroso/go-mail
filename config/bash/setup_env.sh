#!/bin/bash

# Allow gitlab-ci.yml have access to gitlab environment variables.
echo $@

# Store gitlab environment variables to .env file.
# The .env file will then be copied over to application container.
echo EMAIL=$EMAIL >> .env
echo ENV=$ENV >> .env
echo PORT=$PORT >> .env
echo EMAIL_PASSWORD=$EMAIL_PASSWORD >> .env
