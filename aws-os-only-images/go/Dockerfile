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

# --- Runtime stage ---
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /src/bootstrap /var/runtime/bootstrap
RUN chmod +x /var/runtime/bootstrap
CMD ["bootstrap"]
