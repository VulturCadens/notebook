# Python

## Setting Up Environment

```shell
$ mkdir PROJECT
$ cd PROJECT

$ python -m venv env
$ source env/bin/activate

(env) $ python --version

(env) $ python -m pip install PACKAGE

.
.
.

(env) $ deactivate
```

## The Standard Types

### Numbers

* Int **123**

* Float **123.00**

* Complex **12+3j**

### Boolean

* Bool **True**

### Sequence of Characters or Bytes

* String **"Example"**

* Bytearray "mutable" **bytearray([123, 23, 3])**

* Bytes "immutable" **bytes([123, 23, 3])**

### Collections

* List "mutable, ordered" **["Cat", 123]**

* Tuple "immutable, ordered" **("Cat", 123)**

* Set "mutable, unordered, no duplicate elements" **{"Cat", 123)**

* Frozenset "immutable, unordered, no duplicate elements" **frozenset(["Cat", 123])**

* Dictionary "mutable" **{"Cat": "Mirre", "ID": 123}**

## Modules and Packages

The \_\_init\_\_.py file is place to execute any initialization code required to use the package. It's not necessary create an empty \_\_init\_\_.py file. [Implicit Namespace Packages](https://www.python.org/dev/peps/pep-0420/) (since Python 3.3) allow for the creation of a package without any \_\_init\_\_.py file.

```shell
.
├── main.py
└── PACKAGE
    ├── __init__.py
    ├── MODULE_A.py # FUNC1, FUNC2 ...
    └── MODULE_B.py # FUNCA, FUNCB ...
```

```python
# main.py

import PACKAGE.MODULE_A
PACKAGE.MODULE_A.FUNC1()

import PACKAGE.MODULE_B as NAME
NAME.FUNCA()

from PACKAGE.MODULE_A import FUNC1 [, FUNC2] [, FUNC3] ...
FUNC1()
[FUNC2()]...

from PACKAGE.MODULE_B import FUNCB as NAME
NAME()
```

_"The import statement uses the following convention: if a package’s \_\_init\_\_.py code defines a list named \_\_all\_\_, it is taken to be the list of module names that should be imported when from package import * is encountered."_

```python
# __init__.py

__all__ = ["MODULE_A", "MODULE_B"]
```

```python
# main.py

from PACKAGE import *

MODULE_A.FUNC1()
MODULE_A.FUNCA()
```

Using an \_\_init\_\_.py to import functions from all modules.

```python
# __init__.py

from . MODULE_A import *
from . MODULE_B import *
```

```python
# main.py

import PACKAGE
PACKAGE.FUNC1()
PACKAGE.FUNCA()

import PACKAGE as NAME
NAME.FUNC1()
NAME.FUNCA()
```

## The @classmethod Decorator

The class method can access class attributes, but not the instance attributes.

Can be called only using **class.method()**.

```python
class Button:

    buttons = []

    def __init__(self, image, position, function):
        self.image = SDL.image.load(image).convert_alpha()
        self.position = position

        self.function = function

        self.width = self.image.get_width()
        self.height = self.image.get_height()

        Button.buttons.append(self)

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
```

## The @staticmethod Decorator

The static method cannot access the class attributes or the instance attributes.

Can be called **class.method()** and also **object.method()**.

```python
class Cat:
    def __init__(self, name):
        self.name = name

    def getName(self):
        return self.name

    @staticmethod
    def getSpecies():
        return "Cat"
```
