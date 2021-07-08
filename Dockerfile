# Build the manager binary
FROM golang:1.13.6 as builder

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct GOSUMDB=off

WORKDIR /opt/iaso
# Copy the Go Modules manifests
COPY ./go.mod go.mod
COPY ./go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

# Copy the go source
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -a -o iaso main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.11.6
WORKDIR /opt/iaso
COPY --from=builder /opt/iaso .
ENTRYPOINT ["./iaso"]