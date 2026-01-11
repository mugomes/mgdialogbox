// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

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