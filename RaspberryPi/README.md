# Raspberry Pi

### go-rpio

https://github.com/stianeikeland/go-rpio

```console
go get github.com/stianeikeland/go-rpio/v4
```

### The IP address

To determine the IP address, run the nmap (192.168.10.0 -> 192.168.10.255).

```console
sudo nmap -sn 192.168.10.0/24
```

### Access GPIO pins without root

Does **/dev/gpiomem** exist?

```console
ls -l /dev/gpiomem
```

```console
sudo groupadd gpio
sudo usermod -a -G gpio USERNAME

sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem

sudo grep gpio /etc/group
```

Create and edit **/etc/udev/rules.d/gpiomem**.

```console
RUN+="/bin/sh -c 'chown root.gpio /dev/gpiomem && chmod g+rw /dev/gpiomem'"
```

## Disable WiFi and Bluetooth

```console
$ hcitool dev
$ iw dev

$ echo "dtoverlay=disable-bt" | sudo tee -a /boot/firmware/config.txt
$ echo "dtoverlay=disable-wifi" | sudo tee -a /boot/firmware/config.txt

$ sudo systemctl disable hciuart

$ sudo shutdown --reboot now
```