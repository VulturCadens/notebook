#!/usr/bin/env python3

import sys
import socket
import os

def main() -> int:
    if os.path.exists("/tmp/vultur.sock"):
        client = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        client.connect("/tmp/vultur.sock")
    else:
        print("Connection error.")
        return 1

    while True:
        try:
            message = input("Enter a string: ")

            if message != "":
                client.send(message.encode("utf-8"))

        except KeyboardInterrupt:
            print("\n[Ctrl-C]")
            client.close()
            break

    return 0


if __name__ == "__main__":
    sys.exit(main())
