# pylint: skip-file

import sys
import requests


def main() -> int:
    try:
        r = requests.get("https://example.com")

        if r.status_code == requests.codes.ok:
            print("Encoding: " + r.encoding + "\n")

            for h in r.headers:
                print(h + ": " + r.headers[h])

            return 0

        print("ERROR: Status Code:", str(r.status_code))
        return 1

    except requests.ConnectionError as error:
        print("ERROR:", error)
        return 1


if __name__ == "__main__":
    sys.exit(main())
