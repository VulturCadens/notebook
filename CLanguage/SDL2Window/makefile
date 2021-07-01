EXECUTABLE = main

SRCDIR=src
OBJDIR=obj

SRCS = main.c init.c

_OBJS = $(SRCS:.c=.o)
OBJS = $(patsubst %,$(OBJDIR)/%,$(_OBJS))

CC=clang

CFLAGS=-I/usr/include/SDL2 -D_REENTRANT -I./include
LIBS=-lSDL2

all: $(EXECUTABLE)
	@echo OK

$(EXECUTABLE): $(OBJS)
	$(CC) -o $(EXECUTABLE) $(OBJS) $(CFLAGS) $(LIBS)

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	@mkdir -p $(OBJDIR)
	$(CC) $(CFLAGS) -c $< -o $@

.PHONY: clean

clean:
	rm -f $(EXECUTABLE) $(OBJDIR)/*.o
