FROM golang:1.14

ENV GOPATH=/go
ENV GOBIN=/go/bin

WORKDIR $GOPATH/src/provider


# Copy the code into the container
COPY . .
# Copy and download dependency using go mod
RUN go get -d -v ./...

# Build the application
RUN go install ./...
#RUN go install github.com/florian74/nats-provider/cmd/...


# Run the executable
RUN ["chmod", "+x", "/go/bin/main"]
ENTRYPOINT ["main"]

#in terminal
#run docker build -t <image_name> .
#run docker push <image_name>