FROM golang:1.14
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory
WORKDIR /home/tzuhan960691/go_learning

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .
RUN go build -o main .


# Open port
EXPOSE 8080
# Command to run when starting the container
CMD ["/home/tzuhan960691/go_learning/main"]
