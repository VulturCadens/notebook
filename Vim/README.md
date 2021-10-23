# Vim

## Install NERDTree

```consolemk
$ mkdir -p ~/.vim/pack/vendor/start

$ git clone --depth 1 \
  https://github.com/preservim/nerdtree.git \
  ~/.vim/pack/vendor/start/nerdtree

$ vim

:NERDTree
```

## Basic Commands

| Key | Action |
| :--|:--|
|__:q!__|quit without saving|
|__:qw__|quit and save|
|__:qa!__|quit all buffers without saving|
|__:Explore__|open the file explorer|
|__a__|append|
|__i__|insert|
|__o__|open a new line and insert|
|__w__|move  foreward one word|
|__b__|move back one word|
|__r__|replace a single character|
|__x__|delete a single character|
|__ci"__|change inside quotes|
|__ciw__|change a word|
|__dd__|delete a line (cut)|
|__2dd__|delete two lines|
|__di(__|delete inside brackets|
|__yy__|yank a line (copy)|
|__yyp__|clone a line|
|__p__|paste|
|__u__|undo|
|__v__|visual mode (mark text)|
|__>>__ / __<<__|indent / unindent|
|__:new__|open a new window (horizontal)|
|__:vnew__|open a new window (vertical)|
|__ctrl + w w__|toggle between open windows|
|__ctrl + w q__|close a window|
|__:ls__|list buffers|
|__:bd__|delete buffer|  
|__:bd!__|delete buffer without saving|
|__:b2__|go to buffer two|
|__:reg__|show all registers|
|__"2p__|paste from register two|
|__mx__|set mark "x"|
|__'x__|jump to mark "x"|
