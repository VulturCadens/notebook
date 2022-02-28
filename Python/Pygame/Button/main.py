import os
import pygame as SDL

from classes.button import Button

FPS = 30

DIRECTORY = os.path.split(os.path.abspath(__file__))[0]

BOX_PATH = os.path.join(DIRECTORY, "images", "box.png")
BACKGROUND_PATH = os.path.join(DIRECTORY, "images", "background.jpg")
BUTTON_PATH = os.path.join(DIRECTORY, "images", "button.png")


def button_action_1():
    print("[ 1. button was clicked ]")


def button_action_2():
    print("[ 2. button was clicked ]")


def button_action_3():
    print("[ 3. button was clicked ]")


def input_handling():
    for event in SDL.event.get():
        if event.type in (SDL.QUIT, SDL.KEYDOWN):
            return False

        elif event.type == SDL.MOUSEBUTTONDOWN:
            left_button, middle_button, right_button = SDL.mouse.get_pressed()

            if left_button:
                position = SDL.mouse.get_pos()
                print("■☐☐ The left mouse button")

                Button.click(position)

            elif middle_button:
                print("☐■☐ The middle mouse button")

            elif right_button:
                print("☐☐■ The right mouse button")

    return True


def main():
    speed = 1
    x = 100
    clock = SDL.time.Clock()
    is_running = True

    SDL.init()

    screen = SDL.display.set_mode(size=(1000, 500), vsync=1)

    box = SDL.image.load(BOX_PATH).convert()
    background = SDL.image.load(BACKGROUND_PATH).convert()
    window_icon = SDL.image.load("icon.png")  # 32 x 32 pixels

    SDL.display.set_icon(window_icon)
    SDL.display.set_caption("Window Title")

    screen.blit(background, (0, 0))

    Button(BUTTON_PATH, (100, 300), button_action_1)
    Button(BUTTON_PATH, (300, 300), button_action_2)
    Button(BUTTON_PATH, (500, 300), button_action_3)

    Button.draw(screen)

    while is_running:
        is_running = input_handling()

        source_rect = (x, 100, 64, 64)
        screen.blit(source=background, dest=(x, 100), area=source_rect)

        x += speed

        screen.blit(box, (x, 100))

        if x >= 900 or x <= 10:
            speed *= -1

        SDL.display.update()

        # By calling Clock.tick(X) once per frame, the program will never run at more
        # than X frames per second. This function uses SDL_Delay function which is not
        # accurate on every platform, but does not use much CPU.

        clock.tick(FPS)


if __name__ == "__main__":
    main()
    SDL.quit()
