# Simple web server

This document describes a simple web server written in Golang.
The server accepts **HTTPS** connection on port 8443 (locally).

## Requirements

- go 1.16+
- self-signed ssl certificates.

## Installation

- Clone the repository into a directory of your choice.
- Generate self-signed certificates and make sure they are in the same directory as main.go file.

  - Consult the docs for your operating system procedure.

## Testing

Run the following from the terminal. Make sure you are in the same directory as main.go file

    go run main.go
Open a web browser and type 'https://localhost:8433' in the browser address bar.

## Other information

You can do away with ssl/tls security and let server listen on **HTTP** port. This you can accomplish by changing function

    ListenAndServeTLS > ListenAndServe
You need also to edit the server's listening port from

    8433 > 8080

Restart the server and access it via 'http://localhost:8080'

## Future Improvements

Handling dynamic HTML files.
