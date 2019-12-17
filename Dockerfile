# References: https://docs.docker.com/engine/reference/builder

FROM golang:latest

# Maintainer info
LABEL maintainer="Zach Hanson <zacharyjhanson@gmail.com>"

# Current working directory in the container
WORKDIR /app

# Copy over Go's mod and sum files
COPY go.mod go.sum ./

# Download dependencies. If go.mod and go.sum are unchanged, dependencies will be cached
RUN go mod download

# Copy source from current directory to the Working Directory in the container
COPY . .

# Build the app
RUN go build -o main .

# Open required port, 8080
EXPOSE 8081

# Run the executable
CMD ["./main"]up
