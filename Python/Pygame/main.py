import os
import pygame as SDL

FPS = 30

DIRECTORY = os.path.split(os.path.abspath(__file__))[0]

BOX_PATH = os.path.join(DIRECTORY, "images", "box.png")
BACKGROUND_PATH = os.path.join(DIRECTORY, "images", "background.jpg")


def main():
    speed = 1
    x = 100
    clock = SDL.time.Clock()

    SDL.init()

    screen = SDL.display.set_mode(size=(1000, 500), vsync=1)

    box = SDL.image.load(BOX_PATH).convert()
    background = SDL.image.load(BACKGROUND_PATH).convert()
    window_icon = SDL.image.load("icon.png")  # 32 x 32 pixels

    SDL.display.set_icon(window_icon)
    SDL.display.set_caption("Window Title")

    screen.blit(background, (0, 0))

    while True:
        for event in SDL.event.get():
            if event.type in (SDL.QUIT, SDL.KEYDOWN):
                return

            elif event.type == SDL.MOUSEBUTTONDOWN:
                left_button, middle_button, right_button = SDL.mouse.get_pressed()

                if left_button:
                    position = SDL.mouse.get_pos()
                    print("Left mouse button event: {}".format(position))

                elif middle_button:
                    print("Middle mouse button event")

                elif right_button:
                    print("Right mouse button event")

        source_rect = (x, 100, 64, 64)
        screen.blit(source=background, dest=(x, 100), area=source_rect)

        x += speed

        screen.blit(box, (x, 100))

        if x >= 900 or x <= 10:
            speed *= -1

        SDL.display.update()

        clock.tick(FPS)


if __name__ == "__main__":
    main()
    SDL.quit()
