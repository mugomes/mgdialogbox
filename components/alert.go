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

type MessageDialog struct {
	a         fyne.App
	title     string
	message   string
	typeError bool
	buttonOk  string
}

// API EXIGIDA PELO SEU MAIN
func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string) *MessageDialog {
	dlg := &MessageDialog{
		a:         a,
		title:     title,
		message:   message,
		typeError: typeError,
		buttonOk:  buttonOk,
	}
	return dlg
}

func (d *MessageDialog) Show() {
	win := d.a.NewWindow(d.title)
	win.Resize(fyne.NewSize(400, 100))
	win.CenterOnScreen()
	win.SetFixedSize(true)

	var lblIcon *canvas.Text
	color := color.Black
	if d.typeError {
		lblIcon = canvas.NewText("✅", color)
	} else {
		lblIcon = canvas.NewText("🛑", color)
	}
	lblIcon.TextSize = 70

	lblMessage := widget.NewLabel(d.message)
	lblMessage.Wrapping = fyne.TextWrapBreak

	btnClose := widget.NewButtonWithIcon(d.buttonOk, theme.ConfirmIcon(), func() {
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
