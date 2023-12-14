# The base go-image
FROM golang:1.20-alpine
 
# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app

# Run command as described:
# go build will build an executable file named server in the current directory

RUN  go build -o /usr/local/go/bin/dnstwist-go /app/cmd/main.go

RUN chmod +x /usr/local/go/bin/dnstwist-go
 
# Run the server executable
CMD [ "dnstwist-go"]