#!/usr/bin/env python3

import os
import sys
import socket
import threading


def connection_handler(connection: socket):
    while True:
        message = connection.recv(128)

        if not message:
            print(" !! Goodbye, a client went away...")
            break

        print("Client sent: " + message.decode("utf-8"))
        connection.sendall(message)


def main() -> int:
    if os.path.exists("/tmp/vultur.sock"):
        os.remove("/tmp/vultur.sock")

    server = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    server.bind("/tmp/vultur.sock")

    with server:
        while True:
            try:
                server.listen(1)
                connection, _ = server.accept()

                ch = threading.Thread(
                    target=connection_handler,
                    args=(connection,),
                    daemon=True
                )
                
                ch.start()

            except KeyboardInterrupt:
                print("\n[Ctrl-C]")
                return 0


if __name__ == "__main__":
    sys.exit(main())
