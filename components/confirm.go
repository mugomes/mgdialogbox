// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/profmugomes/mgsmartflow/v2"
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
	flow.Resize(lblIcon, 79, lblIcon.MinSize().Height + 57)
	flow.Move(lblIcon, 12,7)

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
