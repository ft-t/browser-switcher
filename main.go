package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"path"
	"regexp"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided")
		return
	}

	homeDir, _ := os.UserHomeDir()
	configPath := path.Join(homeDir, "BrowserSwitcher", "config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err = json.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	targetURL := os.Args[1]
	targetBrowser := findBrowser(targetURL, &config)

	if targetBrowser == nil {
		showUI(&config, targetURL)
		return
	}

	if err = runBrowser(targetBrowser, targetURL); err != nil {
		panic(err)
	}
}

type tappableContainer struct {
	*fyne.Container
	OnTapped func()
}

func (b *tappableContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(b.Container)
}

func (b *tappableContainer) Tapped(ev *fyne.PointEvent) {
	b.OnTapped()
}

func showUI(cfg *Config, targetURL string) {
	myApp := app.NewWithID("com.ft-t.browser-switcher")
	myWindow := myApp.NewWindow("Browser List")

	var elements []fyne.CanvasObject
	for _, browser := range cfg.Browsers {
		name := widget.NewLabel(browser.Name)
		image := canvas.NewImageFromResource(theme.FyneLogo())
		image.ScaleMode = canvas.ImageScaleFastest
		image.FillMode = canvas.ImageFillOriginal
		//image.Resize(fyne.NewSize(128, 128))
		//image.Refresh()

		name.Alignment = fyne.TextAlignCenter

		vbox := container.NewBorder(image, name, nil, nil)
		redBackground := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 0, A: 0})
		redBackground.StrokeWidth = 1
		redBackground.StrokeColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
		withBorder := container.NewStack(vbox, redBackground) // White border around the red background

		//elements = append(elements, &tappableContainer{
		//	Container: vbox,
		//	OnTapped: func() {
		//		if err := runBrowser(browser, targetURL); err != nil {
		//			fmt.Printf("Failed to launch %s: %v\n", browser.Name, err)
		//		}
		//		os.Exit(0)
		//	},
		//})

		elements = append(elements, withBorder)
	}
	
	wrap := container.NewHBox(elements...)
	// Set the content of the window
	//scroll := container.NewScroll(browserList)
	myWindow.SetContent(wrap)

	// Set window size and show
	myWindow.ShowAndRun()
}

func runBrowser(browser *Browser, targetURL string) error {
	return exec.Command(browser.BinaryPath, slices.Concat(
		[]string{targetURL},
		browser.LaunchArgs)...,
	).Start()
}

func findBrowser(targetURL string, config *Config) *Browser {
	for _, browser := range config.Browsers {
		for _, rule := range browser.Rules {
			if matched, _ := regexp.MatchString(rule, targetURL); matched {
				return browser
			}
		}
	}

	return nil
}
