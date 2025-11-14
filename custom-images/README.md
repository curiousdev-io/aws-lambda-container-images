# Custom Images

Custom images need to include all of the necessary components to invoke a Lambda function. In this case, we'll build out two Lambda functions based upon Chainguard images.

| Language | Description | README |
|:--------:|:-----------:|:----:|
| Go | Simple Lambda function triggered by a API Gateway (HTTP) endpoint | [link](./go/README.md)
| Python | Simple Lambda function triggered by a API Gateway (HTTP) endpoint | [link](./python/README.md)

## What is Chainguard?

> Chainguard is a company that provides safe, open source software, continuously rebuilt from source in secure environments with end-to-end integrity.

Chainguard images are **distroless**. Distroless images are container images stripped of unnecessary components such as package managers, shells, or even the underlying operating system distribution. We can use them as the base for our AWS Lambda functions to reduce security vulnerabilities.

## Size Difference

One of the other benefits of using a distroless image as a basis for a Lambda function is the container size difference.

| Language | AWS Base Image | Custom Image |
|:--------:|:-----------:|:----:|
| Python | 185 MB | 40.86 MB|
