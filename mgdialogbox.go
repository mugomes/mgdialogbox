// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package mgdialogbox

import (
	"fyne.io/fyne/v2"
	"github.com/mugomes/mgdialogbox/components"
)

func NewAlert(a fyne.App, title string, message string, typeError bool, buttonOk string) *components.MessageDialog {
	return components.NewAlert(a, title, message, typeError, buttonOk)
}

func NewOpenFile(a fyne.App, title string, exts []string, multiselect bool, onSelect func([]string)) *components.FileDialogOpen {
	return components.NewOpenFile(a, title, exts, multiselect, onSelect)
}