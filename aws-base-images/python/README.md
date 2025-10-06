# AWS Base Images - Python

This repository creates a _very_ simple AWS Lambda function and API Gateway HTTP endpoint that uses a python3.13 base image provided by AWS.

There are two routes used by the API: `/hello` and `/goodbye`. Each route accepts an optional query parameter of `name`. The logical representation of the URL is `https://aws-api-endpoint/hello?name=Foo` or `https://aws-api-endpoint/goodbye?name=Bar`. The actual AWS API endpoint is determined when the application is deployed to AWS.

## Pre-requisites

* [mise](https://mise.jdx.dev/)

* [uv](https://docs.astral.sh/uv/)

* [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html)

## Building the Image

The mise task [create-image](.config/mise/tasks/create-image) will create the image locally.

```bash
mise run create-image
```

Run the following command to verify the image has been created:

```bash
docker images curiousdev-io/python:3.13
REPOSITORY             TAG       IMAGE ID       CREATED          SIZE
curiousdev-io/python   3.13      1420465269cc   35 minutes ago   801MB
```

## Invoking the Image Locally

The following command will build the image locally, start it with Docker, and invoke it:

```bash
mise local-build-and-invoke
```

## Deploying the Image to Amazon Elastic Container Registry (ECR)

AWS Lambda functions built using a container image must be sourced from the Amazon ECR service. The following commands demonstrate how to upload the image to ECR.

```bash
# Retrieve your account number programmatically.
# You may need to pass along the profile flag if you are not using a default AWS profile.
AWS_ACCOUNT_NUMBER=$(aws sts get-caller-identity --query Account --output text)
AWS_REGION="us-east-1" # Update to reflect the region you want to use
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_NUMBER}.dkr.ecr.${AWS_REGION}.amazonaws.com
```

Run the following command to create a repository in ECR and upload the Docker image.

```bash
mise run create-image-repo -- ${AWS_REGION}
mise run publish-image -- ${AWS_REGION}
```

**NOTE**: The mise task `create-image-repo` will only create a repo once. It will fail on subsequent runs.

## Building Your Serverless Application

You can build and deploy your AWS Lambda function once the image has been published to ECR.

```bash
mise run build-and-deploy
```
