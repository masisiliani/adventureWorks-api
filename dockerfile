FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/adventureWorks-api

# Install our dependencies       
RUN go get github.com/labstack/echo
RUN go get github.com/satori/go.uuid
RUN go get github.com/streadway/amqp 
RUN go get github.com/denisenkom/go-mssqldb
RUN go get golang.org/x/net/websocket
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN go get github.com/tylerb/graceful

# Install api binary globally within container
RUN go install adventureWorks-api

# Set binary as entrypoint
ENTRYPOINT /go/bin/adventureWorks-api

# Expose default port (8100)
EXPOSE 8100 