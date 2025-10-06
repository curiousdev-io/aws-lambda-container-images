# AWS Base Images - Python

This repository creates a _very_ simple AWS Lambda function and API Gateway HTTP endpoint that uses a python3.13 base image provided by AWS.

There are two routes used by the API: `/hello` and `/goodbye`. Each route accepts an optional query parameter of `name`. The logical representation of the URL is `https://aws-api-endpoint/hello?name=Foo` or `https://aws-api-endpoint/goodbye?name=Bar`. The actual AWS API endpoint is determined when the application is deployed to AWS.

## Pre-requisites

* [mise](https://mise.jdx.dev/)

* [uv](https://docs.astral.sh/uv/)

* [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html)

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

The mise task [create-image](.config/mise/tasks/create-image) will create the image locally.

```bash
mise run create-image
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 0.5s (9/9) FINISHED                                                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                           0.0s
 => => transferring dockerfile: 382B                                                                                           0.0s
 => [internal] load metadata for public.ecr.aws/lambda/python:3.13                                                             0.5s
 => [internal] load .dockerignore                                                                                              0.0s
 => => transferring context: 2B                                                                                                0.0s
 => [1/4] FROM public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => => resolve public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => [internal] load build context                                                                                              0.0s
 => => transferring context: 1.02kB                                                                                            0.0s
 => CACHED [2/4] COPY requirements.txt /var/task                                                                               0.0s
 => CACHED [3/4] RUN pip install -r requirements.txt                                                                           0.0s
 => CACHED [4/4] COPY src /var/task                                                                                            0.0s
 => exporting to image                                                                                                         0.0s
 => => exporting layers                                                                                                        0.0s
 => => exporting manifest sha256:1420465269ccc88167382c795c1c0111692eefa43c81911e3a10ff76b4461c9e                              0.0s
 => => exporting config sha256:1c329197f1d8c06247571889ec61e7e6c80533ddd299c0c0b4b8e87122e2f2e6                                0.0s
 => => naming to docker.io/curiousdev-io/python:3.13                                                                           0.0s
 => => unpacking to docker.io/curiousdev-io/python:3.13    
```
</details>

Verify the image has been created locally.

```bash
✗ docker images curiousdev-io/python:3.13
REPOSITORY             TAG       IMAGE ID       CREATED          SIZE
curiousdev-io/python   3.13      1420465269cc   35 minutes ago   801MB
```

## Invoking the Image Locally

The following command will build the image locally, start it with Docker, and invoke it. Make note the `create-image` task is running as part of [local-build-and-invoke](.config/mise/tasks/local-build-and-invoke).

```bash
mise local-build-and-invoke
```

<details>
<summary>Sample output</summary>

```bash
✗ mise local-build-and-invoke
[local-build-and-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/local-build-a…
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 0.8s (9/9) FINISHED                                                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                           0.0s
 => => transferring dockerfile: 382B                                                                                           0.0s
 => [internal] load metadata for public.ecr.aws/lambda/python:3.13                                                             0.8s
 => [internal] load .dockerignore                                                                                              0.0s
 => => transferring context: 2B                                                                                                0.0s
 => [1/4] FROM public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => => resolve public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => [internal] load build context                                                                                              0.0s
 => => transferring context: 1.02kB                                                                                            0.0s
 => CACHED [2/4] COPY requirements.txt /var/task                                                                               0.0s
 => CACHED [3/4] RUN pip install -r requirements.txt                                                                           0.0s
 => CACHED [4/4] COPY src /var/task                                                                                            0.0s
 => exporting to image                                                                                                         0.0s
 => => exporting layers                                                                                                        0.0s
 => => exporting manifest sha256:1420465269ccc88167382c795c1c0111692eefa43c81911e3a10ff76b4461c9e                              0.0s
 => => exporting config sha256:1c329197f1d8c06247571889ec61e7e6c80533ddd299c0c0b4b8e87122e2f2e6                                0.0s
 => => naming to docker.io/curiousdev-io/python:3.13                                                                           0.0s
 => => unpacking to docker.io/curiousdev-io/python:3.13                                                                        0.0s
[*] Docker image built successfully.
[local-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/local-invoke
[*] Starting a local instance of the function
[*] Starting a new local instance on port 9000
06 Oct 2025 11:09:53,513 [INFO] (rapid) exec '/var/runtime/bootstrap' (cwd=/var/task, handler=)
[*] Invoking the function locally with /hello?name=Chris
START RequestId: 655c6f9e-7ca6-48cc-87d8-1cf0f1e04b90 Version: $LATEST
06 Oct 2025 11:09:56,452 [INFO] (rapid) INIT START(type: on-demand, phase: init)
06 Oct 2025 11:09:56,452 [INFO] (rapid) The extension's directory "/opt/extensions" does not exist, assuming no extensions to be loaded.
06 Oct 2025 11:09:56,452 [INFO] (rapid) Starting runtime without AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN , Expected?: false
06 Oct 2025 11:09:56,489 [INFO] (rapid) INIT RTDONE(status: success)
06 Oct 2025 11:09:56,489 [INFO] (rapid) INIT REPORT(durationMs: 37.204000)
06 Oct 2025 11:09:56,489 [INFO] (rapid) INVOKE START(requestId: 6ee905fa-1cdc-4e5c-ad1d-67493b0d5e0b)
{"level":"INFO","location":"handler:13","message":{"statusCode":200,"body":"{\"timestamp\": \"2025-10-06T11:09:56.489521+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}","headers":{"Content-Type":"application/json"}},"timestamp":"2025-10-06 11:09:56,489+0000","service":"service_undefined","cold_start":true,"function_name":"test_function","function_memory_size":"3008","function_arn":"arn:aws:lambda:us-east-1:012345678912:function:test_function","function_request_id":"6ee905fa-1cdc-4e5c-ad1d-67493b0d5e0b"}
06 Oct 2025 11:09:56,489 [INFO] (rapid) INVOKE RTDONE(status: success, produced bytes: 0, duration: 0.657000ms)
END RequestId: 6ee905fa-1cdc-4e5c-ad1d-67493b0d5e0b
REPORT RequestId: 6ee905fa-1cdc-4e5c-ad1d-67493b0d5e0b  Init Duration: 0.02 ms  Duration: 37.96 ms      Billed Duration: 38 ms  Memory Size: 3008 MB        Max Memory Used: 3008 MB
{"statusCode": 200, "body": "{\"timestamp\": \"2025-10-06T11:09:56.489521+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}", "headers": {"Content-Type": "application/json"}}[*] Local invocation completed successfully.
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
[create-image-repo] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/create-image-repo …
[*] Creating ECR repo: mycontainerlambdafunction in region us-east-1 for account 123456789012
[*] ECR repo mycontainerlambdafunction created.
```
</details>

Once the repository is created, we can publish the image from our local workstation to our newly created ECR repository `mycontainerlambdafunction` using the [publish-image](.config/mise/tasks/publish-image) task.

```bash
mise run publish-image -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run publish-image -- ${AWS_REGION}
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction]
98cf4f1d4db0: Pushed 
75131756b8ac: Pushed 
15b3e90c2f6f: Pushed 
5a50b1adf161: Pushed 
4549d54d1e1a: Pushed 
3967d372135e: Pushed 
57ea51c444f1: Pushed 
c6d65b6f4c67: Pushed 
e14db851bfe4: Pushed 
3.13: digest: sha256:1420465269ccc88167382c795c1c0111692eefa43c81911e3a10ff76b4461c9e size: 2075
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction:3.13
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
[build-and-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/build-and-deploy us…
[*] Running mise task: create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/create-image us-east-1
[*] Creating image...
[+] Building 0.8s (9/9) FINISHED                                                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                           0.0s
 => => transferring dockerfile: 382B                                                                                           0.0s
 => [internal] load metadata for public.ecr.aws/lambda/python:3.13                                                             0.8s
 => [internal] load .dockerignore                                                                                              0.0s
 => => transferring context: 2B                                                                                                0.0s
 => [1/4] FROM public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => => resolve public.ecr.aws/lambda/python:3.13@sha256:bd53508accbaeb0977c5c358b772cbc088a3b324523e43f191f2ddb784ea544d       0.0s
 => [internal] load build context                                                                                              0.0s
 => => transferring context: 1.02kB                                                                                            0.0s
 => CACHED [2/4] COPY requirements.txt /var/task                                                                               0.0s
 => CACHED [3/4] RUN pip install -r requirements.txt                                                                           0.0s
 => CACHED [4/4] COPY src /var/task                                                                                            0.0s
 => exporting to image                                                                                                         0.0s
 => => exporting layers                                                                                                        0.0s
 => => exporting manifest sha256:1420465269ccc88167382c795c1c0111692eefa43c81911e3a10ff76b4461c9e                              0.0s
 => => exporting config sha256:1c329197f1d8c06247571889ec61e7e6c80533ddd299c0c0b4b8e87122e2f2e6                                0.0s
 => => naming to docker.io/curiousdev-io/python:3.13                                                                           0.0s
 => => unpacking to docker.io/curiousdev-io/python:3.13                                                                        0.0s
[*] Running mise task: publish-image
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction]
75131756b8ac: Layer already exists 
5a50b1adf161: Layer already exists 
c6d65b6f4c67: Layer already exists 
98cf4f1d4db0: Layer already exists 
57ea51c444f1: Layer already exists 
15b3e90c2f6f: Layer already exists 
3967d372135e: Layer already exists 
4549d54d1e1a: Layer already exists 
e14db851bfe4: Layer already exists 
3.13: digest: sha256:1420465269ccc88167382c795c1c0111692eefa43c81911e3a10ff76b4461c9e size: 2075
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction:3.13
[*] Running mise task: sam-build
[sam-build] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/sam-build us-east-1
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
[sam-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/sam-deploy us-east-1
[*] Updating parameter_overrides in /Users/brian/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/sam/samconfig.toml...
[*] Deploying AWS SAM stack to region us-east-1...

        Managed S3 bucket: aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Auto resolution of buckets can be turned off by setting resolve_s3=False
        To use a specific S3 bucket, set --s3-bucket=<bucket_name>
        Above settings can be stored in samconfig.toml

        Deploying with following values
        ===============================
        Stack name                   : aws-base-image-python
        Region                       : us-east-1
        Confirm changeset            : False
        Disable rollback             : False
        Deployment image repository  : 
                                       {
                                           "MyContainerLambdaFunction": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction"
                                       }
        Deployment s3 bucket         : aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Capabilities                 : ["CAPABILITY_IAM"]
        Parameter overrides          : {"EcrImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction:3.13"}
        Signing Profiles             : {}

Initiating deployment
=====================

        Uploading to aws-base-image-python/b26162222dad4588506f96b393ccf98d.template  2649 / 2649  (100.00%)


Waiting for changeset to be created..

CloudFormation stack changeset
---------------------------------------------------------------------------------------------------------------------------------
Operation                        LogicalResourceId                ResourceType                     Replacement                    
---------------------------------------------------------------------------------------------------------------------------------
+ Add                            LambdaPermissionForApiGateway    AWS::Lambda::Permission          N/A                            
+ Add                            MyApiApiGatewayDefaultStage      AWS::ApiGatewayV2::Stage         N/A                            
+ Add                            MyApiLogGroup                    AWS::Logs::LogGroup              N/A                            
+ Add                            MyApi                            AWS::ApiGatewayV2::Api           N/A                            
+ Add                            MyContainerLambdaFunctionGoodb   AWS::Lambda::Permission          N/A                            
                                 yeApiPermission                                                                                  
+ Add                            MyContainerLambdaFunctionHello   AWS::Lambda::Permission          N/A                            
                                 ApiPermission                                                                                    
+ Add                            MyContainerLambdaFunctionLogGr   AWS::Logs::LogGroup              N/A                            
                                 oup                                                                                              
+ Add                            MyContainerLambdaFunctionRole    AWS::IAM::Role                   N/A                            
+ Add                            MyContainerLambdaFunction        AWS::Lambda::Function            N/A                            
---------------------------------------------------------------------------------------------------------------------------------


Changeset created successfully. arn:aws:cloudformation:us-east-1:123456789012:changeSet/samcli-deploy1759750760/28a717b9-0c04-435f-a972-4d623fd26ad8


2025-10-06 07:39:27 - Waiting for stack create/update to complete

CloudFormation events from stack operations (refresh every 5.0 seconds)
---------------------------------------------------------------------------------------------------------------------------------
ResourceStatus                   ResourceType                     LogicalResourceId                ResourceStatusReason           
---------------------------------------------------------------------------------------------------------------------------------
CREATE_IN_PROGRESS               AWS::CloudFormation::Stack       aws-base-image-python            User Initiated                 
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyContainerLambdaFunctionRole    -                              
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyContainerLambdaFunctionRole    Resource creation Initiated    
CREATE_COMPLETE                  AWS::IAM::Role                   MyContainerLambdaFunctionRole    -                              
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyContainerLambdaFunction        -                              
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyContainerLambdaFunction        Resource creation Initiated    
CREATE_IN_PROGRESS -             AWS::Lambda::Function            MyContainerLambdaFunction        Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyContainerLambdaFunctionLogGr   -                              
                                                                  oup                                                             
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    Resource creation Initiated    
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyContainerLambdaFunctionLogGr   Resource creation Initiated    
                                                                  oup                                                             
CREATE_COMPLETE                  AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyContainerLambdaFunctionHello   -                              
                                                                  ApiPermission                                                   
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyContainerLambdaFunctionGoodb   -                              
                                                                  yeApiPermission                                                 
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    Resource creation Initiated    
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyContainerLambdaFunctionHello   Resource creation Initiated    
                                                                  ApiPermission                                                   
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyContainerLambdaFunctionGoodb   Resource creation Initiated    
                                                                  yeApiPermission                                                 
CREATE_COMPLETE                  AWS::Lambda::Permission          MyContainerLambdaFunctionHello   -                              
                                                                  ApiPermission                                                   
CREATE_COMPLETE                  AWS::Lambda::Permission          MyContainerLambdaFunctionGoodb   -                              
                                                                  yeApiPermission                                                 
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyContainerLambdaFunctionLogGr   Eventual consistency check     
CONFIGURATION_COMPLETE                                            oup                              initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyContainerLambdaFunctionLogGr   -                              
                                                                  oup                                                             
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyApiLogGroup                    Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_COMPLETE                  AWS::Lambda::Function            MyContainerLambdaFunction        -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_COMPLETE                  AWS::CloudFormation::Stack       aws-base-image-python            -                              
---------------------------------------------------------------------------------------------------------------------------------

CloudFormation outputs from deployed stack
---------------------------------------------------------------------------------------------------------------------------------
Outputs
---------------------------------------------------------------------------------------------------------------------------------
Key                 HelloEndpoint
Description         URL of the HTTP API endpoint (no stage prefix)
Value               https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/hello

Key                 LambdaFunctionArn
Description         ARN of the Lambda function
Value               arn:aws:lambda:us-east-1:408023262302:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI 

Key                 GoodbyeEndpoint
Description         URL of the HTTP API endpoint (no stage prefix) 
Value               https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/goodbye                                              
---------------------------------------------------------------------------------------------------------------------------------


Successfully created/updated stack - aws-base-image-python in us-east-1

[*] Deploy complete.
[*] Build and deploy completed successfully.
```

</details>

Congrats! At this point - you now have a containerized AWS Lambda function that will respond to requests on the public internet.

## Interacting with Your API

You can open a web browser to the `HelloEndpoint` and `GoodbyeEndpoint`. Leave the query string blank to get the default values.

```bash
✗ curl https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/hello  
{"timestamp": "2025-10-06T11:48:30.629330+00:00", "status": 200, "message": "Hello, World"}%                                        
➜  python git:(main) ✗ curl https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/goodbye
{"timestamp": "2025-10-06T11:48:35.031821+00:00", "status": 200, "message": "Goodbye, World"}% 
```

Pass in the query string parameter of `name` to add some variety.

✗ curl https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/hello\?name\=Chris
{"timestamp": "2025-10-06T11:49:19.531466+00:00", "status": 200, "message": "Hello, Chris"}%                                        
➜  python git:(main) ✗ curl https://iouibh3kyg.execute-api.us-east-1.amazonaws.com/goodbye\?name\=Chris
{"timestamp": "2025-10-06T11:49:27.582635+00:00", "status": 200, "message": "Goodbye, Chris"}%     

Alternately, you can use the task [remote-invoke](.config/mise/tasks/remote-invoke) to invoke the `/hello` and `/goodbye` routes. The task will also retrieve CloudWatch log data over the previous 10 minutes.

```bash
mise run remote-invoke -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run remote-invoke -- ${AWS_REGION}
[remote-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/remote-invoke us-east-1
[*] Invoking the MyContainerLambdaFunction in AWS (stack: aws-base-image-python)
[*] Invoking /hello endpoint with name=Brian
  [*] /hello response: {"timestamp": "2025-10-06T11:59:17.687322+00:00", "status": 200, "message": "Hello, Brian"}
[*] Invoking /goodbye endpoint with name=Brian
  [*] /goodbye response: {"timestamp": "2025-10-06T11:59:18.339764+00:00", "status": 200, "message": "Goodbye, Brian"}
\n[*] Latest CloudWatch log entries for MyContainerLambdaFunction via SAM CLI:
You can now use 'sam logs' without --name parameter, which will pull the logs from all supported resources in your stack.
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:19.530000 START RequestId: 1e333318-a652-42ab-b341-0020008bae72 Version: $LATEST
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:19.531000 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-10-06T11:49:19.531466+00:00\", \"status\": 200, \"message\": \"Hello, Chris\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-10-06 11:49:19,531+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_request_id": "1e333318-a652-42ab-b341-0020008bae72",
  "xray_trace_id": "1-68e3acbf-1f24ceb9385954a62fe8ea4e"
}
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:19.533000 END RequestId: 1e333318-a652-42ab-b341-0020008bae72
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:19.533000 REPORT RequestId: 1e333318-a652-42ab-b341-0020008bae72      Duration: 2.50 ms       Billed Duration: 3 ms   Memory Size: 512 MB     Max Memory Used: 47 MB
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:27.582000 START RequestId: a07ceba9-f6a5-4689-afea-ef148c17df5b Version: $LATEST
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:27.582000 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-10-06T11:49:27.582635+00:00\", \"status\": 200, \"message\": \"Goodbye, Chris\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-10-06 11:49:27,582+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_request_id": "a07ceba9-f6a5-4689-afea-ef148c17df5b",
  "xray_trace_id": "1-68e3acc7-079a262e6ee4be20671e22e1"
}
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:27.584000 END RequestId: a07ceba9-f6a5-4689-afea-ef148c17df5b
2025/10/06/[$LATEST]6d2c579a60d34c1dacd8d192947e13dc 2025-10-06T11:49:27.584000 REPORT RequestId: a07ceba9-f6a5-4689-afea-ef148c17df5b      Duration: 2.10 ms       Billed Duration: 3 ms   Memory Size: 512 MB     Max Memory Used: 48 MB
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:24.373000 START RequestId: 46402e1c-0341-4626-8648-ad926de48595 Version: $LATEST
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:24.374000 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-10-06T11:55:24.374232+00:00\", \"status\": 200, \"message\": \"Hello, Brian\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-10-06 11:55:24,374+0000",
  "service": "service_undefined",
  "cold_start": true,
  "function_name": "aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_request_id": "46402e1c-0341-4626-8648-ad926de48595",
  "xray_trace_id": "1-68e3ae2b-0bf685633cbb43ad4784e3fe"
}
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:24.376000 END RequestId: 46402e1c-0341-4626-8648-ad926de48595
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:24.376000 REPORT RequestId: 46402e1c-0341-4626-8648-ad926de48595      Duration: 2.69 ms       Billed Duration: 203 ms Memory Size: 512 MB     Max Memory Used: 44 MB  Init Duration: 199.53 ms
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:58.938000 START RequestId: fa49c93e-a128-4c57-912c-1fe89ccc2016 Version: $LATEST
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:58.939000 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-10-06T11:55:58.938780+00:00\", \"status\": 200, \"message\": \"Hello, Brian\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-10-06 11:55:58,938+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_request_id": "fa49c93e-a128-4c57-912c-1fe89ccc2016",
  "xray_trace_id": "1-68e3ae4e-33b53b04095a5c807e7efcbd"
}
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:58.940000 END RequestId: fa49c93e-a128-4c57-912c-1fe89ccc2016
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:58.940000 REPORT RequestId: fa49c93e-a128-4c57-912c-1fe89ccc2016      Duration: 1.88 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 44 MB
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:59.610000 START RequestId: e13b1942-7357-4eaf-9623-7488f7921246 Version: $LATEST
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:59.611000 {
  "level": "INFO",
  "location": "handler:13",
  "message": {
    "statusCode": 200,
    "body": "{\"timestamp\": \"2025-10-06T11:55:59.610775+00:00\", \"status\": 200, \"message\": \"Goodbye, Brian\"}",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "timestamp": "2025-10-06 11:55:59,610+0000",
  "service": "service_undefined",
  "cold_start": false,
  "function_name": "aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_memory_size": "512",
  "function_arn": "arn:aws:lambda:us-east-1:123456789012:function:aws-base-image-python-MyContainerLambdaFunction-gryiKyoeEsrI",
  "function_request_id": "e13b1942-7357-4eaf-9623-7488f7921246",
  "xray_trace_id": "1-68e3ae4f-149cdf3a689ed0cc29c5d382"
}
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:59.612000 END RequestId: e13b1942-7357-4eaf-9623-7488f7921246
2025/10/06/[$LATEST]6fab6527a9ab413687b6cc835c80e337 2025-10-06T11:55:59.612000 REPORT RequestId: e13b1942-7357-4eaf-9623-7488f7921246      Duration: 1.91 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 44 MB
```
</details>

## Cleaning Up

If you were following along then both local and AWS resources have been created. To remove them, run the [cleanup](.config/mise/tasks/cleanup) task.

```bash
mise run cleanup -- ${AWS_REGION}
```

<detail>
<summary><summary>

```bash
✗ mise run cleanup -- ${AWS_REGION}
[cleanup] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-base-images/python/.config/mise/tasks/cleanup us-east-1
[*] Stopping local Docker containers for curiousdev-io/python:3.13...
06 Oct 2025 12:04:59,450 [INFO] (rapid) Received signal signal=terminated
06 Oct 2025 12:04:59,450 [INFO] (rapid) Shutting down...
06 Oct 2025 12:04:59,450 [WARNING] (rapid) Reset initiated: SandboxTerminated
06 Oct 2025 12:04:59,450 [INFO] (rapid) Sending SIGKILL to runtime-1(17).
06 Oct 2025 12:04:59,450 [INFO] (rapid) Waiting for runtime domain processes termination
70ec7ba0af95
[*] Removing local Docker image curiousdev-io/python:3.13...
Untagged: curiousdev-io/python:3.13
[*] Removing ECR-tagged local image 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction:3.13...
Untagged: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycontainerlambdafunction:3.13
[*] Deleting ECR repository mycontainerlambdafunction...
[*] Deleting CloudFormation stack aws-base-image-python...
[*] Waiting for CloudFormation stack aws-base-image-python to be deleted...
[*] Cleanup complete.
```

</detail>