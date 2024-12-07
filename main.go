package main

import (
//	"fmt"
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func CheckInput(letters string,required_len int) bool {
	if(len(letters) == required_len) {
		return true
	} else {
		return true
	}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Solve Spelling Bee")
	myWindow.SetPadded(false)

	// Resize ignored by Mobile Platforms
	// - Mobile platforms are always full screen
	// - 27 is a hack determined by Ubuntu/Gnome
	myWindow.Resize(fyne.NewSize(256, (256 + 27)))

	// Prepare Dictionary
	word_list := NewWordList()

	center_letter := widget.NewEntry()
	other_letters := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Center Letter", Widget: center_letter}},
		OnSubmit: func() { // optional, handle form submission
			ok := CheckInput(center_letter.Text,1)
			if(ok) {
				log.Println("Center Letter is fine")
				ok := CheckInput(other_letters.Text,6)
				if(ok) {
					word_list.Solve(center_letter.Text,other_letters.Text)
				}
				log.Println("Center Letter:", center_letter.Text)
				log.Println("Other Letters:", other_letters.Text)
			}
			myWindow.Close()
		},
	}
	// we can also append items
	form.Append("Other Letters", other_letters)

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
