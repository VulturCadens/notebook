EXECUTABLE := executable

.RECIPEPREFIX := ;

MAKE_DIR  := mkdir -p

INCLUDE_DIR := include
SOURCE_DIR := source
BUILD_DIR := build

CC := g++ -std=c++20
FLAGS := `sdl2-config --cflags` -I./$(INCLUDE_DIR)
LIBS := `sdl2-config --libs`
WARNING := -Wall -Wextra -Werror -Wpedantic

SOURCES := $(wildcard $(SOURCE_DIR)/*.cpp)
OBJECTS := $(addprefix $(BUILD_DIR)/, $(notdir $(SOURCES:.cpp=.o)))
HEADERS := $(wildcard $(INCLUDE_DIR)/*.h)

# $(info $(SOURCES))
# $(info $(OBJECTS))

$(shell $(MAKE_DIR) $(BUILD_DIR))

all: $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
; $(CC) $(OBJECTS) $(LIBS) -o $@

$(BUILD_DIR)/%.o: $(SOURCE_DIR)/%.cpp $(HEADERS)
; $(CC) $(FLAGS) $(WARNING) -c $< -o $@

clean:
; rm -f $(EXECUTABLE) $(OBJECTS)

.PHONY: all clean
