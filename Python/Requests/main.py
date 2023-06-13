# pylint: skip-file

import sys
import requests

EXIT_SUCCESS = 0
EXIT_FAILURE = 1

def main() -> int:
    try:
        r = requests.get("https://example.com")

        if r.status_code == requests.codes.ok:
            print("Encoding: " + r.encoding + "\n")

            for h in r.headers:
                print(h + ": " + r.headers[h])

            return EXIT_SUCCESS

        print("ERROR: Status Code:", str(r.status_code))
        return EXIT_FAILURE

    except requests.ConnectionError as error:
        print("ERROR:", error)
        return EXIT_FAILURE


if __name__ == "__main__":
    sys.exit(main())
