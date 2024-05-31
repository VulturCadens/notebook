EXECUTABLE := window

.RECIPEPREFIX := ;

CC := g++
FLAGS := `sdl2-config --cflags`
LIBS := `sdl2-config --libs`
WARNING := -Wall -Wextra -Werror -Wpedantic

SOURCES := $(wildcard *.cpp)
OBJECTS := $(SOURCES:.cpp=.o)

all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
; $(CC) -o $(EXECUTABLE) $(OBJECTS) $(LIBS)

$(OBJECTS): $(SOURCES)
; $(CC) -c $(SOURCES) $(FLAGS) $(WARNING)

clean:
; rm -f $(EXECUTABLE) $(OBJECTS)

.PHONY: all clean
