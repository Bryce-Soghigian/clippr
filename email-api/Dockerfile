FROM golang:latest

# Set the working directory to the root of the project
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the Go server
RUN go build -o server .

# Expose the default port for the Go server
EXPOSE 8080

# Run the Go server when the container starts
CMD ["./server"]

