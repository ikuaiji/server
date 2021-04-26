FROM golang:alpine as build
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/webserver

FROM scratch
COPY --from=build /work/webserver /bin/app
EXPOSE 80
ENTRYPOINT ["/bin/app"]
