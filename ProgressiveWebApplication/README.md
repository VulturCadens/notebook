# Progressive Web Application

PWA must have the following features (minimium).

1. Secure contexts (HTTPS).
2. Service worker.
3. A manifest file.

## Secure Contexts

"A simple tool for making locally-trusted development certificates."

https://github.com/FiloSottile/mkcert

```console
sudo apt install libnss3-tools
```

Build **mkcert** from source (Go) or use the pre-built binaries.

```console
cd www
mkcert -install
mkcert localhost 127.0.0.1
```

Use **WWW/servehttps.py** python script (specify the right files for **keyfile=''** and **certfile=''**).

## Service Worker

## Manifest
