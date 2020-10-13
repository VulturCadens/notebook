# Just a few notes about Ebiten

## Essential Links

Repository: https://github.com/hajimehoshi/ebiten

Website: https://ebiten.org/

API: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2

## Install 

### Install Dependencies

```console
sudo apt install gcc

sudo apt install libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxcursor-dev

sudo apt install libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
```

### Install Ebiten

```console
go get github.com/hajimehoshi/ebiten
```

Or create local module - and Go automatically installs Ebiten.

```console
mkdir foo

cd foo

go mod init github.com/foobar/foo
```