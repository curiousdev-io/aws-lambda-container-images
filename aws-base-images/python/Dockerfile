# Image from https://gallery.ecr.aws/lambda/python
FROM public.ecr.aws/lambda/python:3.13

# Copy requirements.txt
COPY requirements.txt ${LAMBDA_TASK_ROOT}

# Install the specified packages
RUN pip install -r requirements.txt


# Copy function code
COPY src ${LAMBDA_TASK_ROOT}

# Set default CMD for Lambda
CMD [ "lambda_function.handler" ]
