import sys
import pygame as SDL

EXIT_FAILURE = 1
EXIT_SUCCESS = 0

SCREEN_WIDTH = 800
SCREEN_HEIGHT = 600

BACKGROUND = (20, 80, 20)
FPS = 30


class BOB(SDL.sprite.Sprite):
    def __init__(self, image):
        super().__init__()

        self.image = image
        self.rect = self.image.get_rect()


def main() -> int:
    window = SDL.display.set_mode(size=(SCREEN_WIDTH, SCREEN_HEIGHT), vsync=1)

    is_running = True

    try:
        fire = SDL.image.load("fire.png").convert_alpha()

    except FileNotFoundError as error:
        print("Error: {}".format(error))
        return EXIT_FAILURE

    clock = SDL.time.Clock()

    SDL.display.set_caption("Simple Sprite")
    SDL.display.set_icon(fire)

    sprite = BOB(fire)

    # Rect(x, y, width, height) -> Rect
    # Rect((x, y), (width, height)) -> Rect
    # Rect(object) -> Rect
    sprite.rect.x = 100
    sprite.rect.y = 200

    sprites = SDL.sprite.Group()
    sprites.add(sprite)

    direction = 2

    background = SDL.Surface((64, 64))
    background.fill(BACKGROUND)

    window.fill(BACKGROUND)
    SDL.display.update()

    while is_running:

        for event in SDL.event.get():
            if event.type in (SDL.QUIT, SDL.KEYDOWN):
                is_running = False

        window.blit(background, sprite.rect)
        sprites.draw(window)

        SDL.display.update(sprite.rect)

        for sprite in sprites:
            sprite.rect.x += direction

            if sprite.rect.x == 720 or sprite.rect.x == 20:
                direction *= -1

        clock.tick(FPS)

    return EXIT_SUCCESS


if __name__ == "__main__":
    SDL.init()

    exit_status = main()

    SDL.quit()

    sys.exit(exit_status)
