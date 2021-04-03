#!/usr/bin/env python3

import os
import sys
import time
import socket
import threading

from _thread import start_new_thread


def read(connection: socket):
    while True:
        message = connection.recv(128)

        if not message:
            return

        print("Client got: " + message.decode("utf-8"))


def main() -> int:
    if os.path.exists("/tmp/vultur.sock"):
        try:
            connection = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
            connection.connect("/tmp/vultur.sock")

        except ConnectionRefusedError:
            print("[ConnectionRefusedError]")
            return 1

    else:
        print("[Path doesn't exist]")
        return 1

    start_new_thread(read, (connection,))

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
