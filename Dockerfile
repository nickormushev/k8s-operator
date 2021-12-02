# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY webhooks/ webhooks/

COPY ./webhooks_certs/tls.crt /tmp/k8s-webhook-server/serving-certs/
COPY ./webhooks_certs/tls.key /tmp/k8s-webhook-server/serving-certs/
RUN chmod -R 777 /tmp/k8s-webhook-server/serving-certs/tls.crt
RUN chmod -R 777 /tmp/k8s-webhook-server/serving-certs/tls.key

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /tmp/k8s-webhook-server/serving-certs/tls.crt /tmp/k8s-webhook-server/serving-certs/tls.crt
COPY --from=builder /tmp/k8s-webhook-server/serving-certs/tls.key /tmp/k8s-webhook-server/serving-certs/tls.key
USER 65532:65532


ENTRYPOINT ["/manager"]
