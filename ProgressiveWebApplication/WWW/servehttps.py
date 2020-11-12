#!/usr/bin/env python3

from http.server import HTTPServer,SimpleHTTPRequestHandler
from socketserver import BaseServer
import ssl

httpd = HTTPServer(("127.0.0.1", 4443), SimpleHTTPRequestHandler)

httpd.socket = ssl.wrap_socket(
    httpd.socket,
    certfile="localhost.pem",
    keyfile="localhost-key.pem",
    server_side=True)

httpd.serve_forever()
