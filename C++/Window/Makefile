EXECUTABLE := window

.RECIPEPREFIX := ;

INCLUDE_DIR := ./include
HEADERS := $(wildcard  $(INCLUDE_DIR)/*.h)

CC := g++
FLAGS := `sdl2-config --cflags` -I./$(INCLUDE_DIR)
LIBS := `sdl2-config --libs`
WARNING := -Wall -Wextra -Werror -Wpedantic

SOURCES := $(wildcard *.cpp)
OBJECTS := $(SOURCES:.cpp=.o)


all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
; $(CC) -o $(EXECUTABLE) $(OBJECTS) $(LIBS)

$(OBJECTS): $(SOURCES) $(HEADERS)
; $(CC) -c $(SOURCES) $(FLAGS) $(WARNING)

clean:
; rm -f $(EXECUTABLE) $(OBJECTS)

.PHONY: all clean
