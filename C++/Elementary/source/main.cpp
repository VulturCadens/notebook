#include <SDL2/SDL.h>
#include <SDL2/SDL_timer.h>
#include <SDL2/SDL_image.h>

#include <cstdlib>
#include <iostream>

#include "global.h"
#include "init.h"

int main() {

	SDL_Window* window = init();

	std::cout << "The window is " << WIDTH << " x " << HEIGHT << " pixels." << std::flush;

	SDL_Surface* screenSurface = SDL_GetWindowSurface(window);

	/* Set the icon for a window. */

	std::string iconFile(SDL_GetBasePath());
	iconFile.append(ICON_FILE);

	SDL_Surface* iconSurface = IMG_Load(iconFile.c_str());

	if(iconSurface == NULL) {

		std::cerr << "SDL_Error: " << SDL_GetError() << std::endl;
		SDL_Quit();
		exit(EXIT_FAILURE);

	}

	SDL_SetWindowIcon(window, iconSurface);
	SDL_FreeSurface(iconSurface);

	/* The Event Loop. */

	SDL_Event event {};
	bool running { true };

	while(running) {

		while(SDL_PollEvent(&event)) {

			if(event.type == SDL_QUIT) {
				running = false;
			}

		}

		SDL_FillRect(screenSurface, NULL, SDL_MapRGB(screenSurface->format, 0x06, 0x82, 0xBC));
		SDL_UpdateWindowSurface(window);

		SDL_Delay(1000 / FPS);

	}

	SDL_DestroyWindow(window);
	SDL_Quit();

	std::cout << " Exit success." << std::endl;

	return EXIT_SUCCESS;

}
