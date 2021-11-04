import sys
import os
import sqlite3


def main() -> int:
    if os.path.exists("database.db"):
        os.remove("database.db")

    try:
        conn = sqlite3.connect("database.db") # The file is created if it doesn't exist.

        c = conn.cursor()

        c.execute("CREATE TABLE movies(title TEXT, year INTEGER, director TEXT)")

        c.execute("INSERT INTO movies VALUES ('Dune', 2021, 'Denis Villeneuve')")
        c.execute("INSERT INTO movies VALUES ('Murder Mystery', 2019, 'Kyle Newacheck')")

        rows = c.execute("SELECT * FROM movies").fetchall()
        print(rows)

        rows = c.execute(
            "SELECT year, director FROM movies WHERE title = ?", ("Dune",)).fetchall()

        for row in rows:
            print("Year {} and director {}".format(row[0], row[1]))

        c.close()

    except sqlite3.Error as error:
        print("The SQLite error:", error)
        return 1

    finally:
        if conn:
            conn.close()
            print("The SQLite connection is closed.")

    return 0


if __name__ == "__main__":
    sys.exit(main())
