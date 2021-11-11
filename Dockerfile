FROM golang:1.17.3-buster as build
WORKDIR /app
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
COPY . ./

RUN go mod vendor

RUN go build  -o /water-map ./cmd/server

FROM debian:11.1-slim AS release
COPY --from=build /water-map /water-map
RUN addgroup --gid 1001 --system nonroot && \
    adduser --uid 1001 --system --gid 1001 nonroot
USER nonroot
CMD ["/water-map"]
