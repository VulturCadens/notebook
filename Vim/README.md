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

__:q!__ &nbsp; &nbsp; quit without saving  
__:qw__ &nbsp; &nbsp; quit and save  
__:qa__ &nbsp; &nbsp; quit all buffers

__a__ &nbsp; &nbsp; append  
__i__ &nbsp; &nbsp; insert  
__o__ &nbsp; &nbsp; open a new line and insert

__w__ &nbsp; move a word &nbsp; ->  
__b__ &nbsp; move a word &nbsp; <-

---

__r__ &nbsp; &nbsp; replace a single character  
__x__ &nbsp; &nbsp; delete a single character  

__v__ &nbsp; &nbsp; visual mode (mark text)

__dd__ &nbsp; &nbsp; delete a line (cut)  
__2dd__ &nbsp; delete two lines

__yy__ &nbsp; &nbsp; yank a line (copy)  
__yyp__ &nbsp; clone a line

__p__ &nbsp; &nbsp; paste  
__u__ &nbsp; &nbsp; undo

__>>__ / __<<__ &nbsp; &nbsp; indent / unindent

---

__:new__ &nbsp; &nbsp; open a new buffer (horizontal)  
__:vnew__ &nbsp; open a new buffer (vertical)

__:ls__ &nbsp; &nbsp; list buffers  
__:bd__ &nbsp; &nbsp; delete buffer  
__:bd!__ &nbsp; delete buffer without saving  
__:b2__ &nbsp; go to buffer two

---

__:reg__ &nbsp; show all registers  
__"2p__ &nbsp; &nbsp; paste from register two

__mx__ &nbsp; set mark "x"  
__'x__ &nbsp; &nbsp; jump to mark "x"
