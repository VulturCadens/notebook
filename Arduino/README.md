# Arduino

* Language Reference: https://www.arduino.cc/reference/en

## Arduino Micro

* Microcontroller: ATmega32U4
* Operating Voltage: 5V
* Input Voltage: 7-9V
* Digital I/O Pins: 20
* PWM Channels: 7
* Analog Input Channels: 12 (10-bit ADC)
* DC Current per I/O Pin: 20 mA
* DC Current for 3.3V Pin: 50 mA
* Flash Memory: 32 KB of which 4 KB used by bootloader
* Clock Speed: 16 MHz
* LED_BUILTIN: pin 13

## Arduino-cli installation

Download a pre-built binary tarball from the link https://github.com/arduino/arduino-cli.

```console
$ tar -xzf arduino-cli_0.18.3_Linux_64bit.tar.gz -C ~/bin

$ ls -l /dev/ttyACM*

crw-rw---- 1 root dialout 166, 0 Jun 25 20:52 /dev/ttyACM0

$ sudo usermod -a -G dialout <USERNAME>
```

__Reboot the computer.__

```console
$ arduino-cli board list

Port          Type               Board Name     FQBN               Core
/dev/ttyACM0  Serial Port (USB)  Arduino Micro  arduino:avr:micro  arduino:avr
/dev/ttyS0    Serial Port        Unknown
/dev/ttyS1    Serial Port        Unknown
```

```console
$ arduino-cli core install arduino:avr

Downloading packages...
arduino:avr-gcc@7.3.0-atmel3.6.1-arduino7 downloaded
arduino:avrdude@6.3.0-arduino17 downloaded
arduino:arduinoOTA@1.3.0 downloaded
arduino:avr@1.8.3 downloaded
Installing arduino:avr-gcc@7.3.0-atmel3.6.1-arduino7...
arduino:avr-gcc@7.3.0-atmel3.6.1-arduino7 installed
Installing arduino:avrdude@6.3.0-arduino17...
arduino:avrdude@6.3.0-arduino17 installed
Installing arduino:arduinoOTA@1.3.0...
arduino:arduinoOTA@1.3.0 installed
Installing arduino:avr@1.8.3...
Configuring platform...
arduino:avr@1.8.3 installed
```

```console
$ arduino-cli core list

ID           Installed  Latest  Name
arduino:avr  1.8.3      1.8.3   Arduino AVR Boards

$ arduino-cli core update-index

Updating index: package_index.json downloaded
Updating index: package_index.json.sig downloaded
```

## Arduino-cli usage

```console
$ arduino-cli sketch new Blink
```

Edit open and edit the ino file (__./Blink/Blink.ino__ in this case).

```console
$ arduino-cli compile --fqbn arduino:avr:micro Blink

Sketch uses 3958 bytes (13%) of program storage space. Maximum is 28672 bytes.
Global variables use 149 bytes (5%) of dynamic memory, leaving 2411 bytes for local variables. Maximum is 2560 bytes.

$ arduino-cli upload -p /dev/ttyACM0 --fqbn arduino:avr:micro Blink

Connecting to programmer: .
Found programmer: Id = "CATERIN"; type = S
Software Version = 1.0; No Hardware Version given.
Programmer supports auto addr increment.
Programmer supports buffered memory access with buffersize=128 bytes.
Programmer supports the following devices: Device code: 0x44
```
