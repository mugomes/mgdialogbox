# MGDialogBox

[![License](https://img.shields.io/badge/license-MIT-5351FB)](LICENSE)

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

* Go 1.27.2+
* Fyne 2.8.0

---

## 👤 Autor

**Murilo Gomes**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes <profmugomes.com.br>

Licensed under the [MIT](https://github.com/profmugomes/mgdialogbox/blob/main/LICENSE) license.

All contributions to the MGDialogBox are subject to this license.
