import sys
import json


def decoding() -> int:
    """ Deserialize a JSON document (string) from a file -> a Python object """

    filename = "data.json"
    variable = "array"

    # json.load ... decoding a text file or binary file containing a JSON document
    # json.loads ... decoding a string, bytes or bytearray containing a JSON document

    try:
        with open(filename, "r") as f:
            python_object = json.load(f)

    except IOError:
        print("ERROR: could not read file '{}'".format(filename), file=sys.stderr)
        return 1

    print(python_object["object"]["integer"])

    for value in python_object["object"][variable]:
        print(value["first"])

    return 0


def encoding() -> None:
    """ Serialize a Python object -> a JSON document (string) """

    # json.dump ... encoding a Python object as a JSON formatted stream to a file object
    # json.dumps ... encoding a Python object to a JSON formatted string

    python_object = {
        "string": "John Smith",
        "float": 4.200,
        "array": [1, 2, 3]
    }

    print(json.dumps(python_object))

    print(json.dumps(python_object, sort_keys=True, indent=4))


def main() -> int:
    encoding()

    return decoding()


if __name__ == "__main__":
    sys.exit(main())
