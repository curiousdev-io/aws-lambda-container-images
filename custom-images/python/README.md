# Custom Images - Python

This repository creates a _very_ simple AWS Lambda function and API Gateway HTTP endpoint that uses a `python3.13` base image from [Chainguard](https://www.chainguard.dev/).

There are two routes used by the API: `/hello` and `/goodbye`. Each route accepts an optional query parameter of `name`. The logical representation of the URL is `https://aws-api-endpoint/hello?name=Foo` or `https://aws-api-endpoint/goodbye?name=Bar`. The actual AWS API endpoint is determined when the application is deployed to AWS.

**NOTE**: The technical details rely on the great work of Danilo Poccia and his post [New for AWS Lambda - Container Image Support](https://aws.amazon.com/blogs/aws/new-for-aws-lambda-container-image-support/).

## Pre-requisites

* [mise](https://mise.jdx.dev/)

* [uv](https://docs.astral.sh/uv/)

## Installing Project Tools

We're using `mise` for everything we can, including setting up tooling for the project. Run the following command to install `mise` tools.

```bash
mise install
```

<details>
<summary>Sample output</summary>

```bash
✗ mise install
mise mise 2025.9.9 by @jdx – ░░░░░░░░░░░░░░░░░░░░ 0/6 
mise mise 2025.9.9 by @jdx – ░░░░░░░░░░░░░░░░░░░░ 0/6 
mise mise 2025.9.9 by @jdx – ░░░░░░░░░░░░░░░░░░░░ 0/6 
```
</details>

## Building the Image

The mise task [create-image](.config/mise/tasks/create-image) will create the image locally. The task to run is the same as the [AWS Base Image](../../aws-base-images/python/) but the `Dockerfile` is noticably different.

* **We are using a [multi-stage Docker build](https://docs.docker.com/build/building/multi-stage/)**: The intent of running a distroless image is to minimize the footprint inside our running container. However, when building an application with an interpreted language like Python, you will likely need operating package managers (e.g. `apk`) or language package managers(e.g. `pip`). Distroless images may have `dev` versions that include such tooling. In order to take advantage of the minimal footprint we want to move our build artifacts to a runtime container free of these other dependencies. We use the [cgr.dev/chainguard/python:latest-dev](cgr.dev/chainguard/python) image to *build* our application and the [gr.dev/chainguard/python:latest](cgr.dev/chainguard/python) image to *run* our application. In our example, we copy artifacts from the `builder` stage to the `runtime` stage.

* **We add the Runtime Interface Client (RIC)**: The base AWS language container images contain something called the [Runtime Interface Client (RIC)](https://github.com/aws/aws-lambda-python-runtime-interface-client). The RIC implements the Lambda [Runtime API](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-api.html), allowing your custom image to retrieve events from the Lambda service. This is the magic that lets your application run in reponse to Events. We need to install the RIC manually.

* **We add the Runtime Interface Emulator (RIE)**: Similar to the RIC, the [Runtime Interface Emulator (RIE)](https://github.com/aws/aws-lambda-runtime-interface-emulator/) is included in AWS language container images. However, like the RIC, it is *not* included in our default container image. It is used to locally test Lambda functions packaged as a container image. Though not strictly necessary to invoke your Lambda function in AWS, you will want to add it to your custom container images so you can test locally. We use [entry.py](./entry.py) to determine whether we should invoke our functions using the RIC or RIE.

```bash
mise run create-image
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 6.2s (20/20) FINISHED                                                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                            0.0s
 => => transferring dockerfile: 1.14kB                                                                                          0.0s
 => [internal] load metadata for cgr.dev/chainguard/python:latest                                                               0.6s
 => [internal] load metadata for cgr.dev/chainguard/python:latest-dev                                                           0.6s
 => [internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                 0.0s
 => [builder 1/7] FROM cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b914  0.0s
 => => resolve cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b9142ab6c     0.0s
 => [internal] load build context                                                                                               0.0s
 => => transferring context: 1.04kB                                                                                             0.0s
 => CACHED [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download  0.4s
 => [runtime 1/6] FROM cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d987  0.0s
 => => resolve cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d9875         0.0s
 => CACHED [builder 2/7] WORKDIR /app                                                                                           0.0s
 => CACHED [runtime 2/6] WORKDIR /var/task                                                                                      0.0s
 => [builder 3/7] COPY requirements.txt .                                                                                       0.0s
 => [builder 4/7] RUN pip install --target ./packages -r requirements.txt &&     pip install --target ./packages awslambdaric   4.3s
 => [builder 5/7] COPY src/ ./src/                                                                                              0.0s 
 => [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-la  0.0s 
 => [builder 7/7] COPY --chmod=755 entry.py /app/entry.py                                                                       0.0s 
 => [runtime 3/6] COPY --from=builder /app/packages /var/task                                                                   0.2s 
 => [runtime 4/6] COPY --from=builder /app/src/ /var/task/                                                                      0.0s 
 => [runtime 5/6] COPY --from=builder /app/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                         0.0s 
 => [runtime 6/6] COPY --from=builder /app/entry.py /entry.py                                                                   0.0s
 => exporting to image                                                                                                          0.8s
 => => exporting layers                                                                                                         0.6s
 => => exporting manifest sha256:6ab5abbb683a4c3e9495f870bfc66c1e404ddeb070d19a7f0d3c0c4a77500b9f                               0.0s
 => => exporting config sha256:44b6591f542d26dd4358ace6d6eb35c885d45956afaa5a27345339aa8ea533ff                                 0.0s
 => => naming to docker.io/curiousdev-io/custom-python:3.13                                                                     0.0s
 => => unpacking to docker.io/curiousdev-io/custom-python:3.13                                                                  0.2s
```
</details>

Verify the image has been created locally.

```bash
✗ docker images curiousdev-io/custom-python:3.13
REPOSITORY                    TAG       IMAGE ID       CREATED          SIZE
curiousdev-io/custom-python   3.13      6ab5abbb683a   35 seconds ago   161MB
```

## Invoking the Image Locally

The following command will build the image locally, start it with Docker, and invoke it. Make note the `create-image` task is running as part of [local-build-and-invoke](.config/mise/tasks/local-build-and-invoke).

```bash
mise run local-build-and-invoke
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run local-build-and-invoke
[local-build-and-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/local-build-and-…
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 6.0s (20/20) FINISHED                                                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                            0.0s
 => => transferring dockerfile: 1.14kB                                                                                          0.0s
 => [internal] load metadata for cgr.dev/chainguard/python:latest                                                               0.5s
 => [internal] load metadata for cgr.dev/chainguard/python:latest-dev                                                           0.5s
 => [internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                 0.0s
 => [builder 1/7] FROM cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b914  0.0s
 => => resolve cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b9142ab6c     0.0s
 => [runtime 1/6] FROM cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d987  0.0s
 => => resolve cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d9875         0.0s
 => [internal] load build context                                                                                               0.0s
 => => transferring context: 1.04kB                                                                                             0.0s
 => CACHED [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download  0.3s
 => CACHED [runtime 2/6] WORKDIR /var/task                                                                                      0.0s
 => CACHED [builder 2/7] WORKDIR /app                                                                                           0.0s
 => [builder 3/7] COPY requirements.txt .                                                                                       0.0s
 => [builder 4/7] RUN pip install --target ./packages -r requirements.txt &&     pip install --target ./packages awslambdaric   4.2s
 => [builder 5/7] COPY src/ ./src/                                                                                              0.0s 
 => [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-la  0.0s 
 => [builder 7/7] COPY --chmod=755 entry.py /app/entry.py                                                                       0.0s 
 => [runtime 3/6] COPY --from=builder /app/packages /var/task                                                                   0.2s 
 => [runtime 4/6] COPY --from=builder /app/src/ /var/task/                                                                      0.0s 
 => [runtime 5/6] COPY --from=builder /app/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                         0.0s 
 => [runtime 6/6] COPY --from=builder /app/entry.py /entry.py                                                                   0.0s
 => exporting to image                                                                                                          0.8s
 => => exporting layers                                                                                                         0.5s
 => => exporting manifest sha256:bc19deeae8d654f9509ef38fbce79ff35f6d19f85677a4e7846385e37a4cac81                               0.0s
 => => exporting config sha256:d175670d06d60cecb5692e3c6f346ee294a5d5b144e4c7800b64a6f7318e39bd                                 0.0s
 => => naming to docker.io/curiousdev-io/custom-python:3.13                                                                     0.0s
 => => unpacking to docker.io/curiousdev-io/custom-python:3.13                                                                  0.2s
[*] Docker image built successfully.
[local-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/local-invoke
[*] Starting a local instance of the function
[*] Starting a new local instance on port 9000
09 Nov 2025 12:49:18,110 [INFO] (rapid) exec '/usr/bin/python' (cwd=/var/task, handler=lambda_function.handler)
[*] Invoking the function locally with /hello?name=Chris
09 Nov 2025 12:49:20,948 [INFO] (rapid) INIT START(type: on-demand, phase: init)
09 Nov 2025 12:49:20,950 [INFO] (rapid) The extension's directory "/opt/extensions" does not exist, assuming no extensions to be loaded.
START RequestId: eff61dc7-8fc0-4512-909a-4d902eda945a Version: $LATEST
09 Nov 2025 12:49:20,952 [INFO] (rapid) Starting runtime without AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN , Expected?: false
09 Nov 2025 12:49:21,065 [INFO] (rapid) INIT RTDONE(status: success)
09 Nov 2025 12:49:21,065 [INFO] (rapid) INIT REPORT(durationMs: 117.484000)
09 Nov 2025 12:49:21,066 [INFO] (rapid) INVOKE START(requestId: 3d5ca5e9-afb1-4d61-930d-7ec996cb1384)
{"level":"INFO","location":"handler:13","message":{"statusCode":200,"body":"{\"timestamp\": \"2025-11-09T12:49:21.068216+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}","headers":{"Content-Type":"application/json"}},"timestamp":"2025-11-09 12:49:21,068+0000","service":"service_undefined","cold_start":true,"function_name":"test_function","function_memory_size":"3008","function_arn":"arn:aws:lambda:us-east-1:012345678912:function:test_function","function_request_id":"3d5ca5e9-afb1-4d61-930d-7ec996cb1384"}
09 Nov 2025 12:49:21,071 [INFO] (rapid) INVOKE RTDONE(status: success, produced bytes: 0, duration: 4.693000ms)
END RequestId: 3d5ca5e9-afb1-4d61-930d-7ec996cb1384
REPORT RequestId: 3d5ca5e9-afb1-4d61-930d-7ec996cb1384  Init Duration: 0.65 ms  Duration: 123.16 ms     Billed Duration: 124 ms Memory Size: 3008 MB Max Memory Used: 3008 MB
{"statusCode": 200, "body": "{\"timestamp\": \"2025-11-09T12:49:21.068216+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}", "headers": {"Content-Type": "application/json"}}[*] Local invocation completed successfully.
```

</details>

## Deploying the Image to Amazon Elastic Container Registry (ECR)

AWS Lambda functions built using a container image must be sourced from the Amazon ECR service. The following commands demonstrate how to upload the image to ECR.

```bash
# Retrieve your account number programmatically.
# You may need to pass along the profile flag if you are not using a default AWS profile.
AWS_ACCOUNT_NUMBER=$(aws sts get-caller-identity --query Account --output text)
AWS_REGION="us-east-1" # Update to reflect the region you want to use
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_NUMBER}.dkr.ecr.${AWS_REGION}.amazonaws.com
```

Run the [create-image-repo](.config/mise/tasks/create-image-repo) task to create a repository in ECR.

```bash
mise run create-image-repo -- ${AWS_REGION}
```

We are following the pattern of creating one ECR repository per function.

**NOTE**: The mise task `create-image-repo` will only create a repo once. It will fail on subsequent runs.

<details>
<summary>Sample output</summary>

```bash
✗ mise run create-image-repo -- ${AWS_REGION}
[create-image-repo] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/create-image-repo us-…
[*] Creating ECR repo: mycustompythonimagelambdafunction in region us-east-1 for account 123456789012
[*] ECR repo mycustompythonimagelambdafunction created.
```
</details>

Once the repository is created, we can publish the image from our local workstation to our newly created ECR repository `mycustompythonimagelambdafunction` using the [publish-image](.config/mise/tasks/publish-image) task.

```bash
mise run publish-image -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run publish-image -- ${AWS_REGION}
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction]
392c279fac3d: Pushed 
810b6fd0587e: Pushed 
4a9af8c83319: Pushed 
d9abdc2149c9: Pushed 
2cfce6676f11: Pushed 
d520dfe61c27: Pushed 
5d640c11ac87: Pushed 
9bfe0dd71b0d: Pushed 
0ede2fb9f40c: Pushed 
3cc847a8522a: Pushed 
e4a53de388f7: Pushed 
e20defc88c56: Pushed 
b65322036506: Pushed 
afbfef10ad73: Pushed 
5ed89e6dd868: Pushed 
12ea79e160e1: Pushed 
3.13: digest: sha256:bc19deeae8d654f9509ef38fbce79ff35f6d19f85677a4e7846385e37a4cac81 size: 3457
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction:3.13
```

</details>

## Building and Deploying Your Serverless Application

You can build and deploy your AWS Lambda function using the [build-and-deploy](.config/mise/tasks/build-and-deploy) task once the image has been published to ECR. The task will build out a Serverless Application Model (SAM) serverless application. SAM is a flavor of AWS CloudFormation that makes it easier to define and deploy serverless resources.

**NOTE**: This task will do a _full_ build and deploy. Your image will be built locally and pushed to ECR before running a `sam deploy`.

The SAM template [template.yaml](sam/template.yaml) will be deployed to CloudFormation. Outputs of the CloudFormation deployment include the URLs for our `/hello` and `/goodbye` routes.

```bash
mise run build-and-deploy -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run build-and-deploy -- ${AWS_REGION}
[build-and-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/build-and-deploy us-ea…
[*] Running mise task: create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/create-image us-east-1
[*] Creating image...
[+] Building 5.8s (20/20) FINISHED                                                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                            0.0s
 => => transferring dockerfile: 1.14kB                                                                                          0.0s
 => [internal] load metadata for cgr.dev/chainguard/python:latest                                                               0.3s
 => [internal] load metadata for cgr.dev/chainguard/python:latest-dev                                                           0.3s
 => [internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                 0.0s
 => [builder 1/7] FROM cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b914  0.0s
 => => resolve cgr.dev/chainguard/python:latest-dev@sha256:b827e712d2758d10e21b44f71aed260ffb2c1e42ce695b870628ee9b9142ab6c     0.0s
 => [internal] load build context                                                                                               0.0s
 => => transferring context: 1.04kB                                                                                             0.0s
 => CACHED [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download  0.2s
 => [runtime 1/6] FROM cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d987  0.0s
 => => resolve cgr.dev/chainguard/python:latest@sha256:ce8cbd5047393b5c7d6e0f27c89f2adcfc5c2dfdcc755f85c3937f684e7d9875         0.0s
 => CACHED [builder 2/7] WORKDIR /app                                                                                           0.0s
 => CACHED [runtime 2/6] WORKDIR /var/task                                                                                      0.0s
 => [builder 3/7] COPY requirements.txt .                                                                                       0.0s
 => [builder 4/7] RUN pip install --target ./packages -r requirements.txt &&     pip install --target ./packages awslambdaric   4.2s
 => [builder 5/7] COPY src/ ./src/                                                                                              0.0s 
 => [builder 6/7] ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-la  0.0s 
 => [builder 7/7] COPY --chmod=755 entry.py /app/entry.py                                                                       0.0s 
 => [runtime 3/6] COPY --from=builder /app/packages /var/task                                                                   0.2s 
 => [runtime 4/6] COPY --from=builder /app/src/ /var/task/                                                                      0.0s 
 => [runtime 5/6] COPY --from=builder /app/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                         0.0s 
 => [runtime 6/6] COPY --from=builder /app/entry.py /entry.py                                                                   0.0s
 => exporting to image                                                                                                          0.8s
 => => exporting layers                                                                                                         0.6s
 => => exporting manifest sha256:cf32426fca8369cbf291208dead0496effddd62cf4c79ff2dbe9b0cb1b1eeb49                               0.0s
 => => exporting config sha256:18d0bdc3cdb8c2cbaecd298f5b5620a96feba9b4092032fae1d0b72e179780e0                                 0.0s
 => => naming to docker.io/curiousdev-io/custom-python:3.13                                                                     0.0s
 => => unpacking to docker.io/curiousdev-io/custom-python:3.13                                                                  0.2s
[*] Running mise task: publish-image
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction]
3cc847a8522a: Layer already exists 
b65322036506: Layer already exists 
3d6164ecaaab: Pushed 
392c279fac3d: Layer already exists 
9bfe0dd71b0d: Layer already exists 
d520dfe61c27: Layer already exists 
d9abdc2149c9: Layer already exists 
afbfef10ad73: Layer already exists 
810b6fd0587e: Layer already exists 
2cfce6676f11: Layer already exists 
4a9af8c83319: Layer already exists 
0ede2fb9f40c: Layer already exists 
e20defc88c56: Layer already exists 
e81863895374: Pushed 
5d640c11ac87: Layer already exists 
12ea79e160e1: Layer already exists 
3.13: digest: sha256:cf32426fca8369cbf291208dead0496effddd62cf4c79ff2dbe9b0cb1b1eeb49 size: 3457
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction:3.13
[*] Running mise task: sam-build
[sam-build] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/sam-build us-east-1
Building AWS SAM stack...

Build Succeeded

Built Artifacts  : sam/.aws-sam/build
Built Template   : sam/.aws-sam/build/template.yaml

Commands you can use next
=========================
[*] Validate SAM template: sam validate
[*] Invoke Function: sam local invoke -t sam/.aws-sam/build/template.yaml
[*] Test Function in the Cloud: sam sync --stack-name {{stack-name}} --watch
[*] Deploy: sam deploy --guided --template-file sam/.aws-sam/build/template.yaml
Build complete.
[*] Running mise task: sam-deploy
[sam-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/sam-deploy us-east-1
[*] Updating parameter_overrides in /Users/brian/code/curiousdev-io/aws-lambda-container-images/custom-images/python/sam/samconfig.toml...
[*] Deploying AWS SAM stack to region us-east-1...

        Managed S3 bucket: aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Auto resolution of buckets can be turned off by setting resolve_s3=False
        To use a specific S3 bucket, set --s3-bucket=<bucket_name>
        Above settings can be stored in samconfig.toml

        Deploying with following values
        ===============================
        Stack name                   : aws-custom-image-python
        Region                       : us-east-1
        Confirm changeset            : False
        Disable rollback             : False
        Deployment image repository  : 
                                       {
                                           "MyCustomPythonImageLambdaFunction": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction"
                                       }
        Deployment s3 bucket         : aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Capabilities                 : ["CAPABILITY_IAM"]
        Parameter overrides          : {"EcrImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction:3.13"}
        Signing Profiles             : {}

Initiating deployment
=====================

        Uploading to aws-custom-image-python/ea84fb3f645a6d2881ccd6a7287dc114.template  2705 / 2705  (100.00%)


Waiting for changeset to be created..

CloudFormation stack changeset
---------------------------------------------------------------------------------------------------------------------------------
Operation                        LogicalResourceId                ResourceType                     Replacement                    
---------------------------------------------------------------------------------------------------------------------------------
+ Add                            LambdaPermissionForApiGateway    AWS::Lambda::Permission          N/A                            
+ Add                            MyApiApiGatewayDefaultStage      AWS::ApiGatewayV2::Stage         N/A                            
+ Add                            MyApiLogGroup                    AWS::Logs::LogGroup              N/A                            
+ Add                            MyApi                            AWS::ApiGatewayV2::Api           N/A                            
+ Add                            MyCustomPythonImageLambdaFunct   AWS::Lambda::Permission          N/A                            
                                 ionGoodbyeApiPermission                                                                          
+ Add                            MyCustomPythonImageLambdaFunct   AWS::Lambda::Permission          N/A                            
                                 ionHelloApiPermission                                                                            
+ Add                            MyCustomPythonImageLambdaFunct   AWS::Logs::LogGroup              N/A                            
                                 ionLogGroup                                                                                      
+ Add                            MyCustomPythonImageLambdaFunct   AWS::IAM::Role                   N/A                            
                                 ionRole                                                                                          
+ Add                            MyCustomPythonImageLambdaFunct   AWS::Lambda::Function            N/A                            
                                 ion                                                                                              
---------------------------------------------------------------------------------------------------------------------------------


Changeset created successfully. arn:aws:cloudformation:us-east-1:123456789012:changeSet/samcli-deploy1762694688/3e061be5-dd87-42b0-a012-c5bbd3afa304


2025-11-09 08:24:54 - Waiting for stack create/update to complete

CloudFormation events from stack operations (refresh every 5.0 seconds)
---------------------------------------------------------------------------------------------------------------------------------
ResourceStatus                   ResourceType                     LogicalResourceId                ResourceStatusReason           
---------------------------------------------------------------------------------------------------------------------------------
CREATE_IN_PROGRESS               AWS::CloudFormation::Stack       aws-custom-image-python          User Initiated                 
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyCustomPythonImageLambdaFunct   -                              
                                                                  ionRole                                                         
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyCustomPythonImageLambdaFunct   Resource creation Initiated    
                                                                  ionRole                                                         
CREATE_COMPLETE                  AWS::IAM::Role                   MyCustomPythonImageLambdaFunct   -                              
                                                                  ionRole                                                         
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyCustomPythonImageLambdaFunct   -                              
                                                                  ion                                                             
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyCustomPythonImageLambdaFunct   Resource creation Initiated    
                                                                  ion                                                             
CREATE_IN_PROGRESS -             AWS::Lambda::Function            MyCustomPythonImageLambdaFunct   Eventual consistency check     
CONFIGURATION_COMPLETE                                            ion                              initiated                      
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyCustomPythonImageLambdaFunct   -                              
                                                                  ionLogGroup                                                     
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyCustomPythonImageLambdaFunct   Resource creation Initiated    
                                                                  ionLogGroup                                                     
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    Resource creation Initiated    
CREATE_COMPLETE                  AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   -                              
                                                                  ionHelloApiPermission                                           
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   -                              
                                                                  ionGoodbyeApiPermission                                         
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   Resource creation Initiated    
                                                                  ionGoodbyeApiPermission                                         
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   Resource creation Initiated    
                                                                  ionHelloApiPermission                                           
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    Resource creation Initiated    
CREATE_COMPLETE                  AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   -                              
                                                                  ionHelloApiPermission                                           
CREATE_COMPLETE                  AWS::Lambda::Permission          MyCustomPythonImageLambdaFunct   -                              
                                                                  ionGoodbyeApiPermission                                         
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyCustomPythonImageLambdaFunct   Eventual consistency check     
CONFIGURATION_COMPLETE                                            ionLogGroup                      initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyCustomPythonImageLambdaFunct   -                              
                                                                  ionLogGroup                                                     
CREATE_COMPLETE                  AWS::Lambda::Function            MyCustomPythonImageLambdaFunct   -                              
                                                                  ion                                                             
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyApiLogGroup                    Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_COMPLETE                  AWS::CloudFormation::Stack       aws-custom-image-python          -                              
---------------------------------------------------------------------------------------------------------------------------------

CloudFormation outputs from deployed stack
----------------------------------------------------------------------------------------------------------------------------------
Outputs                                                                                                                          
----------------------------------------------------------------------------------------------------------------------------------
Key                 HelloEndpoint                                                                                                
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/hello                                                 

Key                 LambdaFunctionArn                                                                                            
Description         ARN of the Lambda function                                                                                   
Value               arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-                                      
MyCustomPythonImageLambdaF-Y6WzvDrzdyWk                                                                                          

Key                 GoodbyeEndpoint                                                                                              
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/goodbye                                               
----------------------------------------------------------------------------------------------------------------------------------


Successfully created/updated stack - aws-custom-image-python in us-east-1

[*] Deploy complete.
[*] Build and deploy completed successfully.
```

</details>

Congrats! At this point - you now have a containerized AWS Lambda function that will respond to requests on the public internet.

## Interacting with Your API

You can open a web browser to the `HelloEndpoint` and `GoodbyeEndpoint`. Leave the query string blank to get the default values.

```bash
✗ curl https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/hello
{"timestamp": "2025-11-09T13:26:53.498036+00:00", "status": 200, "message": "Hello, World"}%                                         
✗ curl https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/goodbye
{"timestamp": "2025-11-09T13:27:03.681967+00:00", "status": 200, "message": "Goodbye, World"}% 
```

Pass in the query string parameter of `name` to add some variety.

✗ curl https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/hello\?name\=Chris
{"timestamp": "2025-11-09T13:28:07.732541+00:00", "status": 200, "message": "Hello, Chris"}%                                         
✗ curl https://rw5pgapdea.execute-api.us-east-1.amazonaws.com/goodbye\?name\=Chris
{"timestamp": "2025-11-09T13:28:19.533108+00:00", "status": 200, "message": "Goodbye, Chris"}%   

Alternately, you can use the task [remote-invoke](.config/mise/tasks/remote-invoke) to invoke the `/hello` and `/goodbye` routes. The task will also retrieve CloudWatch log data over the previous 10 minutes.

```bash
mise run remote-invoke -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run remote-invoke -- ${AWS_REGION}
[remote-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/python/.config/mise/tasks/remote-invoke us-east-1
[*] Invoking the MyCustomPythonImageLambdaFunction in AWS (stack: aws-custom-image-python)
[*] Invoking /hello endpoint with name=Brian
  [*] /hello response: {"timestamp": "2025-11-09T13:30:44.326551+00:00", "status": 200, "message": "Hello, Brian"}
[*] Invoking /goodbye endpoint with name=Brian
  [*] /goodbye response: {"timestamp": "2025-11-09T13:30:45.002436+00:00", "status": 200, "message": "Goodbye, Brian"}
\n[*] Latest CloudWatch log entries for MyCustomPythonImageLambdaFunction via SAM CLI:
You can now use 'sam logs' without --name parameter, which will pull the logs from all supported resources in your stack.
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:26:53.497000+00:00 START RequestId: 6c0536fe-64c7-4e75-a806-f6700553c8d8 Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:26:53.498000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:26:53.498036+00:00\", \"status\": 200, \"message\": \"Hello, World\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:26:53,498+0000",
  "service": "service_undefined",
  "cold_start": true,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "6c0536fe-64c7-4e75-a806-f6700553c8d8",
  "xray_trace_id": "1-69109698-2581a2bf2a3bac0c5c357416"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:26:53.499000+00:00 END RequestId: 6c0536fe-64c7-4e75-a806-f6700553c8d8
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:26:53.499000+00:00 REPORT RequestId: 6c0536fe-64c7-4e75-a806-f6700553c8d8 Duration: 2.13 ms       Billed Duration: 4213 ms        Memory Size: 512 MB     Max Memory Used: 41 MB  Init Duration: 4210.85 ms
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:27:03.681000+00:00 START RequestId: 4a15724a-732b-4b7a-98ab-366901027974 Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:27:03.682000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:27:03.681967+00:00\", \"status\": 200, \"message\": \"Goodbye, World\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:27:03,682+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "4a15724a-732b-4b7a-98ab-366901027974",
  "xray_trace_id": "1-691096a7-1ee4a6c31584c06b0173f326"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:27:03.683000+00:00 END RequestId: 4a15724a-732b-4b7a-98ab-366901027974
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:27:03.683000+00:00 REPORT RequestId: 4a15724a-732b-4b7a-98ab-366901027974 Duration: 1.37 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 43 MB
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:07.732000+00:00 START RequestId: 813c16f9-793e-4b24-abfb-800df39cb41a Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:07.732000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:28:07.732541+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:28:07,732+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "813c16f9-793e-4b24-abfb-800df39cb41a",
  "xray_trace_id": "1-691096e7-230f44f3059e7e2503891647"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:07.733000+00:00 END RequestId: 813c16f9-793e-4b24-abfb-800df39cb41a
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:07.733000+00:00 REPORT RequestId: 813c16f9-793e-4b24-abfb-800df39cb41a Duration: 1.36 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 43 MB
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:19.532000+00:00 START RequestId: bedd7db5-7051-4c59-954c-0a3beb31d738 Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:19.533000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:28:19.533108+00:00\", \"status\": 200, \"message\": \"Goodbye, Chris\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:28:19,533+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "bedd7db5-7051-4c59-954c-0a3beb31d738",
  "xray_trace_id": "1-691096f3-6a3214733e6e649b11f04bcf"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:19.534000+00:00 END RequestId: bedd7db5-7051-4c59-954c-0a3beb31d738
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:28:19.534000+00:00 REPORT RequestId: bedd7db5-7051-4c59-954c-0a3beb31d738 Duration: 1.34 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 43 MB
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.309000+00:00 START RequestId: 58d144ef-8f3d-4769-bb3e-a592c2684989 Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.310000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:29:14.310114+00:00\", \"status\": 200, \"message\": \"Hello, Brian\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:29:14,310+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "58d144ef-8f3d-4769-bb3e-a592c2684989",
  "xray_trace_id": "1-6910972a-4641eab601f80a7e4d0a925d"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.311000+00:00 END RequestId: 58d144ef-8f3d-4769-bb3e-a592c2684989
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.311000+00:00 REPORT RequestId: 58d144ef-8f3d-4769-bb3e-a592c2684989 Duration: 1.39 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 43 MB
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.954000+00:00 START RequestId: b568f1c9-c5f1-4c9b-b12b-180cbb75a615 Version: $LATEST
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.955000+00:00 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-11-09T13:29:14.955246+00:00\", \"status\": 200, \"message\": \"Goodbye, Brian\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-11-09 13:29:14,955+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-python-MyCustomPythonImageLambdaF-Y6WzvDrzdyWk",
  "function_request_id": "b568f1c9-c5f1-4c9b-b12b-180cbb75a615",
  "xray_trace_id": "1-6910972a-2f6ce19e450431403e0f232d"
}
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.956000+00:00 END RequestId: b568f1c9-c5f1-4c9b-b12b-180cbb75a615
2025/11/09/[$LATEST]1587ef42317b43cfa7a26b7ee349ce37 2025-11-09T13:29:14.956000+00:00 REPORT RequestId: b568f1c9-c5f1-4c9b-b12b-180cbb75a615 Duration: 1.35 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 43 MB
```
</details>

## Cleaning Up

If you were following along then both local and AWS resources have been created. To remove them, run the [cleanup](.config/mise/tasks/cleanup) task.

```bash
mise run cleanup -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run cleanup -- ${AWS_REGION}
[cleanup] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/cleanup us-east-1
[*] Stopping local Docker containers for curiousdev-io/custom-python:3.13...
[*] Removing local Docker image curiousdev-io/custom-python:3.13...
Untagged: curiousdev-io/custom-python:3.13
Deleted: sha256:cf32426fca8369cbf291208dead0496effddd62cf4c79ff2dbe9b0cb1b1eeb49
[*] Removing ECR-tagged local image 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction:3.13...
Untagged: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustompythonimagelambdafunction:3.13
[*] Deleting ECR repository mycustompythonimagelambdafunction...
[*] Deleting CloudFormation stack aws-custom-image-python...
[*] Waiting for CloudFormation stack aws-custom-image-python to be deleted...
[*] Cleanup complete.
```

</details>