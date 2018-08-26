package inter

import "fmt"

type Player interface {
	Play()
}

type Mpv struct {
	cmd string
}
