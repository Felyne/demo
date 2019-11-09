# 编译阶段
FROM 192.168.41.34/pub/golang:1.12 as builder
WORKDIR /build
COPY . .
RUN  GOPROXY=http://192.168.40.131:4000 make

# 运行阶段
FROM alpine:3.10 as runner
ARG GIT_TAG
LABEL git_tag=$GIT_TAG description="test image"
WORKDIR /app
COPY --from=builder /build/app ./
EXPOSE 8080
ENTRYPOINT ["./app"]