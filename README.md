# MGDialogBox

[![License](https://img.shields.io/badge/license-PolyForm%20Perimeter%201.0.1-5351FB)](LICENSE.md)

MGDialogBox é uma biblioteca em Go baseada no **Fyne** que fornece **caixas de diálogo prontas e reutilizáveis** para aplicações desktop, como alertas, confirmações e seleção de arquivos ou diretórios.

O objetivo é simplificar a criação de diálogos comuns, mantendo uma API limpa, consistente e fácil de integrar.

---

## ✨ Recursos

* Alertas simples (informação/erro)
* Caixa de confirmação com múltiplos botões
* Abertura de arquivos (com filtro de extensões)
* Salvamento de arquivos
* Seleção de diretórios
* Suporte a seleção múltipla
* Callbacks para retorno de ações do usuário

---

## 📦 Instalação

```bash
go get github.com/profmugomes/mgdialogbox
```

---

## 🚀 Uso

### 🔔 Alerta

```go
mgdialogbox.NewAlert(
	app,
	"Aviso",
	"Operação concluída com sucesso",
	false,
	"OK",
)

mgdialogbox.NewAlert(
	app,
	"Aviso",
	"Operação concluída com sucesso",
	false,
	"OK",
	func() {
		// Ação que deseja realizar após o fechamento do alerta
	},
)
```

---

### ❓ Confirmação com múltiplos botões

```go
mgdialogbox.NewConfirm(
	app,
	"Confirmação",
	"Deseja continuar?",
	[]string{"Sim", "Não", "Cancelar"},
	func(result int) {
		// result começa em 0 (Sim = 0, Não = 1, Cancelar = 2)
	},
)
```

---

### 📂 Abrir arquivo

```go
mgdialogbox.NewOpenFile(
	app,
	"Abrir Arquivo",
	[]string{"png", "jpg", "pdf"},
	false,
	func(files []string) {
		// arquivos selecionados
	},
)
```

---

### 💾 Salvar arquivo

```go
mgdialogbox.NewSaveFile(
	app,
	"Salvar Arquivo",
	[]string{"txt", "md"},
	func(path string) {
		// caminho escolhido
	},
)
```

---

### 📁 Selecionar diretório

```go
mgdialogbox.NewSelectDirectory(
	app,
	"Selecionar Pasta",
	true,
	func(paths []string) {
		// diretórios selecionados
	},
)
```

---

## 🧱 Estrutura

O pacote principal atua como **facade**, delegando a lógica de UI para o pacote interno `components`, facilitando manutenção e extensões futuras.

---

## 🧩 Compatibilidade

* Go 1.26.5+
* Fyne 2.8.0

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved.

This project is licensed under the PolyForm Perimeter License 1.0.1.

### Summary

This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.