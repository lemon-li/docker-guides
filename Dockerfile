FROM docker.shiyou.kingsoft.com/library/golang:1.17.0

WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY . .

RUN go build -o main

ENV GIN_MODE=release

ENTRYPOINT [ "./main" ]