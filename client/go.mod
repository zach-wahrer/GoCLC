module client

go 1.13

require goclctest v1.0.0

replace goclctest => ../goclctest

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807
	golang.org/x/sys v0.0.0-20200727154430-2d971f7391a4 // indirect
	server v1.0.0
)

replace server => ../server
