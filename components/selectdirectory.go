// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package components

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/mugomes/mgdialogbox/controls"
	"github.com/mugomes/mgsettings/v3"
	"github.com/mugomes/mgsmartflow"
)

type DialogSelectDirectory struct {
	a           fyne.App
	title       string
	multiSelect bool
	onSelect    func([]string)

	lastDir string
}

var selectDirectoryMGconfig, _ = mgsettings.Load("/tmp/mgdialogbox", false)

func NewSelectDirectory(a fyne.App, title string, multiselect bool, onSelect func([]string)) {
	dlg := &DialogSelectDirectory{
		a:           a,
		title:       title,
		multiSelect: multiselect,
		onSelect:    onSelect,
	}
	dlg.sdLoadLastDir()
	dlg.sdShowSelectDirectory()
}

func (d *DialogSelectDirectory) sdShowSelectDirectory() {
	win := d.a.NewWindow(d.title)
	win.Resize(fyne.NewSize(740, 520))
	win.CenterOnScreen()

	dir := d.lastDir
	if dir == "" {
		dir, _ = os.UserHomeDir()
	}

	pathLabel := widget.NewLabel(dir)
	search := widget.NewEntry()
	search.SetPlaceHolder("🔍 Search...")

	files := d.sdListDir(dir)
	filtered := files
	selected := map[int]bool{}

	list := widget.NewList(
		func() int { return len(filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, o fyne.CanvasObject) {
			row := filtered[id]
			name := row.Name()

			if row.IsDir() {
				if selected[id] {
					o.(*widget.Label).SetText("✔ " + name)
				} else {
					o.(*widget.Label).SetText("📁 " + name)
				}
			}
		},
	)

	var (
		click         int = 1
		timeClick     time.Time
		durationClick = 400 * time.Millisecond
	)

	list.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(filtered) {
			return
		}
		f := filtered[id]

		// Abrir diretório
		if f.IsDir() {
			// Selecionar Diretório
			if d.multiSelect {
				selected[id] = !selected[id]
			} else {
				for k := range selected {
					delete(selected, k)
				}
				selected[id] = !selected[id]
			}

			// Remove a selação
			list.Unselect(id)

			// Clique duplo
			click = click + 1
			now := time.Now()
			if now.Sub(timeClick) > durationClick {
				click = 1
			}
			timeClick = now
			if click == 2 {
				search.SetText("")
				dir = filepath.Join(dir, f.Name())
				pathLabel.SetText(dir)
				files = d.sdListDir(dir)
				filtered = d.sdApplyFilter(files, search.Text)
				selected = map[int]bool{}
				list.Refresh()
				click = 1 // reset clique duplo
			}

			return
		}

		list.Refresh()
	}

	// BOTÃO VOLTAR
	btnBack := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		parent := filepath.Dir(dir)
		if parent != dir {
			dir = parent
			pathLabel.SetText(dir)
			files = d.sdListDir(dir)
			filtered = d.sdApplyFilter(files, search.Text)
			selected = map[int]bool{}
			list.Refresh()
		}
	})

	// Botão Selecionar
	btnSelect := widget.NewButtonWithIcon("Selecionar", theme.ConfirmIcon(), func() {
		var out []string
		for i := range selected {
			f := filtered[i]
			if f.IsDir() && selected[i] {
				out = append(out, filepath.Join(dir, f.Name()))
			}
		}

		if len(out) > 0 {
			d.sdSaveLastDir(dir)
			d.onSelect(out)
			win.Close()
		}
	})

	// BUSCA
	search.OnChanged = func(txt string) {
		filtered = d.sdApplyFilter(files, txt)
		selected = map[int]bool{}
		list.Refresh()
	}

	flow := mgsmartflow.New()

	// LAYOUT
	flow.AddColumn(btnBack, container.NewVBox(pathLabel, search))
	flow.SetResize(btnBack, fyne.NewSize(68, 79))

	flow.AddRow(list)
	flow.AddRow(btnSelect)

	controls.OnResize(win, func(size fyne.Size) {
		flow.SetResize(list, fyne.NewSize(size.Width, size.Height-137))
	})

	win.Canvas().Overlays().Add(flow.Container)
	win.Show()
}

//////////////////////////////////////////////////////////////
// FUNÇÕES AUXILIARES
//////////////////////////////////////////////////////////////

func (d *DialogSelectDirectory) sdListDir(path string) []fs.FileInfo {
	entries, _ := os.ReadDir(path)

	var list []fs.FileInfo

	for _, e := range entries {
		info, err := e.Info()
		if info.IsDir() {
			if err != nil {
				continue
			}
			list = append(list, info)
		}
	}

	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		if a.IsDir() != b.IsDir() {
			return a.IsDir()
		}
		return strings.ToLower(a.Name()) < strings.ToLower(b.Name())
	})

	return list
}

func (d *DialogSelectDirectory) sdApplyFilter(files []fs.FileInfo, query string) []fs.FileInfo {
	if query == "" {
		return files
	}

	q := strings.ToLower(query)
	var out []fs.FileInfo

	for _, f := range files {
		if strings.Contains(strings.ToLower(f.Name()), q) {
			out = append(out, f)
		}
	}

	return out
}

//////////////////////////////////////////////////////////////
// ÚLTIMO DIRETÓRIO
//////////////////////////////////////////////////////////////

func (d *DialogSelectDirectory) sdSaveLastDir(dir string) {
	openFileMGconfig.SetString("lastdir", dir)
	openFileMGconfig.Save()
}

func (d *DialogSelectDirectory) sdLoadLastDir() {
	lastDir := openFileMGconfig.GetString("lastdir", "")

	if lastDir != "" {
		d.lastDir = strings.TrimSpace(lastDir)
	}
}
