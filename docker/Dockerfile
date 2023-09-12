FROM golang:1.20-alpine

WORKDIR /app

###########################################################################
# local
###########################################################################
RUN apk --update --no-cache add protobuf make curl git

# install go-micro tools
RUN go install github.com/go-micro/cli/cmd/go-micro@v1.1.4

# debug tool - dlv
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.* ./

RUN go mod download

COPY . /app
