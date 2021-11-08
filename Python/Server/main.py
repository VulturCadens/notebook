# pylint: skip-file

import sys

from http.server import BaseHTTPRequestHandler, HTTPServer

PORT = 8000


class Handler(BaseHTTPRequestHandler):

    def set_headers(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()

    def do_GET(self):
        if self.path == "/":
            self.set_headers()
            self.wfile.write(bytes("GET /", "utf8"))

        elif self.path == "/foo":
            self.set_headers()
            self.wfile.write(bytes("GET /foo", "utf8"))

        elif self.path == "/favicon.ico":
            self.send_response(200)
            self.send_header("Content-type", "image/png")
            self.end_headers()

            self.wfile.write(bytes("data:image/png;base64,iVBORw0KGgo=", "utf8"))

        else:
            self.send_error(404, "File Not Found: {}".format(self.path))

    def do_POST(self):
        self.set_headers()
        self.wfile.write(bytes("POST ALL RESPONSE", "utf8"))


def main() -> int:
    try:
        server = HTTPServer(("localhost", PORT), Handler)
        print("Server running on port {} \n".format(PORT))
        server.serve_forever()

    except KeyboardInterrupt:
        print("\n ... Keyboard interrupt -> stopping server")
        server.socket.close()

    return 0


if __name__ == "__main__":
    sys.exit(main())
