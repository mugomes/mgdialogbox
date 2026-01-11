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

	"github.com/mugomes/mgsmartflow"
)

func NewConfirm(a fyne.App, title, message string, buttons []string, OnResult func(int)) {
	win := a.NewWindow(title)
	win.Resize(fyne.NewSize(400, 100))
	win.CenterOnScreen()
	win.SetFixedSize(true)

	flow := mgsmartflow.New()

	var lblIcon *canvas.Text
	color := color.Black
	lblIcon = canvas.NewText("💬", color)
	lblIcon.TextSize = 70

	lblMessage := widget.NewLabel(message)
	lblMessage.Wrapping = fyne.TextWrapWord

	flow.AddColumn(lblIcon, lblMessage)
	flow.SetResize(lblIcon, fyne.NewSize(79, lblIcon.MinSize().Height))
	flow.SetMove(lblIcon, fyne.NewPos(12,7))

	var btns []fyne.CanvasObject
	for i, btn := range buttons {
		nBtn := widget.NewButton(btn, func() {
			OnResult(i)
			win.Close()
		})

		btns = append(btns, nBtn)
	}

	flow.AddRow(container.NewHBox(layout.NewSpacer(), container.NewHBox(btns...), layout.NewSpacer()))

	win.SetContent(flow.Container)
	win.Show()
}
