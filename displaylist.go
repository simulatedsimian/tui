package tui

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/rect"
)

func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		x++
	}
}

func printAtDef(x, y int, s string) {
	printAt(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func clearRect(r rect.Rectangle, c rune, fg, bg termbox.Attribute) {
	w, h := termbox.Size()
	sz := rect.FromSize(rect.Vec{w, h})

	toClear, ok := rect.Intersection(r, sz)
	if ok {
		for y := toClear.Min.Y; y < toClear.Max.Y; y++ {
			for x := toClear.Min.X; x < toClear.Max.X; x++ {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

func clearRectDef(r rect.Rectangle) {
	clearRect(r, '.', termbox.ColorDefault, termbox.ColorDefault)
}

type Focusable interface {
	GiveFocus()
}

type Drawable interface {
	Draw()
}

type Controlable interface {
	HandleInput(k termbox.Key, r rune)
}

type DisplayList struct {
	list       []Drawable
	focusIndex int
}

func (dl *DisplayList) AddElement(elem Drawable) {
	dl.list = append(dl.list, elem)
}

func (dl *DisplayList) Draw() {
	w, h := termbox.Size()
	clearRectDef(rect.FromSize(rect.Vec{w, h}))

	for _, elem := range dl.list {
		elem.Draw()
	}
}

func (dl *DisplayList) NextFocus() {
	if dl.list != nil && len(dl.list) > 0 {
		for {
			dl.focusIndex++
			if dl.focusIndex >= len(dl.list) {
				dl.focusIndex = 0
			}

			if f, ok := dl.list[dl.focusIndex].(Focusable); ok {
				f.GiveFocus()
				break
			}
		}
	}
}

func (dl *DisplayList) PrevFocus() {
	if dl.list != nil && len(dl.list) > 0 {
		for {
			dl.focusIndex--
			if dl.focusIndex < 0 {
				dl.focusIndex = len(dl.list)
			}

			if f, ok := dl.list[dl.focusIndex].(Focusable); ok {
				f.GiveFocus()
				break
			}
		}
	}
}

func (dl *DisplayList) HandleInput(k termbox.Key, r rune) {

	if dl.list != nil && len(dl.list) > 0 {
		if k == termbox.KeyTab {
			dl.NextFocus()
		} else {
			if c, ok := dl.list[dl.focusIndex].(Controlable); ok {
				c.HandleInput(k, r)
			}

		}
	}
}

type TUIElement struct {
	rect.Rectangle
}
