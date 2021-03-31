# Minimal Docker Image

*"Docker is a set of platform as a service (PaaS) products that use OS-level virtualization to deliver software in packages called containers. Containers are isolated from one another and bundle their own software, libraries and configuration files; they can communicate with each other through well-defined channels. Because all of the containers share the services of a single operating system kernel, they use fewer resources than virtual machines."*

https://en.wikipedia.org/wiki/Docker_(software)

Documentation: https://docs.docker.com

## Install

```bash
sudo snap install docker
sudo groupadd docker
sudo usermod -a -G docker <USER_NAME>
```

## Create Image

```bash
docker build -t <IMAGE> .

docker images
```

## Run Image in Container

```bash
docker run --name <CONTAINER> -p <HOST_PORT>:<CONTAINER_PORT> <IMAGE> &

docker ps -a
```

## Remove Container

```bash
docker stop <CONTAINER>
docker rm <CONTAINER>
```

## Remove Image

```bash
docker image prune
docker image rm <IMAGE>
```
