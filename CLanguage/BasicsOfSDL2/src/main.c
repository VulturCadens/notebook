#include <stdio.h>
#include <stdlib.h>

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>

#include "global.h"
#include "init.h"

#define IMAGE_FILE "./resources/box.png"

int main(void)
{
    SDL_SetHint(SDL_HINT_RENDER_SCALE_QUALITY, "1");

    SDL_Window* window = init();

    if (window == NULL)
        return EXIT_FAILURE;

    SDL_Renderer* renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);

    if (renderer == NULL) {
        fprintf(stderr, "Could not create renderer: %s\n", SDL_GetError());
        return EXIT_FAILURE;
    }

    if (SDL_SetRenderDrawColor(renderer, 0xAA, 0xAA, 0xCC, 0xFF) != 0) {
        fprintf(stderr, "Could not set renderer color: %s\n", SDL_GetError());
        return EXIT_FAILURE;
    }

    SDL_Texture* texture = IMG_LoadTexture(renderer, IMAGE_FILE);

    if (texture == NULL) {
        fprintf(stderr, "Could not load image: %s\n", SDL_GetError());
        return EXIT_FAILURE;
    }

    SDL_Rect* r = (SDL_Rect*)malloc(sizeof(SDL_Rect));

    if (r == NULL) {
        fprintf(stderr, "The memory could not be allocated: malloc(sizeof(SDL_Rect))\n");
        return EXIT_FAILURE;
    }

    r->x = 200;
    r->y = 200;
    r->w = 64;
    r->h = 64;

    /*
     * Uint64 SDL_GetPerformanceFrequency(void);
     *     Returns a platform-specific count per second.
     * 
     * Uint64 SDL_GetPerformanceCounter(void);
     *     Returns the current counter value.
     */

    const double target = (double)SDL_GetPerformanceFrequency() / FPS;

    uint64_t now = SDL_GetPerformanceCounter();
    uint64_t last = 0;
    uint64_t delta = 0;

    while (r->x < 600) {
        r->x++;

        SDL_RenderClear(renderer);
        SDL_RenderCopy(renderer, texture, NULL, r);
        SDL_RenderPresent(renderer);

        delta = 0;
        last = now;

        while (delta < target) {
            now = SDL_GetPerformanceCounter();
            delta = (double)(now - last);
        }
    }

    SDL_Delay(500);

    SDL_DestroyTexture(texture);
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);

    free(r);

    SDL_Quit();

    return EXIT_SUCCESS;
}
