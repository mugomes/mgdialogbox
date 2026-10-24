// Copyright (C) 2025-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: MIT

// Site: https://www.profmugomes.com.br

package mgdialogbox

import (
	"fyne.io/fyne/v2"
	"github.com/profmugomes/mgdialogbox/v3/components"
)

func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string, onClosed func()) {
	components.NewAlert(a, title, message, typeError, buttonOk, onClosed)
}

func NewConfirm(a fyne.App, title, message string, buttons []string, OnResult func(int)) {
	components.NewConfirm(a, title, message, buttons, OnResult)
}

func NewOpenFile(a fyne.App, title string, exts []string, multiselect bool, onSelect func([]string)) {
	components.NewOpenFile(a, title, exts, multiselect, onSelect)
}

func NewSaveFile(a fyne.App, title string, exts []string, onSelect func(string)) {
	components.NewSaveFile(a, title, exts, onSelect)
}

func NewSelectDirectory(a fyne.App, title string, multiselect bool, onSelect func([]string)) {
	components.NewSelectDirectory(a, title, multiselect, onSelect)
}
