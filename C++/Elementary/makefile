EXECUTABLE := executable

.RECIPEPREFIX := ;

INCLUDE_DIR := ./include
SOURCE_DIR := ./source
BUILD_DIR := ./build

CC := g++ -std=c++20
FLAGS := `sdl2-config --cflags` -I./$(INCLUDE_DIR)
LIBS := `sdl2-config --libs`
WARNING := -Wall -Wextra -Werror -Wpedantic

SOURCES := $(wildcard $(SOURCE_DIR)/*.cpp)
OBJECTS := $(BUILD_DIR)/$(notdir $(SOURCES:.cpp=.o))
HEADERS := $(wildcard $(INCLUDE_DIR)/*.h)


all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
; $(CC) $(OBJECTS) $(LIBS) -o $(EXECUTABLE)

$(OBJECTS): $(SOURCES) $(HEADERS)
; $(CC) -c $(SOURCES) $(FLAGS) $(WARNING) -o $(OBJECTS)

clean:
; rm -f $(EXECUTABLE) $(OBJECTS)

.PHONY: all clean
