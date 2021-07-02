EXECUTABLE = main

define NEWLINE


endef

SOURCE_DIR = src
OBJECT_DIR = obj

C_SOURCES = $(notdir $(wildcard $(SOURCE_DIR)/*.c))

OBJS = $(patsubst %,$(OBJECT_DIR)/%,$(C_SOURCES:.c=.o))

CC = clang

CFLAGS = -I/usr/include/SDL2 -D_REENTRANT -I./include
LIBS = -lSDL2

$(info Source files: ${C_SOURCES} ${NEWLINE})
$(shell mkdir -p $(OBJECT_DIR))

all: $(EXECUTABLE)
	$(info ${NEWLINE}Sources compiled and linked successfully.${NEWLINE})
	
$(EXECUTABLE): $(OBJS)
	$(CC) -o $(EXECUTABLE) $(OBJS) $(CFLAGS) $(LIBS)

$(OBJECT_DIR)/%.o: $(SOURCE_DIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@

.PHONY: clean

clean:
	rm -f $(EXECUTABLE) $(OBJECT_DIR)/*.o
