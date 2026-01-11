// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewConfirm(a fyne.App, title, message string, buttons []string, OnResult func(int)) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(400, 100))
	win.CenterOnScreen()
	win.SetFixedSize(true)

	var lblIcon *canvas.Text
	color := color.Black
	lblIcon = canvas.NewText("💬", color)
	lblIcon.TextSize = 70

	lblMessage := widget.NewLabel(message)
	lblMessage.Wrapping = fyne.TextWrapBreak

	var btns []fyne.CanvasObject
	for i, btn := range buttons {
		nBtn := widget.NewButton(btn, func() {
			OnResult(i)
			win.Close()
		})

		btns = append(btns, nBtn)
	}

	top := container.NewBorder(nil, nil, lblIcon, nil,
		container.NewVBox(lblMessage),
	)

	bottom := container.NewHBox(layout.NewSpacer(), container.NewHBox(btns...), layout.NewSpacer())

	win.SetContent(
		container.NewBorder(
			top,
			bottom,
			nil, nil, nil,
		),
	)

	win.Show()
}
