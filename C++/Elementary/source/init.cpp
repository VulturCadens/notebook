#include <SDL2/SDL.h>

#include <cstdlib>
#include <iostream>

bool init() {
	if(SDL_Init(SDL_INIT_VIDEO) != 0) {

		std::cerr << "SDL_Error: " << SDL_GetError() << std::endl;
		return false;

	}

	return  true;
}
