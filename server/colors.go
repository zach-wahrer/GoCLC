package server

import (
	"math/rand"
	"time"
)

var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
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
