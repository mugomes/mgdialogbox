// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

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
	"github.com/mugomes/mgsmartflow"
	"github.com/mugomes/mgsettings/v3"
)

type FileDialogOpen struct {
	a           fyne.App
	title       string
	exts        []string
	multiSelect bool
	onSelect    func([]string)

	lastDir string
}

var openFileMGconfig, _ = mgsettings.Load("/tmp/mgdialogbox", false)

func NewOpenFile(a fyne.App, title string, exts []string, multiselect bool, onSelect func([]string)) {
	dlg := &FileDialogOpen{
		a:           a,
		title:       title,
		exts:        exts,
		multiSelect: multiselect,
		onSelect:    onSelect,
	}
	dlg.loadLastDir()

	dlg.showOpenFile()
}

func (d *FileDialogOpen) showOpenFile() {
	win := d.a.NewWindow(d.title)
	win.Resize(fyne.NewSize(740, 520))
	win.CenterOnScreen()

	flow := mgsmartflow.New()

	dir := d.lastDir
	if dir == "" {
		dir, _ = os.UserHomeDir()
	}

	pathLabel := widget.NewLabel(dir)
	search := widget.NewEntry()
	search.SetPlaceHolder("🔍 Search...")

	files := d.listDir(dir)
	filtered := files
	selected := map[int]bool{}

	list := widget.NewList(
		func() int { return len(filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, o fyne.CanvasObject) {
			row := filtered[id]
			name := row.Name()

			if selected[id] {
				o.(*widget.Label).SetText("✔ " + name)
			} else {
				if row.IsDir() {
					o.(*widget.Label).SetText("📁 " + name)
				} else {
					o.(*widget.Label).SetText("📄 " + name)
				}
			}
		},
	)

	var (
		click int = 1
		timeClick time.Time
		durationClick = 500 * time.Millisecond
	)
	
	list.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(filtered) {
			return
		}
		f := filtered[id]

		// Abrir diretório
		if f.IsDir() {
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
				files = d.listDir(dir)
				filtered = d.applyFilter(files, search.Text)
				selected = map[int]bool{}
				list.Refresh()
				click = 1 // reset clique duplo
			}
			
			return
		}

		// Abrir arquivo (somente single-select)
		if !d.multiSelect {
			d.saveLastDir(dir)
			d.onSelect([]string{filepath.Join(dir, f.Name())})
			win.Close()
			return
		}

		// Clique simples
		if f.IsDir() {
			// apenas destaca a pasta
			selected = map[int]bool{id: true}
		} else {
			if d.multiSelect {
				selected[id] = !selected[id]
			} else {
				selected = map[int]bool{id: true}
			}
		}

		list.Refresh()
	}

	// BOTÃO VOLTAR
	btnBack := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		parent := filepath.Dir(dir)
		if parent != dir {
			dir = parent
			pathLabel.SetText(dir)
			files = d.listDir(dir)
			filtered = d.applyFilter(files, search.Text)
			selected = map[int]bool{}
			list.Refresh()
		}
	})

	// BOTÃO ABRIR
	btnOpen := widget.NewButtonWithIcon("Open", theme.ConfirmIcon(), func() {
		var out []string
		for i := range selected {
			f := filtered[i]
			if !f.IsDir() && selected[i] {
				out = append(out, filepath.Join(dir, f.Name()))
			}
		}

		if len(out) > 0 {
			d.saveLastDir(dir)
			d.onSelect(out)
			win.Close()
		}
	})

	// BUSCA
	search.OnChanged = func(txt string) {
		filtered = d.applyFilter(files, txt)
		selected = map[int]bool{}
		list.Refresh()
	}

	// LAYOUT
	flow.AddColumn(btnBack, container.NewVBox(pathLabel, search))
	flow.SetResize(btnBack, fyne.NewSize(68, 79))

	flow.AddRow(list)
	flow.AddRow(btnOpen)

	controls.OnResize(win, func(size fyne.Size) {
		flow.SetResize(list, fyne.NewSize(size.Width, size.Height-137))
	})

	win.Canvas().Overlays().Add(flow.Container)

	win.Show()
}

//////////////////////////////////////////////////////////////
// FUNÇÕES AUXILIARES
//////////////////////////////////////////////////////////////

func (d *FileDialogOpen) listDir(path string) []fs.FileInfo {
	entries, _ := os.ReadDir(path)

	var list []fs.FileInfo

	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}

		if len(d.exts) > 0 && !info.IsDir() {
			ok := false
			for _, ext := range d.exts {
				if strings.EqualFold(filepath.Ext(info.Name()), ext) {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		list = append(list, info)
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

func (d *FileDialogOpen) applyFilter(files []fs.FileInfo, query string) []fs.FileInfo {
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

func (d *FileDialogOpen) saveLastDir(dir string) {
	openFileMGconfig.SetString("lastdir", dir)
	openFileMGconfig.Save()
}

func (d *FileDialogOpen) loadLastDir() {
	lastDir := openFileMGconfig.GetString("lastdir", "")

	if lastDir != "" {
		d.lastDir = strings.TrimSpace(lastDir)
	}
}
