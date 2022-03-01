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
