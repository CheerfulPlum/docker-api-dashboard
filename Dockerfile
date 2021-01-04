FROM golang:1.15 as build
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w'

FROM scratch as prod

COPY --from=build  /app/docker-dashboard-api /
CMD ["/docker-dashboard-api"]