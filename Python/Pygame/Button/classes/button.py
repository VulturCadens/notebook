import pygame as SDL


class Button:

    buttons = []

    @classmethod
    def click(cls, position):
        for button in cls.buttons:

            x = button.position[0]
            y = button.position[1]

            width = button.width
            height = button.height

            if position[0] > x and position[0] < x + width:
                if position[1] > y and position[1] < y + height:
                    button.function()

    @classmethod
    def draw(cls, target):
        for button in cls.buttons:
            target.blit(button.image, button.position)

    def __init__(self, image, position, function):
        self.image = SDL.image.load(image).convert_alpha()
        self.position = position

        self.function = function

        self.width = self.image.get_width()
        self.height = self.image.get_height()

        Button.buttons.append(self)
