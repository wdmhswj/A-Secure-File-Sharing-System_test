package main

import (
	"A-Secure-File-Sharing-System/client"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_demo/tutorials"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

// 创建菜单
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	checkedItem := fyne.NewMenuItem("Checked", nil)
	checkedItem.Checked = true
	disabledItem := fyne.NewMenuItem("Disabled", nil)
	disabledItem.Disabled = true
	otherItem := fyne.NewMenuItem("Other", nil)
	mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	mailItem.Icon = theme.MailComposeIcon()
	otherItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
		mailItem,
	)
	fileItem := fyne.NewMenuItem("File", func() { fmt.Println("Menu New->File") })
	fileItem.Icon = theme.FileIcon()
	dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		fileItem,
		dirItem,
		otherItem,
	)

	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(440, 520))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut
	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }
	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", newItem, checkedItem, disabledItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)
	checkedItem.Action = func() {
		checkedItem.Checked = !checkedItem.Checked
		main.Refresh()
	}
	return main
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

// 登陆界面
func makeLogin(app fyne.App) error {
	// 登陆窗口
	loginWidget := app.NewWindow("LogIn")

	username := widget.NewEntry()
	username.SetPlaceHolder("John Smith")

	// email := widget.NewEntry()
	// email.SetPlaceHolder("test@example.com")
	// email.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: username, HintText: "Your username"},
			{Text: "Email", Widget: password, HintText: "Your passwrod"},
		},
		OnCancel: func() {
			loginWidget.Close()
			log.Println("quit")
		},
		OnSubmit: func() {
			if authenticate(username.Text, password.Text) {

				User, _ := client.GetUser(username.Text, password.Text)
				showMainWindow(app, User)
				loginWidget.Close()
				// // 发送显示主界面的操作到主 goroutine 中
				// app.Send(func() {
				// 	showMainWindow(fyneApp)
				// })
			} else {
				dialog.ShowError(errors.New("invalid username or password"), loginWidget)
			}
		},
	}

	loginWidget.SetContent(form)
	// loginWidget.SetContent(content)
	loginWidget.Resize(fyne.NewSize(340, 460))
	loginWidget.Show()
	return nil

}

// func makeLogin_v2(app fyne.App) error {
// 	// 登陆窗口
// 	loginWidget := app.NewWindow("LogIn")
// 	// 创建登录界面的代码
// 	// 包括用户名、密码输入框、登录按钮等
// 	username := widget.NewEntry()
// 	username.SetPlaceHolder("John Smith")
// 	password := widget.NewPasswordEntry()
// 	password.SetPlaceHolder("Password")
// 	loginButton := widget.NewButton("click me", func() {
// 		if authenticate(username.Text, password.Text) {
// 			loginWidget.Close()
// 			showMainWindow(app,)
// 		} else {
// 			dialog.ShowError(errors.New("invalid username or password"), loginWidget)
// 		}
// 	})

// 	content := container.New(layout.NewVBoxLayout(), username, password, layout.NewSpacer(), loginButton)
// 	loginWidget.SetContent(container.New(layout.NewVBoxLayout(), content))
// 	// loginWidget.SetContent(content)
// 	loginWidget.Resize(fyne.NewSize(340, 460))
// 	loginWidget.Show()
// 	return nil
// }

func makeLogin_v1(win fyne.Window) (<-chan string, <-chan string, *dialog.FormDialog, error) {
	username := widget.NewEntry()
	username.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "username can only contain letters, numbers, '_', and '-'")
	password := widget.NewPasswordEntry()
	password.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "password can only contain letters, numbers, '_', and '-'")
	remember := false
	items := []*widget.FormItem{
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
		widget.NewFormItem("Remember me", widget.NewCheck("", func(checked bool) {
			remember = checked
		})),
	}

	usernameChan := make(chan string)
	passwordChan := make(chan string)
	newForm := dialog.NewForm("Login...", "Log In", "Cancel", items, func(b bool) {
		if !b {
			// return "", "", errors.New("Verification failed.")
			log.Println("Verification failed.")
		}
		var rememberText string
		if remember {
			rememberText = "and remember this login"
		}

		usernameChan <- username.Text
		passwordChan <- password.Text
		log.Println("Please Authenticate", username.Text, password.Text, rememberText)

	}, win)

	return usernameChan, passwordChan, newForm, nil
}

// 注册界面
func makeRegister() {

}

func setUsernamePassword(un string, ps string) {
	username = un
	password = ps
	log.Printf("username: %s, password: %s\n", username, password)
}

func authenticate(username, password string) bool {
	// 用户认证的逻辑
	// 这里可以是用户名密码验证、API 调用等
	_, err := client.GetUser(username, password)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func showMainWindow(app fyne.App, User *client.User) {

	log.Println("登陆成功")

	w := app.NewWindow("test")

	w.Resize(fyne.NewSize(640, 460)) // 重置窗口大小

	// newNav := makeNav()
	// w.SetContent(newNav)

	newTabs := makeTabs(w, User)
	w.SetContent(newTabs)

	w.Show()
	// log.Println("w.Show()")

}

func unsupportedTutorial(t tutorials.Tutorial) bool {
	return !t.SupportWeb && fyne.CurrentDevice().IsBrowser()
}

// 导航栏
func makeNav() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))

	return content
}

func makeNav_v1(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string { // 树节点的子节点
			return tutorials.TutorialIndex[uid]
		},
		IsBranch: func(uid string) bool { // 是否是分支节点，即是否有子节点
			children, ok := tutorials.TutorialIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject { // 定义了如何创建树节点
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) { // 更新树节点的内容
			t, ok := tutorials.Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			if unsupportedTutorial(t) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			} else {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			}
		},
		OnSelected: func(uid string) {
			if t, ok := tutorials.Tutorials[uid]; ok {
				if unsupportedTutorial(t) {
					return
				}
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious { // 尝试加载之前用户选择的教程，并将其选中
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

func makeTabs(win fyne.Window, User *client.User) fyne.CanvasObject {

	// StoreFileButtion := makeDialogOpenFileButton(win)
	StoreFile := makeStoreFile(win, User)
	LoadFile := makeLoadFile(win, User)
	tabs := container.NewAppTabs(
		container.NewTabItem("StoreFile", StoreFile),
		container.NewTabItem("LoadFile", LoadFile),
		container.NewTabItem("AppendToFile", widget.NewLabel("AppendToFile")),
		container.NewTabItem("CreateInvitation", widget.NewLabel("CreateInvitation")),
		container.NewTabItem("AcceptInvitation", widget.NewLabel("AcceptInvitation")),
		container.NewTabItem("RevokeAccess", widget.NewLabel("RevokeAccess")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}

func makeDialogOpenFileButton(win fyne.Window, filename *string, data *[]byte) *widget.Button {
	openFile := widget.NewButton("File Open Without Filter", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			// imageOpened(reader)
			*filename = reader.URI().Name()
			*data, err = io.ReadAll(reader)
			if err != nil {
				fyne.LogError("Failed to load file data", err)
				return
			}
			log.Printf("the filename: %s, the file content: %s", *filename, string(*data))
		}, win)
		// fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fd.Show()
	})

	return openFile
}

func makeDialogFileSaveButtion(win fyne.Window, filename *string, User *client.User) *widget.Button {
	saveFile := widget.NewButton("File Save", func() {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if writer == nil {
				log.Println("Cancelled")
				return
			}

			fileSaved(writer, win, filename, User)
		}, win)
	})

	return saveFile
}

func fileSaved(f fyne.URIWriteCloser, w fyne.Window, filename *string, User *client.User) {
	defer f.Close()

	// loadFile
	if len(*filename) == 0 {
		dialog.ShowError(errors.New("filename is empty"), w)
		return
	}
	log.Printf("filename: %s", *filename)
	data, err := User.LoadFile(*filename)
	if err != nil {
		log.Println(err.Error())
		dialog.ShowError(err, w)
		return
	}
	log.Printf("data: %s", data)
	_, err = f.Write(data)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	
	log.Println("Saved to...", f.URI())
}

func makeStoreFile(win fyne.Window, User *client.User) fyne.CanvasObject {
	// layout
	// 打开文件按钮
	// 保存上传按钮

	var filename string
	var data []byte

	StoreFileButtion := makeDialogOpenFileButton(win, &filename, &data)

	StoreAndUploadButton := widget.NewButton("upload", func() {
		if len(filename) == 0 || len(data) == 0 {
			log.Println("File does not exist.")
			dialog.ShowInformation("Error", "File does not exist!", win)
			// dialog.NewError(errors.New("file does not exist"), win)
			return
		}
		err := User.StoreFile(filename, data)
		if err != nil {
			fyne.LogError("failed to storefile", err)
			return
		}
		dialog.ShowInformation("Success", "File uploaded successfully.", win)
		filename = ""
		data = nil
		log.Println("test ShowInformation")
	})

	content := container.New(layout.NewVBoxLayout(), StoreFileButtion, layout.NewSpacer(), StoreAndUploadButton)
	return content
}

func makeLoadFile(win fyne.Window, User *client.User) fyne.CanvasObject {
	// 输入文件名的输入框
	filename := widget.NewEntry()
	filename.SetPlaceHolder("Enter the filename")

	// 保存按钮
	LoadAndSaveButton := makeDialogFileSaveButtion(win, &filename.Text, User)

	// 从datastore获取数据
	// 保存到本地文件夹

	content := container.New(layout.NewVBoxLayout(), filename, layout.NewSpacer(), LoadAndSaveButton)

	return content

}
