/*
 * emcc -o index.html main.c -O3 -s WASM=1 --shell-file shell_minimal.html
 * python3 -m http.server
 *
 */

#include <stdio.h>

int main(void)
{
    printf("Minimal\n");
    return 0;
}
