from pyfiglet import Figlet


def ascii_art(text: str, font: str) -> str:
    return Figlet(font=font).renderText(text)
