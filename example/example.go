package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
	"github.com/simulatedsimian/tui"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	doQuit := false

	logDisp := tui.MakeScrollingTextOutput(rect.XYWH(1, 20, 80, 10))
	cmdInput := tui.MakeTextInputField(0, 18, func(cmd string) {
		logDisp.WriteLine(cmd)
		if cmd == "q" {
			doQuit = true
		}
	})

	dl := tui.DisplayList{}
	dl.AddElement(cmdInput)
	dl.AddElement(logDisp)
	dl.AddElement(tui.MakeStaticText(rect.XYWH(0, 0, 1, 1), "StaticText"))

	dl.Draw()
	termbox.Flush()

	for !doQuit {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			dl.HandleInput(ev.Key, ev.Ch)
			dl.Draw()
			termbox.Flush()
		}

		if ev.Type == termbox.EventResize {
			termbox.Flush()
		}
	}

}
