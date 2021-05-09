FROM golang:1.16-alpine AS build
RUN apk add --update alpine-sdk
WORKDIR /src
COPY . .
RUN make build
RUN chmod +x /src/.bin/discount

FROM alpine
COPY --from=build /src/.bin/discount /discount
ENTRYPOINT ["/discount"]
