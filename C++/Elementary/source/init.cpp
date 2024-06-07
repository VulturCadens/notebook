#include <SDL2/SDL.h>

#include <cstdlib>
#include <iostream>

#include "global.h"

SDL_Window* init() {

	if(SDL_Init(SDL_INIT_VIDEO) != 0) {

		std::cerr << "SDL_Error: " << SDL_GetError() << std::endl;
		exit(EXIT_FAILURE);

	}

	SDL_Window* window = SDL_CreateWindow("Window",
			SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED,
			WIDTH, HEIGHT, SDL_WINDOW_SHOWN);

	if(window == NULL) {

			std::cerr << "SDL_Error: " << SDL_GetError() << std::endl;
			SDL_Quit();
			exit(EXIT_FAILURE);

	}

	return window;

}
