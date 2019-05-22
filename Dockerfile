FROM golang:1.11 AS builder

# Download and install the latest release of dep
#ADD https://github.com/golang/dep/releases/download/v0.5.3/dep-linux-amd64 /usr/bin/dep
#RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/pavolloffay/github-changelog
COPY Gopkg.toml Gopkg.lock Makefile ./
RUN make install
COPY . ./
RUN make build
RUN cp -r ./build /build

FROM alpine:latest as certs
RUN apk add --update --no-cache ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build ./
ENTRYPOINT ["./gch"]
