import pygame as SDL


SDL.init()

BACKGROUND = (20, 80, 20)

FPS = 30

window = SDL.display.set_mode(size=(500, 500), vsync=1)

is_running = True

fire = SDL.image.load("fire.png").convert_alpha()
clock = SDL.time.Clock()

SDL.display.set_caption("Simple Sprite")
SDL.display.set_icon(fire)


class BOB(SDL.sprite.Sprite):
    def __init__(self, image):
        super().__init__()

        self.image = image
        self.rect = self.image.get_rect()


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

        if sprite.rect.x == 420 or sprite.rect.x == 20:
            direction *= -1

    clock.tick(FPS)

SDL.quit()
