FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/adventureWorks-api

# Install our dependencies       
RUN go get github.com/denisenkom/go-mssqldb

# Install api binary globally within container
RUN go install adventureWorks-api

# Set binary as entrypoint
ENTRYPOINT /go/bin/adventureWorks-api

# Expose default port (8100)
EXPOSE 8100 