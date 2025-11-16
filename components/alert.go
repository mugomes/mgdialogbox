// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package components

import (
	//"image/color"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// API EXIGIDA PELO SEU MAIN
func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(400, 100))
	win.CenterOnScreen()
	win.SetFixedSize(true)

	var lblIcon *canvas.Text
	color := color.Black
	if typeError {
		lblIcon = canvas.NewText("🛑", color)
	} else {
		lblIcon = canvas.NewText("✅", color)
	}
	lblIcon.TextSize = 70

	lblMessage := widget.NewLabel(message)
	lblMessage.Wrapping = fyne.TextWrapBreak

	btnClose := widget.NewButtonWithIcon(buttonOk, theme.ConfirmIcon(), func() {
		win.Close()
	})

	top := container.NewBorder(nil, nil, lblIcon, nil,
		container.NewVBox(lblMessage),
	)

	bottom := container.NewHBox(layout.NewSpacer(), btnClose, layout.NewSpacer())

	win.SetContent(
		container.NewBorder(
			top,
			bottom,
			nil, nil, nil,
		),
	)

	win.Show()
	
}