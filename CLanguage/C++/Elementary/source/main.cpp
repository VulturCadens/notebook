#include <SDL2/SDL.h>
#include <SDL2/SDL_timer.h>
#include <SDL2/SDL_image.h>

#include <cstdlib>
#include <iostream>

#include "global.h"
#include "init.h"

int main()
{
	SDL_Window* window { init() };

	std::cout << "The windows is " << WIDTH << " x " << HEIGHT << " pixels.\n";
	std::cout << "Press ESC to exit..." << std::flush;

	SDL_Surface* screenSurface { SDL_GetWindowSurface(window) };

	/* Set the icon for a window. */

	std::string iconFile { SDL_GetBasePath() };
	iconFile.append(ICON_FILE);

	SDL_Surface* iconSurface { IMG_Load(iconFile.c_str()) };

	if (iconSurface == NULL) {

		std::cerr << "SDL_Error: " << SDL_GetError() << std::endl;
		SDL_Quit();
		exit(EXIT_FAILURE);

	}

	SDL_SetWindowIcon(window, iconSurface);
	SDL_FreeSurface(iconSurface);

	/* Create an SDL surface using an SVG from an string. */

	const std::string svgString =
		"<svg height='100' width='100'>"
		"<circle cx='50' cy='50' r='45' stroke='#202020' stroke-width='2' fill='#237847'/>"
		"</svg>";

	SDL_RWops *rw = SDL_RWFromConstMem(svgString.c_str(), svgString.size());
	SDL_Surface *svgSurface = IMG_Load_RW(rw, 1);

	struct SDL_Rect target { 50, 50, 100, 100 };

	/* The event loop. */

	SDL_Event event {};
	bool running { true };

	while (running) {

		while (SDL_PollEvent(&event)) {

			/* https://github.com/libsdl-org/SDL/blob/SDL2/include/SDL_keycode.h */

			if(event.type == SDL_KEYDOWN) {

				switch(event.key.keysym.sym) {
					case SDLK_ESCAPE:
						running = false;
						break;
					default:
						break;
				}

			} else if (event.type == SDL_MOUSEBUTTONDOWN) {

				if (event.button.button == SDL_BUTTON_LEFT) {
					SDL_GetMouseState(&target.x, &target.y);
				}

			} else if(event.type == SDL_QUIT) {

				running = false;

			}

		}

		SDL_FillRect(screenSurface, NULL, SDL_MapRGB(screenSurface->format, 0x06, 0x82, 0xBC));
		SDL_BlitSurface(svgSurface, NULL, screenSurface, &target);
		SDL_UpdateWindowSurface(window);

		SDL_Delay(1000 / FPS);

	}

	SDL_DestroyWindow(window);
	SDL_Quit();

	std::cout << " Exit success." << std::endl;

	return EXIT_SUCCESS;
}
