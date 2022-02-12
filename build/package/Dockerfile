#
# Build container
#
FROM golang as builder
ARG SERVICE

ENV CGO_ENABLED=0

COPY . /go/src/github.com/ra9dev/go-e
WORKDIR /go/src/github.com/ra9dev/go-e
RUN cd cmd/$SERVICE && \
    go build -a -tags $SERVICE -installsuffix $SERVICE -o /go/bin/$SERVICE

#
# Userspace preparation container
#
FROM alpine as alpine
ARG SERVICE

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S $SERVICE && adduser -S $SERVICE -G $SERVICE

ENTRYPOINT [ "/bin/$SERVICE" ]

#
# Runtime container
#
FROM scratch
ARG SERVICE

COPY --from=builder /go/bin/$SERVICE /bin/$SERVICE
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER $SERVICE

ENTRYPOINT ["/bin/$SERVICE"]