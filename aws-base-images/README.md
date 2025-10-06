# AWS Base Images

AWS Base images are provided for several runtimes. The full list can be found in the [Amazon ECR Public Gallery](https://gallery.ecr.aws/lambda/). More recent language images make use of the [AL2023 Minimal container image
](https://docs.aws.amazon.com/linux/al2023/ug/minimal-container.html). These images are there to make it as easy as possible to run your Lambda function in a container. All dependencies, including the Lambda Runtime Interface Client, are included.

| Language | Description | README |
|:--------:|:-----------:|:----:|
| Python | Simple Lambda function triggered by a API Gateway (HTTP) endpoint | [link](./python/README.md)