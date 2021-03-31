#!/usr/bin/env python3

import sys
import socket
import os


def main() -> int:
    if os.path.exists("/tmp/vultur.sock"):
        try:
            client = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
            client.connect("/tmp/vultur.sock")

        except ConnectionRefusedError:
            print("[ConnectionRefusedError]")
            return 1

    else:
        print("[Path doesn't exist]")
        return 1

    with client:
        while True:
            try:
                message = input("Enter a string: ")

                if message != "":
                    client.send(message.encode("utf-8"))

            except BrokenPipeError:
                print("[BrokenPipeError]")
                return 1

            except KeyboardInterrupt:
                print("\n[Ctrl-C]")
                return 0


if __name__ == "__main__":
    sys.exit(main())
