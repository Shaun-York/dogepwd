package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"

	"dogepwd/characters"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var charset = characters.Charset{}

func mkpwd(length int) string {
	rand.Seed(time.Now().UnixNano())

	if charset.NotSet() {
		charset.All()
	}
	
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
    	buf[i] = charset.Chars[rand.Intn(len(charset.Chars))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
    	buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}

func main() {
	goapp := app.New()
	w := goapp.NewWindow("dogepwd")

	somelabel := widget.NewLabel("...")
	someentry := widget.NewEntry()
	someentry.PlaceHolder = "length"

	w.SetContent(widget.NewVBox(
		somelabel,
		someentry,
		widget.NewButton("Lowercase", func() { 
			charset.Atozl() 
		}),
		widget.NewButton("Uppercase", func() { charset.Atozu() }),
		widget.NewButton("Digits", func() { charset.Digits() }),
		widget.NewButton("Special", func() { charset.Special() }),
		widget.NewButton("Generate Password", func() {
			n := someentry.Text
			if n == "" {
				n = "8"
			}
			i, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				dialog.NewCustom("ERROR", "OK", canvas.NewText(err.Error(), color.White), w).Show()
			}
			pwd := mkpwd(int(i))
			cb := w.Clipboard()
			go cb.SetContent(pwd)
			go somelabel.SetText(pwd)
		}),
		widget.NewButton("Clear", func() { charset.Clear() }),
	))
	w.ShowAndRun()
}