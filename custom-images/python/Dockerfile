ARG LAMBDA_TASK_ROOT="/var/task"

# --- Build stage ---
FROM cgr.dev/chainguard/python:latest-dev AS builder

ARG LAMBDA_TASK_ROOT
WORKDIR /app

# Copy requirements and install dependencies
COPY requirements.txt .
RUN pip install --target ./packages -r requirements.txt && \
    pip install --target ./packages awslambdaric

# Copy source code
COPY src/ ./src/

# Download RIE to a writable location and set permissions
USER root
ADD --chmod=755 https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /app/aws-lambda-rie

# Create entrypoint script in build stage
COPY --chmod=755 entry.py /app/entry.py

# --- Runtime stage ---
FROM cgr.dev/chainguard/python:latest AS runtime

ARG LAMBDA_TASK_ROOT
WORKDIR ${LAMBDA_TASK_ROOT}

# Copy everything from builder (no RUN commands needed)
COPY --from=builder /app/packages ${LAMBDA_TASK_ROOT}
COPY --from=builder /app/src/ ${LAMBDA_TASK_ROOT}/
COPY --from=builder /app/aws-lambda-rie /usr/local/bin/aws-lambda-rie
COPY --from=builder /app/entry.py /entry.py

ENTRYPOINT ["python", "/entry.py"]
CMD ["lambda_function.handler"]
