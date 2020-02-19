Transform UDP datagrams in HTTP POST messages.

This repository contains three folders that represent individual components that can be used for demos.

# udp2http
This application listens on UDP port 1053 for datagrams. Upon receival of one, the payload is forwarded to http://127.0.0.1:51000 using a POST request.

# udpSend
This application sends the first command line argument as a payload in a UDP datagram to localhost port 1053.

# httpServer
This program is a HTTP server listening on port 51000 that print all requests to stdout.

# Demo workflow
1. Start the HTTP server
1. Start the udp2http
1. Start sending UDP datagrams with arbitrary payload using the udpSend binary 
