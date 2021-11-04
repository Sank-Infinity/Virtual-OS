package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/internal/widget"
	// "fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func showGalleryApp(w fyne.Window) {

	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(800,600))
	
	root_src := "C:\\Users\\Sank\\OneDrive\\Desktop\\Go tut\\images"

	files , err := ioutil.ReadDir(root_src)

	tabs := container.NewAppTabs()
	if err != nil {
		fmt.Print("Hello")
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() == false {
			extention := strings.Split(file.Name() , ".")[1]

			if extention == "png" || extention == "jpeg" || extention == "jpg" {
				tabs.Append( container.NewTabItem( file.Name(), canvas.NewImageFromFile( root_src+"\\"+file.Name() ) ) )
			}
		}
		
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs),)

	w.Show()

}