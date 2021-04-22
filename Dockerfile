FROM golang:alpine as build
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/webserver

FROM atfeng.com/laokuaiji/web as static

FROM scratch
COPY --from=build /work/webserver /bin/app
COPY --from=static /usr/share/nginx/html /web/dist
EXPOSE 80
ENTRYPOINT ["/bin/app"]
