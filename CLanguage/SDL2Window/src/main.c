#include <stdio.h>

#include <SDL2/SDL.h>

#include "init.h"

int main(int argc, char* args[])
{
    SDL_Window* window = init();

    if (window == NULL)
        return EXIT_FAILURE;

    SDL_Surface* screenSurface = SDL_GetWindowSurface(window);

    SDL_FillRect(screenSurface, NULL, SDL_MapRGB(screenSurface->format, 0xFF, 0xFF, 0xFF));
    SDL_UpdateWindowSurface(window);

    SDL_Delay(1000);
    
    SDL_DestroyWindow(window);
    SDL_Quit();

    return EXIT_SUCCESS;
}
