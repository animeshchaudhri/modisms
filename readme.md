# Concurrent HTTP PATCH Request Sender in Go

This project is a Go program that sends concurrent HTTP PATCH requests with multipart/form-data. The program demonstrates how to use goroutines, channels, and WaitGroups to handle concurrency effectively in Go.

## Features

- Sends multiple concurrent HTTP PATCH requests
- Handles multipart/form-data requests
- Uses goroutines for concurrency
- Collects and prints responses from all requests

## Prerequisites

- Go 1.13+ installed on your machine

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/go-concurrent-http-request-sender.git
   cd go-concurrent-http-request-sender
   ```

2. Build the Go program:

   ```sh
   go build -o concurrent-request-sender main.go

   ```

3. usage

```sh
  ./concurrent-request-sender
```
