EXECUTABLE = main

define NEWLINE


endef

SOURCE_DIR = src
OBJECT_DIR = obj
INCLUDE_DIR = include

C_SOURCES = $(notdir $(wildcard $(SOURCE_DIR)/*.c))

HEADERS = $(wildcard $(INCLUDE_DIR)/*.h)

OBJS = $(patsubst %,$(OBJECT_DIR)/%,$(C_SOURCES:.c=.o))

CC = clang -std=c99 -pedantic -Wall

CFLAGS = -I/usr/include/SDL2 -D_REENTRANT -I./include
LIBS = -lSDL2 -lSDL2_image

$(info Source files: ${C_SOURCES} ${NEWLINE})
$(shell mkdir -p $(OBJECT_DIR))

all: $(EXECUTABLE)
	$(info ${NEWLINE}Sources compiled and linked successfully.${NEWLINE})
	
$(EXECUTABLE): $(OBJS)
	$(CC) -o $(EXECUTABLE) $(OBJS) $(CFLAGS) $(LIBS)

$(OBJECT_DIR)/%.o: $(SOURCE_DIR)/%.c $(HEADERS)
	$(CC) $(CFLAGS) -c $< -o $@

$(HEADERS): ;

.PHONY: clean

clean:
	rm -f $(EXECUTABLE) $(OBJECT_DIR)/*.o
