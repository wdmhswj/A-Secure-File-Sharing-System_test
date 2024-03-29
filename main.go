package main

import (
	"A-Secure-File-Sharing-System/client"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// 这里要写一个shell的界面以方便用户交互
func main() {
	user1, erro := client.InitUser("cyny666", "123456")
	if erro != nil {
		fmt.Print(erro)
	}
	user2, erro2 := client.GetUser("cyny666", "123456")
	if erro2 != nil {
		fmt.Println(erro2)
	}
	fmt.Print(user1)
	fmt.Println(user2)

	// myApp := app.New()
	// myWindow := myApp.NewWindow("Hello")
	// // myWindow.SetContent(widget.NewLabel("Hello"))
	// clock := widget.NewLabel("")
	// formatted := time.Now().Format("Time: 03:04:05")
	// clock.SetText(formatted)
	// myWindow.SetContent(clock)

	// myWindow.Show()
	// myApp.Run()

	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Resize(fyne.NewSize(200, 100))
	w.ShowAndRun()

	tidyUp()

}

func tidyUp() {
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
