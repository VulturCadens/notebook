EXECUTABLE := window

.RECIPEPREFIX := ;

CC := g++
CFLAGS := `sdl2-config --libs --cflags` -Wall -lm

SRCS := window.cpp
OBJS := $(SRCS:.cpp=.o)

all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJS)
; $(CC) -o $(EXECUTABLE) $(OBJS) $(CFLAGS)

clean:
; rm -f $(EXECUTABLE) $(OBJS)

.PHONY: all clean
