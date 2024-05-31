#include <SDL2/SDL.h>
#include <SDL2/SDL_timer.h>

#include <cstdlib>
#include <stdio.h>

static const int WIDTH = 500;
static const int HEIGHT = 300;

int main() {

	SDL_Window* window = NULL;
	SDL_Surface* screenSurface = NULL;

	if( SDL_Init(SDL_INIT_VIDEO) != 0) {

		printf("SDL_Error: %s\n", SDL_GetError());
		return EXIT_FAILURE;

	}

	window = SDL_CreateWindow("Window",
		SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED,
		WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

	if( window == NULL )  {

		printf("SDL_Error: %s\n", SDL_GetError());
		return EXIT_FAILURE;

	}

	screenSurface = SDL_GetWindowSurface(window);

	SDL_Event event;
	bool running = true;

	while(running) {

		while(SDL_PollEvent(&event)) {

			if(event.type == SDL_QUIT) {
				running = false;
			}

		}

		SDL_FillRect(screenSurface, NULL, SDL_MapRGB(screenSurface->format, 0xAA, 0xAA, 0xFF));
		SDL_UpdateWindowSurface(window);

		SDL_Delay(1000 / 60);

	}

	SDL_DestroyWindow(window);
	SDL_Quit();

	return EXIT_SUCCESS;

}
