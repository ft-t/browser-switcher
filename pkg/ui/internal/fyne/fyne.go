package fyne

//
//type tappableContainer struct {
//	*fyne.Container
//	OnTapped func()
//}
//
//func (b *tappableContainer) CreateRenderer() fyne.WidgetRenderer {
//	return widget.NewSimpleRenderer(b.Container)
//}
//
//func (b *tappableContainer) Tapped(ev *fyne.PointEvent) {
//	b.OnTapped()
//}
//
//func showUI(cfg *config2.Config, targetURL string) {
//	myApp := app.NewWithID("com.ft-t.browser-switcher")
//	myWindow := myApp.NewWindow("Browser List")
//
//	var elements []fyne.CanvasObject
//	for _, browser := range cfg.Browsers {
//		name := widget.NewLabel(browser.Name)
//		image := canvas.NewImageFromResource(theme.FyneLogo())
//		image.ScaleMode = canvas.ImageScaleFastest
//		image.FillMode = canvas.ImageFillOriginal
//		//image.Resize(fyne.NewSize(128, 128))
//		//image.Refresh()
//
//		name.Alignment = fyne.TextAlignCenter
//
//		vbox := container.NewBorder(image, name, nil, nil)
//		redBackground := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 0, A: 0})
//		redBackground.StrokeWidth = 1
//		redBackground.StrokeColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
//		withBorder := container.NewStack(vbox, redBackground) // White border around the red background
//
//		//elements = append(elements, &tappableContainer{
//		//	Container: vbox,
//		//	OnTapped: func() {
//		//		if err := runBrowser(browser, targetURL); err != nil {
//		//			fmt.Printf("Failed to launch %s: %v\n", browser.Name, err)
//		//		}
//		//		os.Exit(0)
//		//	},
//		//})
//
//		elements = append(elements, withBorder)
//	}
//
//	wrap := container.NewHBox(elements...)
//	// Set the content of the window
//	//scroll := container.NewScroll(browserList)
//	myWindow.SetContent(wrap)
//
//	// Set window size and show
//	myWindow.ShowAndRun()
//}
