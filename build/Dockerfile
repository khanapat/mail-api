FROM amd64/alpine:latest
# FROM scratch

WORKDIR /app

COPY goapp ./
COPY ./template ./template
ENV TZ=Asia/Bangkok

ENTRYPOINT ["/app/goapp"]