FROM golang:1.18-alpine AS builder

WORKDIR /intellimall

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY .. ./

# Build the binary.
RUN go build -ldflags="-s -w" -v -o monolith ./cmd/intellimall

FROM alpine:latest AS runtime

COPY --from=builder /intellimall/docker/wait-for .
RUN chmod +x /wait-for

# Copy the binary to the production image from the builder stage.
COPY --from=builder /intellimall/monolith /intellimall/monolith

CMD ["/intellimall/monolith"]
