#!/usr/bin/env python3

# pip3 install pyserial
# sudo usermod -a -G dialout <USERNAME>

import serial
import time

PORT = "/dev/ttyACM0"


def main():
    led_state = "HIGH"

    try:
        arduino = serial.Serial(PORT, baudrate=19200)
    except:
        print("Could not open port {}".format(PORT))
        return

    print("Using {}\n".format(arduino.name))

    time.sleep(1)

    with arduino:
        while(True):
            try:
                data = arduino.readline()
                value = int.from_bytes(data[0:2], byteorder="big")
                print("\rAnalog value (A0): {}  ".format(value), end="")

                led_state = "HIGH" if led_state == "LOW" else "LOW"
                arduino.write(bytes("[{}]".format(led_state), "utf-8"))

            except KeyboardInterrupt:
                print("\n[Ctrl-C]")
                break


if __name__ == "__main__":
    main()
