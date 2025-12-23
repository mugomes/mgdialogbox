// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package controls

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// type WindowResizeObserver struct {
// 	last     fyne.Size
// 	OnChange func(size fyne.Size)
// }

// func NewWindowResizeObserver(onChange func(size fyne.Size), objects ...fyne.CanvasObject) *fyne.Container {
// 	obs := &WindowResizeObserver{
// 		OnChange: onChange,
// 	}
// 	return container.New(obs, objects...)
// }

// func (o *WindowResizeObserver) Layout(objects []fyne.CanvasObject, size fyne.Size) {
// 	if size != o.last {
// 		o.last = size
// 		if o.OnChange != nil {
// 			o.OnChange(size)
// 		}
// 	}

// 	for _, obj := range objects {
// 		obj.Resize(obj.Size())
// 		obj.Move(obj.Position())
// 	}
// }

// func (o *WindowResizeObserver) MinSize(objects []fyne.CanvasObject) fyne.Size {
// 	return fyne.NewSize(1, 1)
// }

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