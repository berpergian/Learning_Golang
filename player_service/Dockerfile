FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bin/player_service

FROM gcr.io/distroless/base-debian12
COPY --from=build /bin/player_service /player_service
EXPOSE 8082
ENTRYPOINT ["/player_service"]
