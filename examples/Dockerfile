FROM golang:1.19.6-buster as build

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . . 
RUN go build -o /server

FROM gcr.io/distroless/base-debian11
COPY --from=build /server /server

EXPOSE 8088
ENTRYPOINT ["/server"]