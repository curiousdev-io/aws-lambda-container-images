# Custom Images

Custom images need to include all of the necessary components to invoke a Lambda function. In this case, we'll build out two Lambda functions based upon custom OCI images.

| Language | Description | README |
|:--------:|:-----------:|:----:|
| Go | Simple Lambda function triggered by a API Gateway (HTTP) endpoint | [link](./go/README.md)
| Python | Simple Lambda function triggered by a API Gateway (HTTP) endpoint | [link](./python/README.md)

## Why Use Custom Images?

You may have a need to control what is included in your images. Your employer may enforce requirements about build and runtime containers. Even if you don't, controlling what is included in an image can result in smaller image sizes with a better security posture.

## What Do I Need in a Custom Image?

The image needs the ability to execute your AWS Lambda function. Python functions need a Python runtime along with all module dependencies. Go functions are statically compiled and can run on a minimal container.

However, you will also need to include the Runtime Interface Client (RIC) for a deployed function to work with the AWS Lambda service. The RIC polls the AWS Lambda Runtime API for new events and sends responses back to the Lambda service.

If you plan on invoking your function locally, you will also need to include the Runtime Interface Emulator (RIE). 

## Size Difference

One of the potential benefits of using a custom image as a basis for a Lambda function is the container size difference.

| Language | AWS Base Image | AWS OS-only Image | Custom Image |
|:--------:|:-----------:|:----:|:----:|
| Go | n/a | 42.40 MB | 11 MB |
| Python | 185 MB | n/a | 40.86 MB|
