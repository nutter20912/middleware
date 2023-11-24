FROM golang:1.20-alpine AS base

WORKDIR /app

COPY go.* ./

RUN go mod download

###########################################################################
# local
###########################################################################
FROM base AS local

RUN apk --update --no-cache add protobuf make curl git

# install go-micro tools
#RUN go install github.com/go-micro/cli/cmd/go-micro@v1.1.4

# debug tool - dlv
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /app

ENTRYPOINT ["sleep", "infinity"]

###########################################################################
# build image
###########################################################################
FROM base AS builder

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux \
    go build .

FROM scratch AS image

WORKDIR /app

COPY --from=builder /app/middleware ./main
COPY --from=builder /app/.env.example ./.env
COPY --from=builder /app/config/*.yaml ./config/

ENTRYPOINT ["./main"]
