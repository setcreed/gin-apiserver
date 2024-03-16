FROM golang:1.21-alpine as builder
WORKDIR /app
ARG VERSION
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
COPY ./go.sum ./
#RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X 'main.version=${VERSION}'" -o gin-apiserver ./cmd

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/gin-apiserver /app/
COPY etc/gin-apiserver.yaml /etc/gin-apiserver/gin-apiserver.yaml

ENTRYPOINT ["./gin-apiserver"]
