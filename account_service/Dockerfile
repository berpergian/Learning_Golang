FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bin/account_service

FROM gcr.io/distroless/base-debian12
COPY --from=build /bin/account_service /account_service
EXPOSE 8081
ENTRYPOINT ["/account_service"]
