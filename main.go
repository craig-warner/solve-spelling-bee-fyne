package main

import (
//	"fmt"
//	"log"
	"unicode"	
	"strings"	
//	"regexp"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func IsAllLetters(str string) bool {
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func CheckInput(letters string,required_len int) bool {
	if(len(letters) == required_len) {
		if(IsAllLetters(letters)) {
			return true
		} else { 
			return(false)
		}
	} else {
		return false 
	}
}

func main() {
	var solution []string
	var one_str string
	var popup *widget.PopUp

	myApp := app.New()
	myWindow := myApp.NewWindow("Solve Spelling Bee")
	myWindow.SetPadded(false)

	// Resize ignored by Mobile Platforms
	// - Mobile platforms are always full screen
	// - 27 is a hack determined by Ubuntu/Gnome
	myWindow.Resize(fyne.NewSize(256, (256 + 27)))

	// Prepare Dictionary
	word_list := NewWordList()

	// Output
	start_str:= "Solution will be displayed here\n"
	text := widget.NewLabel(start_str)
	text.TextStyle = fyne.TextStyle{Monospace: true}

	center_letter_entry := widget.NewEntry()
	other_letters_entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Center Letter", Widget: center_letter_entry}},
		OnSubmit: func() { // optional, handle form submission
			center_letter := center_letter_entry.Text
			center_letter = strings.ToLower(center_letter)
			other_letters := other_letters_entry.Text
			other_letters = strings.ToLower(other_letters)
			ok := CheckInput(center_letter,1)
			if(ok) {
//				log.Println("Center Letter is fine")
				ok := CheckInput(other_letters,6)
				if(ok) {
					solution = word_list.Solve(center_letter,other_letters)
					for i:=0;i<len(solution);i++ {
						one_str = one_str + solution[i] + "\n"
					}
//					log.Println(one_str)
					text.SetText(one_str)
				} else { 
					err_msg := widget.NewLabel("Other Letters must be six letter")
					popUpContent := container.NewVBox(
						err_msg,
						widget.NewButton("Ok", func() {
							popup.Hide()
						}),
					)
					popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
					popup.Show()
				}
//				log.Println("Center Letter:", center_letter.Text)
//				log.Println("Other Letters:", other_letters.Text)
			} else {
				err_msg := widget.NewLabel("Center Letter must be one letter")
				popUpContent := container.NewVBox(
					err_msg,
					widget.NewButton("Ok", func() {
						popup.Hide()
					}),
				)
				popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
				popup.Show()
			}
			// Update display
			//myWindow.Close()
		},
	}
	// we can also append items
	form.Append("Other Letters", other_letters_entry)

	// Containing it all
	output_container := container.NewVScroll(text)
	output_container.SetMinSize(fyne.NewSize(200,200))

	top_containter := container.NewVBox(
					form,
					output_container)

	myWindow.SetContent(top_containter)
	myWindow.ShowAndRun()
}
