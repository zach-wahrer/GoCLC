"""Python ASCII Art Helper for GoCLC."""

import sys
from pyfiglet import Figlet
from ascii_art import ascii_art


if __name__ == "__main__":
    res = ""
    for index, word in enumerate(sys.argv[2:]):
        res += word
        if index != len(sys.argv[2:]) - 1:
            res += " "
    print(ascii_art(res, sys.argv[1]))
