FROM golang:1.16.0-alpine3.13 as build

WORKDIR /app

COPY ./Notification-Service .

RUN go mod download
RUN go build -o /notificationservice

FROM alpine:3.13
COPY --from=build /app/templates/ ./templates/
COPY --from=build /app/config/ ./config/
COPY --from=build /notificationservice /notificationservice

ENTRYPOINT  ["/notificationservice"]