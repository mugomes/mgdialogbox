# MGDialogBox

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
go get github.com/mugomes/mgdialogbox
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

## 🖥️ Requisitos

* Go 1.25.5+
* Fyne 2.7.1+

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://mugomes.github.io](https://mugomes.github.io)

📺 [https://youtube.com/@mugomesoficial](https://youtube.com/@mugomesoficial)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgdialogbox/blob/main/LICENSE) license.

All contributions to the MGDialogBox are subject to this license.