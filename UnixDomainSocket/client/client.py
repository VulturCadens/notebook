#!/usr/bin/env python3

import os
import sys
import time
import socket
import threading


def read_socket(connection: socket):
    while True:
        message = connection.recv(128)

        if not message:
            return

        print("Client got: " + message.decode("utf-8"))


def main() -> int:
    if not os.path.exists("/tmp/vultur.sock"):
        print("[Path doesn't exist]")
        return 1

    try:
        connection = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        connection.connect("/tmp/vultur.sock")

    except ConnectionRefusedError:
        print("[ConnectionRefusedError]")
        return 1

    rs = threading.Thread(target=read_socket, args=(connection,), daemon=True)
    rs.start()

    with connection:
        while True:
            try:
                message = input("Enter a string: ")

                if message:
                    connection.send(message.encode("utf-8"))

            except BrokenPipeError:
                print("[BrokenPipeError]")
                return 1

            except KeyboardInterrupt:
                print("\n[Ctrl-C]")
                return 0

            time.sleep(0.25)


if __name__ == "__main__":
    sys.exit(main())
