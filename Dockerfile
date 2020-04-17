FROM golang:1.14 as base
ENV APP_ENV=prod
EXPOSE 3000
RUN curl -sL https://taskfile.dev/install.sh | BINDIR=/usr/local/bin sh
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


FROM base as dev
ENV APP_ENV=dev
RUN go get github.com/cosmtrek/air
CMD ["/go/bin/air"]

FROM dev as dev-source
COPY . .

FROM dev-source as debug
ENV GOTRACEBACK=all
CMD ["dlv", "debug", "/app", "--accept-multiclient", "--api-version=2", "--headless", "--listen=:2345", "--log"]

FROM base as source
COPY . .
RUN task build

FROM scratch as prod
WORKDIR /app
COPY --from=source /app/main.exe ./main
CMD ["./main"]
