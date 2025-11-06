# AWS OS-Only Image - Go

This repository creates a _very_ simple AWS Lambda function and API Gateway HTTP endpoint that uses a `provided.al2023` base image from AWS.

There are two routes used by the API: `/hello` and `/goodbye`. Each route accepts an optional query parameter of `name`. The logical representation of the URL is `https://aws-api-endpoint/hello?name=Foo` or `https://aws-api-endpoint/goodbye?name=Bar`. The actual AWS API endpoint is determined when the application is deployed to AWS.

## Pre-requisites

* [mise](https://mise.jdx.dev/)

## Installing Project Tools

We're using `mise` for everything we can, including setting up tooling for the project. Run the following command to install `mise` tools.

```bash
mise install
```

## Building the Image

The mise task [create-image](.config/mise/tasks/create-image) will create the image locally.

```bash
mise run create-image
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 27.3s (16/16) FINISHED                                                                                                                                                                docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                                                               0.0s
 => => transferring dockerfile: 482B                                                                                                                                                                               0.0s
 => WARN: FromPlatformFlagConstDisallowed: FROM --platform flag should not use constant value "linux/amd64" (line 2)                                                                                               0.0s
 => [internal] load metadata for public.ecr.aws/lambda/provided:al2023                                                                                                                                             0.6s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                                                                                                              1.2s
 => [internal] load .dockerignore                                                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                                                    0.0s
 => [stage-1 1/3] FROM public.ecr.aws/lambda/provided:al2023@sha256:838693f555a26743ece11c97cef4d1bb6f90b37766c9844288881da7ef14fa02                                                                               0.0s
 => => resolve public.ecr.aws/lambda/provided:al2023@sha256:838693f555a26743ece11c97cef4d1bb6f90b37766c9844288881da7ef14fa02                                                                                       0.0s
 => [build 1/7] FROM docker.io/library/golang:1.25-alpine@sha256:aee43c3ccbf24fdffb7295693b6e33b21e01baec1b2a55acc351fde345e9ec34                                                                                  2.2s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:aee43c3ccbf24fdffb7295693b6e33b21e01baec1b2a55acc351fde345e9ec34                                                                                        0.0s
 => => sha256:f3f5ae8826faeb0e0415f8f29afbc9550ae5d655f3982b2924949c93d5efd5c8 126B / 126B                                                                                                                         0.2s
 => => sha256:91631faa732ae651543f888b70295cbfe29a433d3c8da02b9966f67f238d3603 60.15MB / 60.15MB                                                                                                                   1.3s
 => => sha256:85e8836fcdb2966cd3e43a5440ccddffd1828d2d186a49fa7c17b605db8b3bb3 291.15kB / 291.15kB                                                                                                                 0.5s
 => => sha256:2d35ebdb57d9971fea0cac1582aa78935adf8058b2cc32db163c98822e5dfa1b 3.80MB / 3.80MB                                                                                                                     0.6s
 => => extracting sha256:2d35ebdb57d9971fea0cac1582aa78935adf8058b2cc32db163c98822e5dfa1b                                                                                                                          0.0s
 => => extracting sha256:85e8836fcdb2966cd3e43a5440ccddffd1828d2d186a49fa7c17b605db8b3bb3                                                                                                                          0.0s
 => => extracting sha256:91631faa732ae651543f888b70295cbfe29a433d3c8da02b9966f67f238d3603                                                                                                                          0.9s
 => => extracting sha256:f3f5ae8826faeb0e0415f8f29afbc9550ae5d655f3982b2924949c93d5efd5c8                                                                                                                          0.0s
 => => extracting sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1                                                                                                                          0.0s
 => [internal] load build context                                                                                                                                                                                  0.0s
 => => transferring context: 381B                                                                                                                                                                                  0.0s
 => [build 2/7] WORKDIR /src                                                                                                                                                                                       0.1s
 => [build 3/7] COPY go.mod go.sum ./                                                                                                                                                                              0.0s
 => [build 4/7] RUN go mod download                                                                                                                                                                                1.9s
 => [build 5/7] COPY internal/ ./internal/                                                                                                                                                                         0.0s
 => [build 6/7] COPY cmd/ ./cmd/                                                                                                                                                                                   0.0s
 => [build 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go                                                                                                    21.8s
 => CACHED [stage-1 2/3] WORKDIR /var/runtime                                                                                                                                                                      0.0s
 => CACHED [stage-1 3/3] COPY --from=build /src/bootstrap ./bootstrap                                                                                                                                              0.0s
 => exporting to image                                                                                                                                                                                             0.0s
 => => exporting layers                                                                                                                                                                                            0.0s
 => => exporting manifest sha256:441bea90348515f0b98a3c3626bc1447c766aaa05a09e296ecf7a4dd74dbdc2f                                                                                                                  0.0s
 => => exporting config sha256:dbc14264f95a57ce1e13e1be018531c7c5306e483129446612825dadf0621ce0                                                                                                                    0.0s
 => => naming to docker.io/curiousdev-io/go:1.25                                                                                                                                                                   0.0s
 => => unpacking to docker.io/curiousdev-io/go:1.25                                                                                                                                                                0.0s

 1 warning found (use docker --debug to expand):
 - FromPlatformFlagConstDisallowed: FROM --platform flag should not use constant value "linux/amd64" (line 2)
```
</details>

Verify the image has been created locally.

```bash
✗ docker images curiousdev-io/go:1.25
REPOSITORY         TAG       IMAGE ID       CREATED              SIZE
curiousdev-io/go   1.25      441bea903485   About a minute ago   193MB
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
[local-build-and-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/local-build-and-invoke
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 0.6s (16/16) FINISHED                                                                                                                                                                 docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                                                               0.0s
 => => transferring dockerfile: 482B                                                                                                                                                                               0.0s
 => WARN: FromPlatformFlagConstDisallowed: FROM --platform flag should not use constant value "linux/amd64" (line 2)                                                                                               0.0s
 => [internal] load metadata for public.ecr.aws/lambda/provided:al2023                                                                                                                                             0.6s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                                                                                                              0.4s
 => [internal] load .dockerignore                                                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                                                    0.0s
 => [build 1/7] FROM docker.io/library/golang:1.25-alpine@sha256:aee43c3ccbf24fdffb7295693b6e33b21e01baec1b2a55acc351fde345e9ec34                                                                                  0.0s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:aee43c3ccbf24fdffb7295693b6e33b21e01baec1b2a55acc351fde345e9ec34                                                                                        0.0s
 => [internal] load build context                                                                                                                                                                                  0.0s
 => => transferring context: 381B                                                                                                                                                                                  0.0s
 => [stage-1 1/3] FROM public.ecr.aws/lambda/provided:al2023@sha256:838693f555a26743ece11c97cef4d1bb6f90b37766c9844288881da7ef14fa02                                                                               0.0s
 => => resolve public.ecr.aws/lambda/provided:al2023@sha256:838693f555a26743ece11c97cef4d1bb6f90b37766c9844288881da7ef14fa02                                                                                       0.0s
 => CACHED [stage-1 2/3] WORKDIR /var/runtime                                                                                                                                                                      0.0s
 => CACHED [build 2/7] WORKDIR /src                                                                                                                                                                                0.0s
 => CACHED [build 3/7] COPY go.mod go.sum ./                                                                                                                                                                       0.0s
 => CACHED [build 4/7] RUN go mod download                                                                                                                                                                         0.0s
 => CACHED [build 5/7] COPY internal/ ./internal/                                                                                                                                                                  0.0s
 => CACHED [build 6/7] COPY cmd/ ./cmd/                                                                                                                                                                            0.0s
 => CACHED [build 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go                                                                                              0.0s
 => CACHED [stage-1 3/3] COPY --from=build /src/bootstrap ./bootstrap                                                                                                                                              0.0s
 => exporting to image                                                                                                                                                                                             0.0s
 => => exporting layers                                                                                                                                                                                            0.0s
 => => exporting manifest sha256:441bea90348515f0b98a3c3626bc1447c766aaa05a09e296ecf7a4dd74dbdc2f                                                                                                                  0.0s
 => => exporting config sha256:dbc14264f95a57ce1e13e1be018531c7c5306e483129446612825dadf0621ce0                                                                                                                    0.0s
 => => naming to docker.io/curiousdev-io/go:1.25                                                                                                                                                                   0.0s
 => => unpacking to docker.io/curiousdev-io/go:1.25                                                                                                                                                                0.0s

 1 warning found (use docker --debug to expand):
 - FromPlatformFlagConstDisallowed: FROM --platform flag should not use constant value "linux/amd64" (line 2)
[*] Docker image built successfully.
[local-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/local-invoke
[*] Starting a local instance of the function
  [!] A local instance is already running on port 9000
[*] Invoking the function locally with /hello?name=Chris
START RequestId: 20750e69-2599-4699-b8a2-e80299d08afd Version: $LATEST
05 Nov 2025 00:41:24,413 [INFO] (rapid) INVOKE START(requestId: 11b7fb55-79d6-4269-94c5-0687d8613098)
{"time":"2025-11-05T00:41:24.414353465Z","level":"INFO","msg":"lambda event received","path":"/hello","query":{"name":"Chris"}}
{"time":"2025-11-05T00:41:24.414471715Z","level":"INFO","msg":"request processed","path":"/hello","query":{"name":"Chris"},"status":200,"message":"Hello, Chris"}
{"time":"2025-11-05T00:41:24.414495673Z","level":"INFO","msg":"lambda response","status":200,"body":"{\"timestamp\":\"2025-11-05T00:41:24Z\",\"status\":200,\"message\":\"Hello, Chris\"}"}
{"statusCode":200,"headers":{"Content-Type":"application/json"},"multiValueHeaders":null,"body":"{\"timestamp\":\"2025-11-05T00:41:24Z\",\"status\":200,\"message\":\"Hello, Chris\"}","cookies":null}05 Nov 2025 00:41:24,415 [INFO] (rapid) INVOKE RTDONE(status: success, produced bytes: 0, duration: 1.565000ms)
END RequestId: 11b7fb55-79d6-4269-94c5-0687d8613098
REPORT RequestId: 11b7fb55-79d6-4269-94c5-0687d8613098  Duration: 2.07 ms       Billed Duration: 3 ms   Memory Size: 3008 MB    Max Memory Used: 3008 MB

[*] Local invocation completed successfully.
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
[create-image-repo] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/create-image-repo us-east-1
[*] Creating ECR repo: myosonlycontainerlambdafunction in region us-east-1 for account 123456789012
[*] ECR repo myosonlycontainerlambdafunction created.
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
✗ mise run build-and-deploy        
[build-and-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/build-and-deploy
[*] Running mise task: create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/create-image us-east-1
[*] Creating image...
[+] Building 0.8s (17/17) FINISHED                                                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                            0.0s
 => => transferring dockerfile: 508B                                                                                            0.0s
 => [internal] load metadata for public.ecr.aws/lambda/provided:al2023                                                          0.7s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                           0.7s
 => [internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                 0.0s
 => [build 1/8] FROM docker.io/library/golang:1.25-alpine@sha256:8b6b77a5e6a9dda591e864e1a2856d436d94219befa5f54d7ce76d2a77cc7  0.0s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:8b6b77a5e6a9dda591e864e1a2856d436d94219befa5f54d7ce76d2a77cc7a06     0.0s
 => [stage-1 1/3] FROM public.ecr.aws/lambda/provided:al2023@sha256:d3698fc80929e77245679ccd4b07fca2274e1d1916b50e0443e8e2696d  0.0s
 => => resolve public.ecr.aws/lambda/provided:al2023@sha256:d3698fc80929e77245679ccd4b07fca2274e1d1916b50e0443e8e2696d89baa1    0.0s
 => [internal] load build context                                                                                               0.0s
 => => transferring context: 381B                                                                                               0.0s
 => CACHED [build 2/8] WORKDIR /src                                                                                             0.0s
 => CACHED [build 3/8] COPY go.mod go.sum ./                                                                                    0.0s
 => CACHED [build 4/8] RUN go mod download                                                                                      0.0s
 => CACHED [build 5/8] COPY internal/ ./internal/                                                                               0.0s
 => CACHED [build 6/8] COPY cmd/ ./cmd/                                                                                         0.0s
 => CACHED [build 7/8] RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go           0.0s
 => CACHED [build 8/8] RUN chmod +x bootstrap                                                                                   0.0s
 => CACHED [stage-1 2/3] COPY --from=build /src/bootstrap /var/runtime/bootstrap                                                0.0s
 => CACHED [stage-1 3/3] RUN chmod +x /var/runtime/bootstrap                                                                    0.0s
 => exporting to image                                                                                                          0.0s
 => => exporting layers                                                                                                         0.0s
 => => exporting manifest sha256:58b859b62c55e1813bc49560bef62aac3131197e6258aa8d13409b0e8ef23134                               0.0s
 => => exporting config sha256:2d170723b662f19ad9210c032b5a500b1b89abbc9541a0736a7c255f96b7ccd0                                 0.0s
 => => naming to docker.io/curiousdev-io/go:1.25                                                                                0.0s
 => => unpacking to docker.io/curiousdev-io/go:1.25                                                                             0.0s
[*] Running mise task: publish-image
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction]
4f4fb700ef54: Pushed 
11b72ab0c7e9: Pushed 
ec7f7c5bda60: Pushed 
b1a17001d5c3: Pushed 
75131756b8ac: Pushed 
d0e8f1f8e3c2: Pushed 
1.25: digest: sha256:58b859b62c55e1813bc49560bef62aac3131197e6258aa8d13409b0e8ef23134 size: 1483
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction:1.25
[*] Running mise task: sam-build
[sam-build] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/sam-build us-east-1
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
[sam-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/sam-deploy us-east-1
[*] Updating parameter_overrides in /Users/brian/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/sam/samconfig.toml...
[*] Deploying AWS SAM stack to region us-east-1...

        Managed S3 bucket: aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Auto resolution of buckets can be turned off by setting resolve_s3=False
        To use a specific S3 bucket, set --s3-bucket=<bucket_name>
        Above settings can be stored in samconfig.toml

        Deploying with following values
        ===============================
        Stack name                   : aws-lambda-os-only-image
        Region                       : us-east-1
        Confirm changeset            : False
        Disable rollback             : False
        Deployment image repository  : 
                                       {
                                           "MyOsOnlyContainerLambdaFunction": "123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction"
                                       }
        Deployment s3 bucket         : aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Capabilities                 : ["CAPABILITY_IAM"]
        Parameter overrides          : {"EcrImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction:1.25"}
        Signing Profiles             : {}

Initiating deployment
=====================

        Uploading to aws-lambda-os-only-image/d511223d7902ecbe7011a92ad789c4ce.template  2691 / 2691  (100.00%)


Waiting for changeset to be created..

CloudFormation stack changeset
---------------------------------------------------------------------------------------------------------------------------------
Operation                        LogicalResourceId                ResourceType                     Replacement                    
---------------------------------------------------------------------------------------------------------------------------------
+ Add                            LambdaPermissionForApiGateway    AWS::Lambda::Permission          N/A                            
+ Add                            MyApiApiGatewayDefaultStage      AWS::ApiGatewayV2::Stage         N/A                            
+ Add                            MyApiLogGroup                    AWS::Logs::LogGroup              N/A                            
+ Add                            MyApi                            AWS::ApiGatewayV2::Api           N/A                            
+ Add                            MyOsOnlyContainerLambdaFunctio   AWS::Lambda::Permission          N/A                            
                                 nGoodbyeApiPermission                                                                            
+ Add                            MyOsOnlyContainerLambdaFunctio   AWS::Lambda::Permission          N/A                            
                                 nHelloApiPermission                                                                              
+ Add                            MyOsOnlyContainerLambdaFunctio   AWS::Logs::LogGroup              N/A                            
                                 nLogGroup                                                                                        
+ Add                            MyOsOnlyContainerLambdaFunctio   AWS::IAM::Role                   N/A                            
                                 nRole                                                                                            
+ Add                            MyOsOnlyContainerLambdaFunctio   AWS::Lambda::Function            N/A                            
                                 n                                                                                                
---------------------------------------------------------------------------------------------------------------------------------


Changeset created successfully. arn:aws:cloudformation:us-east-1:123456789012:changeSet/samcli-deploy1762406400/9295174f-a187-42b1-a84d-bcf13034d825


2025-11-06 00:20:07 - Waiting for stack create/update to complete

CloudFormation events from stack operations (refresh every 5.0 seconds)
---------------------------------------------------------------------------------------------------------------------------------
ResourceStatus                   ResourceType                     LogicalResourceId                ResourceStatusReason           
---------------------------------------------------------------------------------------------------------------------------------
CREATE_IN_PROGRESS               AWS::CloudFormation::Stack       aws-lambda-os-only-image         User Initiated                 
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nRole                                                           
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyOsOnlyContainerLambdaFunctio   Resource creation Initiated    
                                                                  nRole                                                           
CREATE_COMPLETE                  AWS::IAM::Role                   MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nRole                                                           
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyOsOnlyContainerLambdaFunctio   -                              
                                                                  n                                                               
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyOsOnlyContainerLambdaFunctio   Resource creation Initiated    
                                                                  n                                                               
CREATE_IN_PROGRESS -             AWS::Lambda::Function            MyOsOnlyContainerLambdaFunctio   Eventual consistency check     
CONFIGURATION_COMPLETE                                            n                                initiated                      
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nLogGroup                                                       
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    Resource creation Initiated    
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyOsOnlyContainerLambdaFunctio   Resource creation Initiated    
                                                                  nLogGroup                                                       
CREATE_COMPLETE                  AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nHelloApiPermission                                             
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nGoodbyeApiPermission                                           
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   Resource creation Initiated    
                                                                  nHelloApiPermission                                             
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   Resource creation Initiated    
                                                                  nGoodbyeApiPermission                                           
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    Resource creation Initiated    
CREATE_COMPLETE                  AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nHelloApiPermission                                             
CREATE_COMPLETE                  AWS::Lambda::Permission          MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nGoodbyeApiPermission                                           
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyOsOnlyContainerLambdaFunctio   Eventual consistency check     
CONFIGURATION_COMPLETE                                            nLogGroup                        initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyOsOnlyContainerLambdaFunctio   -                              
                                                                  nLogGroup                                                       
CREATE_COMPLETE                  AWS::Lambda::Function            MyOsOnlyContainerLambdaFunctio   -                              
                                                                  n                                                               
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyApiLogGroup                    Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_COMPLETE                  AWS::CloudFormation::Stack       aws-lambda-os-only-image         -                              
---------------------------------------------------------------------------------------------------------------------------------

CloudFormation outputs from deployed stack
----------------------------------------------------------------------------------------------------------------------------------
Outputs                                                                                                                          
----------------------------------------------------------------------------------------------------------------------------------
Key                 HelloEndpoint                                                                                                
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/hello                                                 

Key                 LambdaFunctionArn                                                                                            
Description         ARN of the Lambda function                                                                                   
Value               arn:aws:lambda:us-east-1:123456789012:function:aws-lambda-os-only-image-MyOsOnlyContainerLambdaFu-           
AXWU3LK7lEIj                                                                                                                     

Key                 GoodbyeEndpoint                                                                                              
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/goodbye                                               
----------------------------------------------------------------------------------------------------------------------------------


Successfully created/updated stack - aws-lambda-os-only-image in us-east-1

[*] Deploy complete.
[*] Build and deploy completed successfully.
```

</details>

Congrats! At this point - you now have a containerized AWS Lambda function that will respond to requests on the public internet.

## Interacting with Your API

You can open a web browser to the `HelloEndpoint` and `GoodbyeEndpoint`. Leave the query string blank to get the default values.

```bash
✗ curl https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/hello
{"timestamp":"2025-11-06T05:23:56Z","status":200,"message":"Hello, World"}%                                                          
✗ curl https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/goodbye
{"timestamp":"2025-11-06T05:24:03Z","status":200,"message":"Goodbye, World"}% 
```

Pass in the query string parameter of `name` to add some variety.

```bash
✗ curl https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/hello\?name\=Chris
{"timestamp":"2025-11-06T05:24:41Z","status":200,"message":"Hello, Chris"}%                                                          
✗ curl https://hu5ak1ibjh.execute-api.us-east-1.amazonaws.com/goodbye\?name\=Chris
{"timestamp":"2025-11-06T05:24:50Z","status":200,"message":"Goodbye, Chris"}%  
```

Alternately, you can use the task [remote-invoke](.config/mise/tasks/remote-invoke) to invoke the `/hello` and `/goodbye` routes. The task will also retrieve CloudWatch log data over the previous 10 minutes.

```bash
mise run remote-invoke -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run remote-invoke -- ${AWS_REGION}
[remote-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/remote-invoke us-east-1
[*] Invoking the MyContainerLambdaFunction in AWS (stack: aws-lambda-os-only-image)
[*] Invoking /hello endpoint with name=Brian
  [*] /hello response: {"timestamp":"2025-11-06T05:25:27Z","status":200,"message":"Hello, Brian"}
[*] Invoking /goodbye endpoint with name=Brian
  [*] /goodbye response: {"timestamp":"2025-11-06T05:25:28Z","status":200,"message":"Goodbye, Brian"}
\n[*] Latest CloudWatch log entries for MyOsOnlyContainerLambdaFunction via SAM CLI:
You can now use 'sam logs' without --name parameter, which will pull the logs from all supported resources in your stack.
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:23:56.551000+00:00 START RequestId: e979cd06-3fd1-4e48-84af-2d00e98882da Version: $LATEST
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:23:56.552000+00:00 2025/11/06 05:23:56 INFO Lambda handler invoked path=/hello method=GET
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:23:56.552000+00:00 {
  "time": "2025-11-06T05:23:56.552637069Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/hello",
  "query": null,
  "status": 200,
  "message": "Hello, World"
}
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:23:56.554000+00:00 END RequestId: e979cd06-3fd1-4e48-84af-2d00e98882da
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:23:56.554000+00:00 REPORT RequestId: e979cd06-3fd1-4e48-84af-2d00e98882da Duration: 2.39 ms       Billed Duration: 1761 ms        Memory Size: 512 MB     Max Memory Used: 26 MB  Init Duration: 1757.93 ms
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:03.651000+00:00 START RequestId: ab4f31e9-94a7-4896-bf12-0defd33e3862 Version: $LATEST
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:03.652000+00:00 2025/11/06 05:24:03 INFO Lambda handler invoked path=/goodbye method=GET
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:03.652000+00:00 {
  "time": "2025-11-06T05:24:03.652262789Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/goodbye",
  "query": null,
  "status": 200,
  "message": "Goodbye, World"
}
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:03.653000+00:00 END RequestId: ab4f31e9-94a7-4896-bf12-0defd33e3862
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:03.653000+00:00 REPORT RequestId: ab4f31e9-94a7-4896-bf12-0defd33e3862 Duration: 1.40 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 26 MB
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:41.428000+00:00 START RequestId: 457a5fbf-da14-4131-b543-2459972c789c Version: $LATEST
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:41.428000+00:00 2025/11/06 05:24:41 INFO Lambda handler invoked path=/hello method=GET
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:41.428000+00:00 {
  "time": "2025-11-06T05:24:41.428368306Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/hello",
  "query": {
    "name": "Chris"
  },
  "status": 200,
  "message": "Hello, Chris"
}
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:41.429000+00:00 END RequestId: 457a5fbf-da14-4131-b543-2459972c789c
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:41.429000+00:00 REPORT RequestId: 457a5fbf-da14-4131-b543-2459972c789c Duration: 1.22 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 26 MB
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:50.361000+00:00 START RequestId: 2545cc4c-7a58-4345-b8c4-68623c7e1a5e Version: $LATEST
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:50.361000+00:00 2025/11/06 05:24:50 INFO Lambda handler invoked path=/goodbye method=GET
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:50.361000+00:00 {
  "time": "2025-11-06T05:24:50.361573085Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/goodbye",
  "query": {
    "name": "Chris"
  },
  "status": 200,
  "message": "Goodbye, Chris"
}
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:50.362000+00:00 END RequestId: 2545cc4c-7a58-4345-b8c4-68623c7e1a5e
2025/11/06/[$LATEST]6bf4dacac1074651810a7e04d66e9885 2025-11-06T05:24:50.363000+00:00 REPORT RequestId: 2545cc4c-7a58-4345-b8c4-68623c7e1a5e Duration: 1.20 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 26 MB
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
✗ mise run cleanup -- ${AWS_REGION}
[cleanup] $ ~/code/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/.config/mise/tasks/cleanup us-east-1
[*] Stopping local Docker containers for myosonlycontainerlambdafunction:1.25...
[*] Removing local Docker image myosonlycontainerlambdafunction:1.25...
Error response from daemon: No such image: myosonlycontainerlambdafunction:1.25
[*] Removing ECR-tagged local image 123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction:1.25...
Untagged: 123456789012.dkr.ecr.us-east-1.amazonaws.com/myosonlycontainerlambdafunction:1.25
[*] Deleting ECR repository myosonlycontainerlambdafunction...
[*] Deleting CloudFormation stack aws-lambda-os-only-image...
[*] Waiting for CloudFormation stack aws-lambda-os-only-image to be deleted...
[*] Cleanup complete.
```

</details>