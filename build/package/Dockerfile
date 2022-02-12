#
# Build container
#
FROM golang as builder
ARG SERVICE

ENV CGO_ENABLED=0

COPY . /go/src/github.com/ra9dev/go-e
WORKDIR /go/src/github.com/ra9dev/go-e
RUN cd cmd/$SERVICE && \
    go build -a -o /go/bin/service

#
# Userspace preparation container
#
FROM alpine as alpine

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S go-e && adduser -S go-e -G go-e

ENTRYPOINT [ "/bin/service" ]

#
# Runtime container
#
FROM scratch

COPY --from=builder /go/bin/service /bin/service
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER go-e

ENTRYPOINT ["/bin/service"]