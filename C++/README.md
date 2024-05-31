# C++ Notes

## Setting Up SDL 2

Simple DirectMedia Layer: https://www.libsdl.org

```
$ apt-cache search libsdl2

    libsdl2-2.0-0 - Simple DirectMedia Layer
    libsdl2-doc - Reference manual for libsdl2
    libsdl2-dev - Simple DirectMedia Layer development files
    libsdl2-gfx-1.0-0 - drawing and graphical effects extension for SDL2
    libsdl2-gfx-dev - development files for SDL2_gfx
    libsdl2-gfx-doc - documentation files for SDL2_gfx
    libsdl2-image-2.0-0 - Image loading library for Simple DirectMedia Layer 2, libraries
    libsdl2-image-dev - Image loading library for Simple DirectMedia Layer 2, development files
    libsdl2-mixer-2.0-0 - Mixer library for Simple DirectMedia Layer 2, libraries
    libsdl2-mixer-dev - Mixer library for Simple DirectMedia Layer 2, development files
    libsdl2-net-2.0-0 - Network library for Simple DirectMedia Layer 2, libraries
    libsdl2-net-dev - Network library for Simple DirectMedia Layer 2, development files
    libsdl2-ttf-2.0-0 - TrueType Font library for Simple DirectMedia Layer 2, libraries
    libsdl2-ttf-dev - TrueType Font library for Simple DirectMedia Layer 2, development files

$ apt-get install libsdl2-dev

    ...

$ sdl2-config --libs --cflags

    -lSDL2
    -I/usr/include/SDL2 -D_REENTRANT
```
