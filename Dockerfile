# Build the Go Binary.
FROM golang:1.20 as build
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /service

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/app/services/cora-email
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Run the Go Binary in Alpine.
FROM alpine:3.17
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S appuser && \
    adduser -u 1000 -h /service -G appuser -S appuser
COPY --from=build --chown=appuser:appuser /service/app/services/cora-email/cora-email /service/cora-email
WORKDIR /service
USER appuser
CMD ["./cora-email"]
