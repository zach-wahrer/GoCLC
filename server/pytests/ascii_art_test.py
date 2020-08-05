from ascii_art.ascii_art import ascii_art


def test_ascii_art():
    font = "rectangles"
    want = """| |_ ___ ___| |_
|  _| -_|_ -|  _|
|_| |___|___|_|
"""

    got = ascii_art("test", font)
    for line in want.split("\n"):
        assert line in got
