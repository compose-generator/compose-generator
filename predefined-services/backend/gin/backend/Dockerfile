FROM golang:1.16-alpine
WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["${{GIN_APP_NAME}}"]