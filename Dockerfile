FROM golang:1.13

RUN apt-get -y update \
    && apt-get -y upgrade 
    

WORKDIR /app

# Enables go modules inside GOPATH
ENV GO111MODULE=on

# Copy go modules
COPY go.mod go.sum ./

# Install dependecies.
RUN go mod download

COPY . ./

# RUN make build
RUN go build -v -o eth-stats cmd/server/server.go

CMD ["/app/eth-stats"]
