FROM golang:alpine as builder
RUN apk --no-cache add ca-certificates

RUN mkdir /build
ADD . /build/
WORKDIR /build

# Used some compile-time parameters in the build process to instruct the
# go compiler to statically link the runtime libraries into the binary itself.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o mail .

# Minimize the entire runtime environment.
# Compile Golang app to a single binary, of which is statically-linked to its dependencies.
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build /app/
WORKDIR /app

CMD ["./mail"]
