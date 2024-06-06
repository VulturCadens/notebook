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
HEADERS := $(wildcard $(INCLUDE_DIR)/*.h)

LIST := $(SOURCES:.cpp=.o)
TEMP := $(notdir $(LIST))
OBJECTS := $(BUILD_DIR)/$(TEMP)

all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
; $(CC) $(OBJECTS) $(LIBS) -o $(EXECUTABLE)

$(OBJECTS): $(SOURCES) $(HEADERS)
; $(CC) -c $(SOURCES) $(FLAGS) $(WARNING) -o $(OBJECTS)

clean:
; rm -f $(EXECUTABLE) $(OBJECTS)

.PHONY: all clean
