
# Start from the latest golang base image
FROM golang:latest
# Add Maintainer Info
LABEL maintainer="Vikas"

ENV TZ="Asia/Kolkata"
RUN date

# Set the Current Working Directory inside the container
RUN mkdir -p /loco-transaction/logs
WORKDIR /loco-transaction

# Copy the current directory contents into the container at /app
ADD . /loco-transaction

RUN go build -o main .

EXPOSE 8003

# Define environment variable
ENV GIN_MODE=release

# Run the executable
CMD ["./main"]
