FROM golang:${{FIBER_GO_VERSION}}-alpine
WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["${{FIBER_APP_NAME}}"]