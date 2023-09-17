
# Start from the latest golang base image
FROM golang:latest
# Add Maintainer Info
LABEL maintainer="Vikas"

ENV TZ="Asia/Kolkata"
RUN date

# Set the Current Working Directory inside the container
RUN mkdir -p /loco-transaction/logs
WORKDIR /loco-transaction

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download
RUN apt-get -y update
RUN apt-get install -y supervisor
RUN apt-get install net-tools
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main .
# Expose port 8069 to the outside world
EXPOSE 8003
