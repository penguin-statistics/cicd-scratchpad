FROM golang:1.20.7-alpine AS base
WORKDIR /app

# builder
FROM base AS builder
ENV GOOS linux
ENV GOARCH amd64

RUN apk --no-cache add bash git openssh

COPY . .

ARG VERSION

# inject versioning information & build the binary
RUN export BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ"); go build -ldflags "-X 'main.BuildTime=$BUILD_TIME' -X 'main.Version=$VERSION'" -o cicdscratchpad .

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app/cicdscratchpad /app/cicdscratchpad
EXPOSE 8080

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app/cicdscratchpad" ]
