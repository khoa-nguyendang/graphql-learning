FROM golang:1.22-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -tags=viper_bind_struct -o /server

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/config/config.yml /app/config/config.yml
COPY --from=build-stage /app/migrations /migrations
COPY --from=build-stage /server /server
EXPOSE 8080
CMD ["/server"]