package main

import (
	"golang.org/x/crypto/argon2"
	"encoding/hex"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/theme"
)

func main() {
	
  	app := app.New()
	app.Settings().SetTheme(theme.DarkTheme())
  	app.SetIcon(resourceIconPng)
  	w := app.NewWindow("Argon2id")

  	passEntry := widget.NewPasswordEntry()
  	passEntry.SetPlaceHolder("with good entropy")
	saltEntry := widget.NewEntry()
	saltEntry.SetPlaceHolder("with 16 random bytes minimum")

  	passwordBox := widget.NewHBox(widget.NewLabel("Input a password:"), layout.NewSpacer(), passEntry)
	saltinput := widget.NewForm(widget.NewFormItem("Input a salt:", saltEntry))
	
  	hashBtn := widget.NewButton("copy key to clipboard", func() {
	key := argon2.IDKey([]byte(passEntry.Text), []byte(saltEntry.Text), 1, 64*1024, 4, 32)
	result := hex.EncodeToString(key)
	w.Clipboard().SetContent(result)
  	})
	
	overwriteBtn := widget.NewButton("overwrite clipboard", func() {
	w.Clipboard().SetContent("blahblub ...")
	})

  	content := widget.NewVBox(passwordBox, saltinput, hashBtn, overwriteBtn)
  	w.SetContent(content)
  	w.ShowAndRun()
}
