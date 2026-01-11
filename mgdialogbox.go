// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package mgdialogbox

import (
	"fyne.io/fyne/v2"
	"github.com/mugomes/mgdialogbox/components"
)

func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string) {
	components.NewAlert(a, title, message, typeError, buttonOk)
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
