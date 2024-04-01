package main

import (
	"A-Secure-File-Sharing-System/client"
	"fmt"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var username string
var password string

func main() {
	// 初始化测试用户
	client.InitUser("test", "test")

	// 创建应用程序
	fyneApp := app.NewWithID("test")

	// 创建登录界面
	err := makeLogin()
	if err != nil {
		log.Fatal(err)
	}

	// 运行应用程序
	fyneApp.Run()
}

func tidyUp() {
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func makeUI() (*widget.Label, *widget.Entry) {
	// return widget.NewLabel("Hello world!"),
	// 	widget.NewEntry()

	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

// 这里要写一个shell的界面以方便用户交互
func main_2() {
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
	w.SetMaster()
	w.Show()

	// w2 := a.NewWindow("window2")
	// // w2.SetContent(widget.NewLabel("window2"))
	// w2.SetContent(widget.NewButton("Open new", func() {
	// 	w3 := a.NewWindow("window3")
	// 	w3.SetContent(widget.NewLabel("window3"))
	// 	w3.Show()
	// }))

	// w2.Show()

	// w4 := a.NewWindow("window4")
	// w4.SetContent(container.NewVBox(makeUI()))
	// w4.Show()

	// myWindow := a.NewWindow("Canvas")
	// myCanvas := myWindow.Canvas()

	// blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	// rect := canvas.NewRectangle(blue)
	// myCanvas.SetContent(rect)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	// 	rect.FillColor = green
	// 	rect.Refresh()
	// }()

	// myWindow.Resize(fyne.NewSize(100, 100))
	// myWindow.Show()

	// myWindow2 := a.NewWindow("Widget")

	// myWindow2.SetContent(widget.NewEntry())
	// myWindow2.Show()

	myWindow3 := a.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	text3 := canvas.NewText("Fore", green)
	text3.Move(fyne.NewPos(0, 40))
	content := container.NewWithoutLayout(text1, text2, text3)
	// content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow3.SetContent(content)
	myWindow3.Show()

	myWindow4 := a.NewWindow("List Widget")
	var data = []string{"a", "string", "list"}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow4.SetContent(list)
	myWindow4.Show()

	myWindow5 := a.NewWindow("Table Widget")

	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"a", "b", "c"}
			case "a":
				return []widget.TreeNodeID{"a1", "a2"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "a"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		})

	myWindow5.SetContent(tree)
	myWindow5.Show()

	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
	myInt = 6
	boundInt.Reload()
	i, _ = boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)

	w6 := a.NewWindow("Simple")

	str := binding.NewString()
	str.Set("Initial value")

	text := widget.NewLabelWithData(str)
	w.SetContent(text)

	go func() {
		time.Sleep(time.Second * 5)
		str.Set("A new string")
	}()

	w6.Show()

	a.Run()
	tidyUp()

}
