module client

go 1.13

require goclctest v1.0.0

replace goclctest => ../goclctest

require (
	github.com/jroimartin/gocui v0.4.0
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
	server v1.0.0
)

replace server => ../server
