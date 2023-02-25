import pygame as SDL

from typing import Tuple
from pygame import gfxdraw


def rotate_center(image, angle, x, y) -> Tuple:
    rotated_image = SDL.transform.rotate(image, angle)
    center = image.get_rect(center=(x, y)).center
    rect = rotated_image.get_rect(center=center)

    return (rotated_image, rect)


SDL.init()

BACKGROUND = (20, 20, 20)
TEXT_COLOR = (200, 200, 200)
LINE_COLOR = (200, 150, 150)

FPS = 30

window = SDL.display.set_mode(size=(500, 500), vsync=1)

is_running = True

ship = SDL.image.load("ship.png").convert_alpha()
font = SDL.font.Font("Roboto-Regular.ttf", 20)
clock = SDL.time.Clock()

rotation = 0    # The degree as a unit of rotations.

SDL.display.set_caption("Rotation")
SDL.display.set_icon(ship)

while is_running:
    window.fill(BACKGROUND)

    text = font.render("Rotation: {}°".format(int(rotation % 360)),
                       True,
                       TEXT_COLOR,
                       BACKGROUND)

    window.blit(text, (10, 10))

    for event in SDL.event.get():
        if event.type in (SDL.QUIT, SDL.KEYDOWN):
            is_running = False

    (image, rect) = rotate_center(ship, rotation, 250, 250)

    window.blit(image, rect)

    SDL.draw.line(window, LINE_COLOR, (250, 50), (250, 450))
    SDL.draw.line(window, LINE_COLOR, (50, 250), (450, 250))
    gfxdraw.aacircle(window, 250, 250, 40, LINE_COLOR)
    gfxdraw.aacircle(window, 250, 250, 150, LINE_COLOR)

    rotation += 0.5

    SDL.display.update()

    clock.tick(FPS)

SDL.quit()
