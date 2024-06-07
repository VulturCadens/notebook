#!/usr/bin/env python

import os
import sys
import subprocess

EXECUTABLE = "executable"
SPACE = " "
EMPTY = ""

INCLUDE_DIR = "./include"
SOURCE_DIR = "./source"
BUILD_DIR = "./build"

CC = "g++ -std=c++20"
FLAGS = "`sdl2-config --cflags` -I" + INCLUDE_DIR
LIBS = "`sdl2-config --libs`"
WARNING = "-Wall -Wextra -Werror -Wpedantic"


def run_command(c):
    command = (SPACE.join(c))

    print("  + Running: " + command)

    result = subprocess.run([command], shell=True,
                            capture_output=True, text=True)

    if result.stdout != EMPTY:
        lines = result.stdout.splitlines()

        for line in lines:
            print("\t" + line)

        if result.returncode != 0:
            sys.exit("ERROR: " + result.stderr)


def main():
    sources = []
    objects = []

    for file in os.listdir(SOURCE_DIR):
        if os.path.isfile(os.path.join(SOURCE_DIR, file)):
            sources.append(file)

    print("\nCompilation:")

    for source in sources:
        source_file = os.path.join(SOURCE_DIR, source)
        object_file = os.path.join(BUILD_DIR, source.replace(".cpp", ".o"))

        run_command([CC, "-c", source_file, WARNING, FLAGS, "-o", object_file])

        objects.append(object_file)

    print("\nLinking:")

    run_command([CC, "-o", EXECUTABLE, SPACE.join(objects), LIBS])


def clean():
    print("\nCleaning:")
    run_command(["rm -fv", BUILD_DIR + "/*.o", EXECUTABLE])


if __name__ == "__main__":
    if len(sys.argv) == 1:
        main()

    elif sys.argv[1] == "clean":
        clean()

    else:
        sys.exit("\nERROR: An argument is invalid.\n")

    print()
