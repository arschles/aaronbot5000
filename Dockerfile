ARG GOLANG_VERSION=1.15
ARG ALPINE_VERSION=3.11.5

FROM golang:${GOLANG_VERSION}-alpine AS builder

ENV GO111MODULE=on

# TODO: why do we need this?
RUN mkdir -p /src/aaronbot5000
WORKDIR /src/aaronbot5000

COPY . .


RUN CGO_ENABLED=0 GOPROXY="https://proxy.golang.org" go build -o /bin/aaronbot5000 .

FROM alpine:${ALPINE_VERSION}

COPY --from=builder /bin/aaronbot5000 /bin/aaronbot5000


EXPOSE 8080

CMD ["/bin/aaronbot5000"]
