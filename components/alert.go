// Copyright (C) 2025-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: MIT

// Site: https://www.profmugomes.com.br

package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/profmugomes/mgsmartflow/v3"
)

func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string, onClosed func()) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(400, 100))
	win.CenterOnScreen()
	win.SetFixedSize(true)

	flow := mgsmartflow.New()

	var lblIcon *canvas.Text
	color := color.Black
	if typeError {
		lblIcon = canvas.NewText("🛑", color)
	} else {
		lblIcon = canvas.NewText("✅", color)
	}
	lblIcon.TextSize = 70

	lblMessage := widget.NewLabel(message)
	lblMessage.Wrapping = fyne.TextWrapWord

	flow.AddColumn(lblIcon, lblMessage)
	flow.Resize(lblIcon, 79, lblIcon.MinSize().Height+57)
	flow.Move(lblIcon, 12, 7)

	btnClose := widget.NewButtonWithIcon(buttonOk, theme.ConfirmIcon(), func() {
		win.Close()

		if onClosed != nil {
			onClosed()
		}
	})

	flow.AddRow(container.NewHBox(layout.NewSpacer(), btnClose, layout.NewSpacer()))

	win.SetContent(flow.Container)
	win.Show()
}
