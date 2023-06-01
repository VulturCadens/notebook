import curses
import os

INVISIBLE = 0
VISIBLE = 1

KEY_ESC = 27

def init_screen():
    os.environ.setdefault("ESCDELAY", "0")

    s = curses.initscr()
    s.keypad(True)
    s.refresh()

    curses.noecho()
    curses.cbreak()
    curses.curs_set(INVISIBLE)

    return s

def loop(screen):
    win = curses.newwin(20, 30, 5, 10) # height, width, y, x
 
    win.addstr(1, 1, "Notebook", curses.A_BOLD)
    win.border()
    win.refresh()

    while True:
        c = screen.getch()

        if c == ord("q"):
            break
        elif c == KEY_ESC:
            break

if __name__ == "__main__":
    screen = init_screen()

    loop(screen)

    curses.curs_set(VISIBLE)
    curses.nocbreak()
    curses.echo()
    curses.endwin()

