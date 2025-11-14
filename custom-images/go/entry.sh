#!/bin/sh

if [ -z "${AWS_LAMBDA_RUNTIME_API}" ]; then
    # Running locally with RIE
    exec /usr/local/bin/aws-lambda-rie /var/task/bootstrap
else
    # Running in AWS Lambda
    exec /var/task/bootstrap
fi