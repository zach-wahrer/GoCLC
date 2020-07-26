package server

import (
	"math/rand"
	"time"
)

var (
	colorReset  = "\u001b[0m"
	colorBold   = "\u001b[1m"
	colorRed    = "\u001b[31m"
	colorGreen  = "\u001b[32m"
	colorYellow = "\u001b[33m"
	colorBlue   = "\u001b[34m"
	colorPurple = "\u001b[35m"
	colorCyan   = "\u001b[36m"
	colorWhite  = "\u001b[37m"
)
var colors = []string{
	colorRed,
	colorGreen,
	colorYellow,
	colorBlue,
	colorPurple,
	colorCyan,
	colorWhite}

func randomColor() string {
	rand.Seed(time.Now().UnixNano())
	return colors[rand.Intn(len(colors)-1)]
}
