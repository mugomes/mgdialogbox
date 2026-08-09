// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package controls

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type resizableLayout struct {
	last     fyne.Size
	onChange func(fyne.Size)
}

func (l *resizableLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if size != l.last {
		l.last = size
		if l.onChange != nil {
			l.onChange(size)
		}
	}
}

func (l *resizableLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(1, 1)
}

func OnResize(win fyne.Window, fn func(fyne.Size)) {
	if win == nil || win.Canvas() == nil {
		return
	}

	lytResize := &resizableLayout{
		onChange: fn,
	}

	passThrough := container.New(lytResize, layout.NewSpacer())
	win.SetContent(passThrough)
}