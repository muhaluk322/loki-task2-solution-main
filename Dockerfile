FROM golang:1.21-alpine AS build-stage

WORKDIR /app
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go mod download && go build -o webapp

FROM alpine

WORKDIR /app
COPY --from=build-stage /app /app
EXPOSE 3000

ENTRYPOINT ["/app/webapp"]
