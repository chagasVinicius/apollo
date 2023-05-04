# Build the Go Binary.
FROM golang:1.20 as build_apollo-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /service

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/cmd/services/apollo-api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Build the admin binary.
WORKDIR /service/cmd/tooling/admin
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.17
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S appuser && \
    adduser -u 1000 -h /service -G appuser -S appuser
COPY --from=build_apollo-api --chown=appuser:appuser /service/cmd/services/apollo-api/apollo-api /service/apollo-api
COPY --from=build_apollo-api --chown=appuser:appuser /service/cmd/tooling/admin/admin /service/admin

WORKDIR /service
USER appuser
CMD ["./apollo-api"]
