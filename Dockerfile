FROM golang:1.16-alpine AS build

WORKDIR /app/

COPY . /app/
RUN go build -o app /app/server/

FROM scratch
WORKDIR /app
COPY --from=build /app/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]
# CMD["--help"]