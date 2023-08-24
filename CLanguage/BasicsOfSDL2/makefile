#
# sudo apt install clang libsdl2-dev libsdl2-image-dev
#

EXECUTABLE := ./app.exe

.RECIPEPREFIX := ;

SOURCE_DIR := ./src
BUILD_DIR := ./build
INCLUDE_DIR := ./include

CC := gcc
WARNING := -Wall -Wextra -Werror -Wpedantic -std=c11
CFLAGS = -I/usr/include/SDL2 -D_REENTRANT -I./include
LIBS = -lSDL2 -lSDL2_image

HEADERS := $(wildcard  $(INCLUDE_DIR)/*.h)
OBJECTS := $(addprefix $(BUILD_DIR)/, $(addsuffix .o, \
		$(basename $(notdir $(wildcard  $(SOURCE_DIR)/*.c)))))

.PHONY: all
all: $(EXECUTABLE)
; @echo
; @echo + The program has been compiled and linked successfully.
; @echo

$(EXECUTABLE): $(OBJECTS)
; $(CC) -o $(EXECUTABLE) $(OBJECTS) $(LIBS)

$(BUILD_DIR)/%.o: $(SOURCE_DIR)/%.c $(HEADERS)
; $(CC) -c $(WARNING) $(CFLAGS) $< -o $@

.PHONY: clean
clean:
; @rm -fv $(EXECUTABLE)
; @rm -fv $(BUILD_DIR)/*.o
; @echo
; @echo + All the object files and executables has been deleted.
; @echo
