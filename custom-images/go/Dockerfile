ARG LAMBDA_TASK_ROOT="/var/task"

# --- Build stage ---
FROM golang:1.25-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY internal/ ./internal/
COPY cmd/ ./cmd/
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o bootstrap ./cmd/main.go
RUN chmod +x bootstrap

# Download RIE for local testing
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /tmp/aws-lambda-rie
RUN chmod +x /tmp/aws-lambda-rie

COPY entry.sh /tmp/entry.sh
RUN chmod +x /tmp/entry.sh

# --- Runtime stage ---
FROM alpine:3.22 AS runtime

ARG LAMBDA_TASK_ROOT
WORKDIR ${LAMBDA_TASK_ROOT}

COPY --from=build /src/bootstrap ${LAMBDA_TASK_ROOT}
COPY --from=build /tmp/aws-lambda-rie /usr/local/bin/aws-lambda-rie
COPY --from=build /tmp/entry.sh /entry.sh
RUN chmod +x /usr/local/bin/aws-lambda-rie
RUN chmod +x bootstrap
RUN chmod +x /entry.sh
CMD ["/entry.sh"]
