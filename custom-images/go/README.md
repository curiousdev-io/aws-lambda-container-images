# Custom Images - Go

This repository creates a _very_ simple AWS Lambda function and API Gateway HTTP endpoint that uses a base [golang:1.25-alpine](https://hub.docker.com/_/golang/) image in the **build** stage and a [alpine:3.22](https://hub.docker.com/_/alpine) image in the **runtime** stage.

There are two routes used by the API: `/hello` and `/goodbye`. Each route accepts an optional query parameter of `name`. The logical representation of the URL is `https://aws-api-endpoint/hello?name=Foo` or `https://aws-api-endpoint/goodbye?name=Bar`. The actual AWS API endpoint is determined when the application is deployed to AWS.

**NOTE**: The technical details rely on the great work of Danilo Poccia and his post [New for AWS Lambda - Container Image Support](https://aws.amazon.com/blogs/aws/new-for-aws-lambda-container-image-support/).

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
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 8.5s (26/26) FINISHED                                                                                                                                                                                                                                      docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                                                    0.0s
 => => transferring dockerfile: 953B                                                                                                                                                                                                                                                    0.0s
 => [internal] load metadata for docker.io/library/alpine:3.22                                                                                                                                                                                                                          0.6s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                                                                                                                                                                                   0.6s
 => [internal] load .dockerignore                                                                                                                                                                                                                                                       0.0s
 => => transferring context: 2B                                                                                                                                                                                                                                                         0.0s
 => CACHED [build  1/11] FROM docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d831bbb                                                                                                                                              0.0s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d831bbb                                                                                                                                                             0.0s
 => [internal] load build context                                                                                                                                                                                                                                                       0.0s
 => => transferring context: 618B                                                                                                                                                                                                                                                       0.0s
 => [runtime 1/8] FROM docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412                                                                                                                                                            0.0s
 => => resolve docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412                                                                                                                                                                    0.0s
 => CACHED [build 10/11] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tmp/aws-lambda-rie                                                                                                                                   0.6s
 => [build  2/11] RUN apk add --no-cache python3 py3-pip                                                                                                                                                                                                                                2.2s
 => [build  3/11] WORKDIR /src                                                                                                                                                                                                                                                          0.0s 
 => [build  4/11] COPY go.mod go.sum ./                                                                                                                                                                                                                                                 0.0s 
 => [build  5/11] RUN go mod download                                                                                                                                                                                                                                                   0.9s 
 => [build  6/11] COPY internal/ ./internal/                                                                                                                                                                                                                                            0.0s 
 => [build  7/11] COPY cmd/ ./cmd/                                                                                                                                                                                                                                                      0.0s 
 => [build  8/11] RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go                                                                                                                                                                        3.8s 
 => [build  9/11] RUN chmod +x bootstrap                                                                                                                                                                                                                                                0.1s
 => [build 10/11] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tmp/aws-lambda-rie                                                                                                                                          0.0s
 => [build 11/11] RUN chmod +x /tmp/aws-lambda-rie                                                                                                                                                                                                                                      0.1s
 => CACHED [runtime 2/8] WORKDIR /var/task                                                                                                                                                                                                                                              0.0s
 => CACHED [runtime 3/8] COPY --from=build /src/bootstrap /var/task                                                                                                                                                                                                                     0.0s
 => [runtime 4/8] COPY --from=build /tmp/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                                                                                                                                                                                   0.0s
 => [runtime 5/8] COPY entry.sh /entry.sh                                                                                                                                                                                                                                               0.0s
 => [runtime 6/8] RUN chmod +x /usr/local/bin/aws-lambda-rie                                                                                                                                                                                                                            0.1s
 => [runtime 7/8] RUN chmod +x bootstrap                                                                                                                                                                                                                                                0.1s
 => [runtime 8/8] RUN chmod +x /entry.sh                                                                                                                                                                                                                                                0.1s
 => exporting to image                                                                                                                                                                                                                                                                  0.2s
 => => exporting layers                                                                                                                                                                                                                                                                 0.2s
 => => exporting manifest sha256:18e6b16f1a8907dbd205524a50ee0603a5fb95aa6d00bf5c14450fb31049d02d                                                                                                                                                                                       0.0s
 => => exporting config sha256:2c4b9879494722116646048326bbc41f3f06e29b36c3d6c083a22d4b9a0b8b6a                                                                                                                                                                                         0.0s
 => => naming to docker.io/curiousdev-io/custom-go:1.25                                                                                                                                                                                                                                 0.0s
 => => unpacking to docker.io/curiousdev-io/custom-go:1.25                                                                                                                                                                                                                              0.0s
```
</details>

Verify the image has been created locally.

```bash
✗ docker images curiousdev-io/custom-go:1.25
REPOSITORY                TAG       IMAGE ID       CREATED          SIZE
curiousdev-io/custom-go   1.25      18e6b16f1a89   40 seconds ago   37.6MB
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
[local-build-and-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/local-build-and-invoke
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/create-image
[*] Creating image...
[+] Building 0.4s (28/28) FINISHED                                                                                                                                                                                                                                      docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                                                    0.0s
 => => transferring dockerfile: 1.03kB                                                                                                                                                                                                                                                  0.0s
 => [internal] load metadata for docker.io/library/alpine:3.22                                                                                                                                                                                                                          0.2s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                                                                                                                                                                                   0.1s
 => [internal] load .dockerignore                                                                                                                                                                                                                                                       0.0s
 => => transferring context: 2B                                                                                                                                                                                                                                                         0.0s
 => [build  1/13] FROM docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d831bbb                                                                                                                                                     0.0s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d831bbb                                                                                                                                                             0.0s
 => [build 10/13] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tmp/aws-lambda-rie                                                                                                                                          0.2s
 => [runtime 1/8] FROM docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412                                                                                                                                                            0.0s
 => => resolve docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412                                                                                                                                                                    0.0s
 => [internal] load build context                                                                                                                                                                                                                                                       0.0s
 => => transferring context: 409B                                                                                                                                                                                                                                                       0.0s
 => CACHED [runtime 2/8] WORKDIR /var/task                                                                                                                                                                                                                                              0.0s
 => CACHED [build  2/13] RUN apk add --no-cache python3 py3-pip                                                                                                                                                                                                                         0.0s
 => CACHED [build  3/13] WORKDIR /src                                                                                                                                                                                                                                                   0.0s
 => CACHED [build  4/13] COPY go.mod go.sum ./                                                                                                                                                                                                                                          0.0s
 => CACHED [build  5/13] RUN go mod download                                                                                                                                                                                                                                            0.0s
 => CACHED [build  6/13] COPY internal/ ./internal/                                                                                                                                                                                                                                     0.0s
 => CACHED [build  7/13] COPY cmd/ ./cmd/                                                                                                                                                                                                                                               0.0s
 => CACHED [build  8/13] RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go                                                                                                                                                                 0.0s
 => CACHED [build  9/13] RUN chmod +x bootstrap                                                                                                                                                                                                                                         0.0s
 => CACHED [build 10/13] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tmp/aws-lambda-rie                                                                                                                                   0.0s
 => CACHED [build 11/13] RUN chmod +x /tmp/aws-lambda-rie                                                                                                                                                                                                                               0.0s
 => CACHED [build 12/13] COPY entry.sh /tmp/entry.sh                                                                                                                                                                                                                                    0.0s
 => CACHED [build 13/13] RUN chmod +x /tmp/entry.sh                                                                                                                                                                                                                                     0.0s
 => CACHED [runtime 3/8] COPY --from=build /src/bootstrap /var/task                                                                                                                                                                                                                     0.0s
 => CACHED [runtime 4/8] COPY --from=build /tmp/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                                                                                                                                                                            0.0s
 => CACHED [runtime 5/8] COPY --from=build /tmp/entry.sh /entry.sh                                                                                                                                                                                                                      0.0s
 => CACHED [runtime 6/8] RUN chmod +x /usr/local/bin/aws-lambda-rie                                                                                                                                                                                                                     0.0s
 => CACHED [runtime 7/8] RUN chmod +x bootstrap                                                                                                                                                                                                                                         0.0s
 => CACHED [runtime 8/8] RUN chmod +x /entry.sh                                                                                                                                                                                                                                         0.0s
 => exporting to image                                                                                                                                                                                                                                                                  0.0s
 => => exporting layers                                                                                                                                                                                                                                                                 0.0s
 => => exporting manifest sha256:19d2e0d77d765c2804100083b7c0ba5da5b2e5890b289f0000423d1404af5b7f                                                                                                                                                                                       0.0s
 => => exporting config sha256:1836c5cfba3726a95943c20610b0fce5229def66ecf3b3d88ac708d7a3eff2a4                                                                                                                                                                                         0.0s
 => => naming to docker.io/curiousdev-io/custom-go:1.25                                                                                                                                                                                                                                 0.0s
 => => unpacking to docker.io/curiousdev-io/custom-go:1.25                                                                                                                                                                                                                              0.0s
[*] Docker image built successfully.
[local-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/local-invoke
[*] Starting a local instance of the function
[*] Starting a new local instance on port 9000
14 Nov 2025 01:26:37,154 [INFO] (rapid) exec '/var/task/bootstrap' (cwd=/var/task, handler=)
[*] Invoking the function locally with /hello?name=Chris
14 Nov 2025 01:26:40,043 [INFO] (rapid) INIT START(type: on-demand, phase: init)
14 Nov 2025 01:26:40,044 [INFO] (rapid) The extension's directory "/opt/extensions" does not exist, assuming no extensions to be loaded.
START RequestId: 5438d4b0-a21c-46c0-8309-5dc50ff10372 Version: $LATEST
14 Nov 2025 01:26:40,047 [INFO] (rapid) Starting runtime without AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN , Expected?: false
14 Nov 2025 01:26:40,065 [INFO] (rapid) INIT RTDONE(status: success)
14 Nov 2025 01:26:40,065 [INFO] (rapid) INIT REPORT(durationMs: 23.160000)
14 Nov 2025 01:26:40,066 [INFO] (rapid) INVOKE START(requestId: da277603-b1f8-493c-9f2f-056c70330ad1)
2025/11/14 01:26:40 INFO Lambda handler invoked path=/hello method=GET
{"time":"2025-11-14T01:26:40.068674013Z","level":"INFO","msg":"request processed","path":"/hello","query":{"name":"Chris"},"status":200,"message":"Hello, Chris"}
14 Nov 2025 01:26:40,071 [INFO] (rapid) INVOKE RTDONE(status: success, produced bytes: 0, duration: 5.295000ms)
END RequestId: da277603-b1f8-493c-9f2f-056c70330ad1
REPORT RequestId: da277603-b1f8-493c-9f2f-056c70330ad1  Init Duration: 0.42 ms  Duration: 29.32 ms      Billed Duration: 30 ms  Memory Size: 3008 MB    Max Memory Used: 3008 MB
{"statusCode":200,"headers":{"Content-Type":"application/json"},"multiValueHeaders":null,"body":"{\"timestamp\":\"2025-11-14T01:26:40Z\",\"status\":200,\"message\":\"Hello, Chris\"}","cookies":null}
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
[create-image-repo] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/create-image-repo us-east…
[*] Creating ECR repo: mycustomgoimagelambdafunction in region us-east-1 for account 123456789012
[*] ECR repo mycustomgoimagelambdafunction created.
```
</details>

Once the repository is created, we can publish the image from our local workstation to our newly created ECR repository `mycustomgoimagelambdafunction` using the [publish-image](.config/mise/tasks/publish-image) task.

```bash
mise run publish-image -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run publish-image -- ${AWS_REGION}
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction]
4f4fb700ef54: Pushed 
cebaf02df8c8: Pushed 
d5938dc1ac89: Pushed 
0710e66e3c6b: Pushed 
6b59a28fa201: Pushed 
b0b975ca786f: Pushed 
1.25: digest: sha256:19d2e0d77d765c2804100083b7c0ba5da5b2e5890b289f0000423d1404af5b7f size: 1866
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction:1.25
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
[build-and-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/build-and-deploy us-east-1
[*] Running mise task: create-image
[create-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/create-image us-east-1
[*] Creating image...
[+] Building 1.2s (28/28) FINISHED                                                                              docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                            0.0s
 => => transferring dockerfile: 1.03kB                                                                                          0.0s
 => [internal] load metadata for docker.io/library/alpine:3.22                                                                  0.7s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                           0.7s
 => [internal] load .dockerignore                                                                                               0.0s
 => => transferring context: 2B                                                                                                 0.0s
 => [build  1/13] FROM docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d8  0.0s
 => => resolve docker.io/library/golang:1.25-alpine@sha256:d3f0cf7723f3429e3f9ed846243970b20a2de7bae6a5b66fc5914e228d831bbb     0.0s
 => [internal] load build context                                                                                               0.0s
 => => transferring context: 409B                                                                                               0.0s
 => [runtime 1/8] FROM docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412    0.0s
 => => resolve docker.io/library/alpine:3.22@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412            0.0s
 => [build 10/13] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tm  0.5s
 => CACHED [runtime 2/8] WORKDIR /var/task                                                                                      0.0s
 => CACHED [build  2/13] RUN apk add --no-cache python3 py3-pip                                                                 0.0s
 => CACHED [build  3/13] WORKDIR /src                                                                                           0.0s
 => CACHED [build  4/13] COPY go.mod go.sum ./                                                                                  0.0s
 => CACHED [build  5/13] RUN go mod download                                                                                    0.0s
 => CACHED [build  6/13] COPY internal/ ./internal/                                                                             0.0s
 => CACHED [build  7/13] COPY cmd/ ./cmd/                                                                                       0.0s
 => CACHED [build  8/13] RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go         0.0s
 => CACHED [build  9/13] RUN chmod +x bootstrap                                                                                 0.0s
 => CACHED [build 10/13] ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-  0.0s
 => CACHED [build 11/13] RUN chmod +x /tmp/aws-lambda-rie                                                                       0.0s
 => CACHED [build 12/13] COPY entry.sh /tmp/entry.sh                                                                            0.0s
 => CACHED [build 13/13] RUN chmod +x /tmp/entry.sh                                                                             0.0s
 => CACHED [runtime 3/8] COPY --from=build /src/bootstrap /var/task                                                             0.0s
 => CACHED [runtime 4/8] COPY --from=build /tmp/aws-lambda-rie /usr/local/bin/aws-lambda-rie                                    0.0s
 => CACHED [runtime 5/8] COPY --from=build /tmp/entry.sh /entry.sh                                                              0.0s
 => CACHED [runtime 6/8] RUN chmod +x /usr/local/bin/aws-lambda-rie                                                             0.0s
 => CACHED [runtime 7/8] RUN chmod +x bootstrap                                                                                 0.0s
 => CACHED [runtime 8/8] RUN chmod +x /entry.sh                                                                                 0.0s
 => exporting to image                                                                                                          0.0s
 => => exporting layers                                                                                                         0.0s
 => => exporting manifest sha256:19d2e0d77d765c2804100083b7c0ba5da5b2e5890b289f0000423d1404af5b7f                               0.0s
 => => exporting config sha256:1836c5cfba3726a95943c20610b0fce5229def66ecf3b3d88ac708d7a3eff2a4                                 0.0s
 => => naming to docker.io/curiousdev-io/custom-go:1.25                                                                         0.0s
 => => unpacking to docker.io/curiousdev-io/custom-go:1.25                                                                      0.0s
[*] Running mise task: publish-image
[publish-image] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/publish-image us-east-1
[*] Publishing image to AWS ECR in region us-east-1
Login Succeeded
The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction]
0710e66e3c6b: Layer already exists 
cebaf02df8c8: Layer already exists 
4f4fb700ef54: Layer already exists 
6b59a28fa201: Layer already exists 
b0b975ca786f: Layer already exists 
d5938dc1ac89: Layer already exists 
1.25: digest: sha256:19d2e0d77d765c2804100083b7c0ba5da5b2e5890b289f0000423d1404af5b7f size: 1866
[*] Image published: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction:1.25
[*] Running mise task: sam-build
[sam-build] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/sam-build us-east-1
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
[sam-deploy] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/sam-deploy us-east-1
[*] Updating parameter_overrides in /Users/brian/code/curiousdev-io/aws-lambda-container-images/custom-images/go/sam/samconfig.toml...
[*] Deploying AWS SAM stack to region us-east-1...

        Managed S3 bucket: aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Auto resolution of buckets can be turned off by setting resolve_s3=False
        To use a specific S3 bucket, set --s3-bucket=<bucket_name>
        Above settings can be stored in samconfig.toml

        Deploying with following values
        ===============================
        Stack name                   : aws-custom-image-go
        Region                       : us-east-1
        Confirm changeset            : False
        Disable rollback             : False
        Deployment image repository  : 
                                       {
                                           "MyCustomGoImageLambdaFunction": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction"
                                       }
        Deployment s3 bucket         : aws-sam-cli-managed-default-samclisourcebucket-1n87uhj5iy5vy
        Capabilities                 : ["CAPABILITY_IAM"]
        Parameter overrides          : {"EcrImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction:1.25"}
        Signing Profiles             : {}

Initiating deployment
=====================

        Uploading to aws-custom-image-go/1cd16603ced8a8646a2bffe8f6bc446c.template  2677 / 2677  (100.00%)


Waiting for changeset to be created..

CloudFormation stack changeset
---------------------------------------------------------------------------------------------------------------------------------
Operation                        LogicalResourceId                ResourceType                     Replacement                    
---------------------------------------------------------------------------------------------------------------------------------
+ Add                            LambdaPermissionForApiGateway    AWS::Lambda::Permission          N/A                            
+ Add                            MyApiApiGatewayDefaultStage      AWS::ApiGatewayV2::Stage         N/A                            
+ Add                            MyApiLogGroup                    AWS::Logs::LogGroup              N/A                            
+ Add                            MyApi                            AWS::ApiGatewayV2::Api           N/A                            
+ Add                            MyCustomGoImageLambdaFunctionG   AWS::Lambda::Permission          N/A                            
                                 oodbyeApiPermission                                                                              
+ Add                            MyCustomGoImageLambdaFunctionH   AWS::Lambda::Permission          N/A                            
                                 elloApiPermission                                                                                
+ Add                            MyCustomGoImageLambdaFunctionL   AWS::Logs::LogGroup              N/A                            
                                 ogGroup                                                                                          
+ Add                            MyCustomGoImageLambdaFunctionR   AWS::IAM::Role                   N/A                            
                                 ole                                                                                              
+ Add                            MyCustomGoImageLambdaFunction    AWS::Lambda::Function            N/A                            
---------------------------------------------------------------------------------------------------------------------------------


Changeset created successfully. arn:aws:cloudformation:us-east-1:123456789012:changeSet/samcli-deploy1763086775/b393986c-fc1e-4a5f-8c64-b214c0440566


2025-11-13 21:19:42 - Waiting for stack create/update to complete

CloudFormation events from stack operations (refresh every 5.0 seconds)
---------------------------------------------------------------------------------------------------------------------------------
ResourceStatus                   ResourceType                     LogicalResourceId                ResourceStatusReason           
---------------------------------------------------------------------------------------------------------------------------------
CREATE_IN_PROGRESS               AWS::CloudFormation::Stack       aws-custom-image-go              User Initiated                 
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyCustomGoImageLambdaFunctionR   -                              
                                                                  ole                                                             
CREATE_IN_PROGRESS               AWS::IAM::Role                   MyCustomGoImageLambdaFunctionR   Resource creation Initiated    
                                                                  ole                                                             
CREATE_COMPLETE                  AWS::IAM::Role                   MyCustomGoImageLambdaFunctionR   -                              
                                                                  ole                                                             
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyCustomGoImageLambdaFunction    -                              
CREATE_IN_PROGRESS               AWS::Lambda::Function            MyCustomGoImageLambdaFunction    Resource creation Initiated    
CREATE_IN_PROGRESS -             AWS::Lambda::Function            MyCustomGoImageLambdaFunction    Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyCustomGoImageLambdaFunctionL   -                              
                                                                  ogGroup                                                         
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyCustomGoImageLambdaFunctionL   Resource creation Initiated    
                                                                  ogGroup                                                         
CREATE_IN_PROGRESS               AWS::Lambda::Permission          LambdaPermissionForApiGateway    Resource creation Initiated    
CREATE_COMPLETE                  AWS::Lambda::Permission          LambdaPermissionForApiGateway    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Api           MyApi                            Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Api           MyApi                            -                              
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionH   -                              
                                                                  elloApiPermission                                               
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionG   -                              
                                                                  oodbyeApiPermission                                             
CREATE_IN_PROGRESS               AWS::Logs::LogGroup              MyApiLogGroup                    Resource creation Initiated    
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionH   Resource creation Initiated    
                                                                  elloApiPermission                                               
CREATE_IN_PROGRESS               AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionG   Resource creation Initiated    
                                                                  oodbyeApiPermission                                             
CREATE_COMPLETE                  AWS::Lambda::Function            MyCustomGoImageLambdaFunction    -                              
CREATE_COMPLETE                  AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionH   -                              
                                                                  elloApiPermission                                               
CREATE_COMPLETE                  AWS::Lambda::Permission          MyCustomGoImageLambdaFunctionG   -                              
                                                                  oodbyeApiPermission                                             
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyCustomGoImageLambdaFunctionL   Eventual consistency check     
CONFIGURATION_COMPLETE                                            ogGroup                          initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyCustomGoImageLambdaFunctionL   -                              
                                                                  ogGroup                                                         
CREATE_IN_PROGRESS -             AWS::Logs::LogGroup              MyApiLogGroup                    Eventual consistency check     
CONFIGURATION_COMPLETE                                                                             initiated                      
CREATE_COMPLETE                  AWS::Logs::LogGroup              MyApiLogGroup                    -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_IN_PROGRESS               AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      Resource creation Initiated    
CREATE_COMPLETE                  AWS::ApiGatewayV2::Stage         MyApiApiGatewayDefaultStage      -                              
CREATE_COMPLETE                  AWS::CloudFormation::Stack       aws-custom-image-go              -                              
---------------------------------------------------------------------------------------------------------------------------------

CloudFormation outputs from deployed stack
----------------------------------------------------------------------------------------------------------------------------------
Outputs                                                                                                                          
----------------------------------------------------------------------------------------------------------------------------------
Key                 HelloEndpoint                                                                                                
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://z7mo376j24.execute-api.us-east-1.amazonaws.com/hello                                                 

Key                 LambdaFunctionArn                                                                                            
Description         ARN of the Lambda function                                                                                   
Value               arn:aws:lambda:us-east-1:123456789012:function:aws-custom-image-go-MyCustomGoImageLambdaFunction-            
eItSiiXmxgnh                                                                                                                     

Key                 GoodbyeEndpoint                                                                                              
Description         URL of the HTTP API endpoint (no stage prefix)                                                               
Value               https://z7mo376j24.execute-api.us-east-1.amazonaws.com/goodbye                                               
----------------------------------------------------------------------------------------------------------------------------------


Successfully created/updated stack - aws-custom-image-go in us-east-1

[*] Deploy complete.
[*] Build and deploy completed successfully.
```

</details>

Congrats! At this point - you now have a containerized AWS Lambda function that will respond to requests on the public internet.

## Interacting with Your API

You can open a web browser to the `HelloEndpoint` and `GoodbyeEndpoint`. Leave the query string blank to get the default values.

```bash
✗ curl https://z7mo376j24.execute-api.us-east-1.amazonaws.com/hello
{"timestamp":"2025-11-14T02:24:14Z","status":200,"message":"Hello, World"}%                                                      
✗ curl https://z7mo376j24.execute-api.us-east-1.amazonaws.com/goodbye
{"timestamp":"2025-11-14T02:24:22Z","status":200,"message":"Goodbye, World"}% 
```

Pass in the query string parameter of `name` to add some variety.

```bash
✗ curl https://z7mo376j24.execute-api.us-east-1.amazonaws.com/hello\?name\=Chris
{"timestamp":"2025-11-14T02:26:30Z","status":200,"message":"Hello, Chris"}%                                                         
✗ curl https://z7mo376j24.execute-api.us-east-1.amazonaws.com/goodbye\?name\=Chris
{"timestamp":"2025-11-14T02:26:41Z","status":200,"message":"Goodbye, Chris"}%  
```

Alternately, you can use the task [remote-invoke](.config/mise/tasks/remote-invoke) to invoke the `/hello` and `/goodbye` routes. The task will also retrieve CloudWatch log data over the previous 10 minutes.

```bash
mise run remote-invoke -- ${AWS_REGION}
```

<details>
<summary>Sample output</summary>

```bash
✗ mise run remote-invoke -- ${AWS_REGION}
[remote-invoke] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/remote-invoke us-east-1
[*] Invoking the MyCustomGoImageLambdaFunction in AWS (stack: aws-custom-image-go)
[*] Invoking /hello endpoint with name=Brian
  [*] /hello response: {"timestamp":"2025-11-14T02:32:00Z","status":200,"message":"Hello, Brian"}
[*] Invoking /goodbye endpoint with name=Brian
  [*] /goodbye response: {"timestamp":"2025-11-14T02:32:01Z","status":200,"message":"Goodbye, Brian"}
\n[*] Latest CloudWatch log entries for MyCustomGoImageLambdaFunction via SAM CLI:
You can now use 'sam logs' without --name parameter, which will pull the logs from all supported resources in your stack.
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:14.855000+00:00 START RequestId: 8b345087-ba6d-4419-8efd-9f2b3e2f27df Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:14.857000+00:00 2025/11/14 02:24:14 INFO Lambda handler invoked path=/hello method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:14.857000+00:00 {
  "time": "2025-11-14T02:24:14.857694742Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/hello",
  "query": null,
  "status": 200,
  "message": "Hello, World"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:14.859000+00:00 END RequestId: 8b345087-ba6d-4419-8efd-9f2b3e2f27df
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:14.859000+00:00 REPORT RequestId: 8b345087-ba6d-4419-8efd-9f2b3e2f27df Duration: 4.38 ms       Billed Duration: 1822 ms        Memory Size: 512 MB     Max Memory Used: 28 MB  Init Duration: 1816.66 ms
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:22.844000+00:00 START RequestId: 89c68597-3f62-4dae-8a02-f449a7e8aed4 Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:22.844000+00:00 2025/11/14 02:24:22 INFO Lambda handler invoked path=/goodbye method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:22.844000+00:00 {
  "time": "2025-11-14T02:24:22.844299593Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/goodbye",
  "query": null,
  "status": 200,
  "message": "Goodbye, World"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:22.845000+00:00 END RequestId: 89c68597-3f62-4dae-8a02-f449a7e8aed4
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:24:22.845000+00:00 REPORT RequestId: 89c68597-3f62-4dae-8a02-f449a7e8aed4 Duration: 1.16 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 28 MB
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:30.002000+00:00 START RequestId: 54cc8193-ec41-4383-8fb6-717064b9cb32 Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:30.003000+00:00 2025/11/14 02:26:30 INFO Lambda handler invoked path=/hello method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:30.003000+00:00 {
  "time": "2025-11-14T02:26:30.003082768Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/hello",
  "query": {
    "name": "Chris"
  },
  "status": 200,
  "message": "Hello, Chris"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:30.003000+00:00 END RequestId: 54cc8193-ec41-4383-8fb6-717064b9cb32
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:30.003000+00:00 REPORT RequestId: 54cc8193-ec41-4383-8fb6-717064b9cb32 Duration: 1.05 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 29 MB
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:41.364000+00:00 START RequestId: 7b5183d9-0395-4099-9461-46106cbe708d Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:41.364000+00:00 2025/11/14 02:26:41 INFO Lambda handler invoked path=/goodbye method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:41.364000+00:00 {
  "time": "2025-11-14T02:26:41.364885047Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/goodbye",
  "query": {
    "name": "Chris"
  },
  "status": 200,
  "message": "Goodbye, Chris"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:41.365000+00:00 END RequestId: 7b5183d9-0395-4099-9461-46106cbe708d
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:26:41.365000+00:00 REPORT RequestId: 7b5183d9-0395-4099-9461-46106cbe708d Duration: 1.04 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 29 MB
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.103000+00:00 START RequestId: a9548284-dc33-4258-b047-e7b038f0f1d8 Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.104000+00:00 2025/11/14 02:27:53 INFO Lambda handler invoked path=/hello method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.104000+00:00 {
  "time": "2025-11-14T02:27:53.104067065Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/hello",
  "query": {
    "name": "Brian"
  },
  "status": 200,
  "message": "Hello, Brian"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.104000+00:00 END RequestId: a9548284-dc33-4258-b047-e7b038f0f1d8
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.104000+00:00 REPORT RequestId: a9548284-dc33-4258-b047-e7b038f0f1d8 Duration: 1.15 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 30 MB
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.740000+00:00 START RequestId: 5ab9e771-8e87-4a73-9d04-191893d49874 Version: $LATEST
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.740000+00:00 2025/11/14 02:27:53 INFO Lambda handler invoked path=/goodbye method=GET
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.740000+00:00 {
  "time": "2025-11-14T02:27:53.740392289Z",
  "level": "INFO",
  "msg": "request processed",
  "path": "/goodbye",
  "query": {
    "name": "Brian"
  },
  "status": 200,
  "message": "Goodbye, Brian"
}
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.741000+00:00 END RequestId: 5ab9e771-8e87-4a73-9d04-191893d49874
2025/11/14/[$LATEST]3870d93beada43ebb9d021dfd93bce56 2025-11-14T02:27:53.741000+00:00 REPORT RequestId: 5ab9e771-8e87-4a73-9d04-191893d49874 Duration: 1.02 ms       Billed Duration: 2 ms   Memory Size: 512 MB     Max Memory Used: 30 MB
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
[cleanup] $ ~/code/curiousdev-io/aws-lambda-container-images/custom-images/go/.config/mise/tasks/cleanup us-east-1
[*] Stopping local Docker containers for mycustomgoimagelambdafunction:1.25...
[*] Removing local Docker image curiousdev-io/custom-go:1.25...
Untagged: curiousdev-io/custom-go:1.25
[*] Removing ECR-tagged local image 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction:1.25...
Untagged: 123456789012.dkr.ecr.us-east-1.amazonaws.com/mycustomgoimagelambdafunction:1.25
[*] Deleting ECR repository mycustomgoimagelambdafunction...
[*] Deleting CloudFormation stack aws-custom-image-go...
[*] Waiting for CloudFormation stack aws-custom-image-go to be deleted...
[*] Cleanup complete.
```

</details>