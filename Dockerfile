FROM golang:1.21.5-bullseye AS build

RUN apt-get update && apt-get install -y git

WORKDIR /app

RUN echo user-service

RUN git clone https://github.com/akshay0074700747/user-service-grpc.git .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o bin/user-service

COPY /cmd/.env /app/cmd/bin/

FROM busybox:latest

WORKDIR /user-service

COPY --from=build /app/cmd/bin/user-service .

COPY --from=build /app/cmd/bin/.env .

EXPOSE 50002

CMD ["./user-service"]