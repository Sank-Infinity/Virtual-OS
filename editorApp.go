package main

import (
	"io/ioutil"
	"strconv"

	// "fyne.io/fyne/dialog"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	// "fyne.io/fyne/v2/internal/widget"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor( w fyne.Window ) {
	// a := app.New()
	// wt := myApp.NewWindow("Text Editor")
	// wt.Resize(fyne.NewSize(800, 600))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
		),
	)	

	content.Add( widget.NewButton("Add New File" , func(){
		content.Add(widget.NewLabel("New File "+ strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text ...")

	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton( "Save text File" , func() {
		saveFileDialog := dialog.NewFileSave(
			func (uc fyne.URIWriteCloser, _ error)  {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName( "New File "+ strconv.Itoa(count - 1) + ".txt" )
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, e error) {
				readData , _ := ioutil.ReadAll(r)
				
				output := fyne.NewStaticResource("New File", readData)
				
				viewData := widget.NewMultiLineEntry()

				// viewData.Resize(fyne.NewSize(400 , 400))

				viewData.SetText(string(output.StaticContent))
				

				wp:=fyne.CurrentApp().NewWindow(
					string(output.StaticName))

					saveSubBtn := widget.NewButton( "Save text File" , func() {
						saveFileDialog := dialog.NewFileSave(
							func (uc fyne.URIWriteCloser, _ error)  {
								textData := []byte(viewData.Text)
								uc.Write(textData)
							}, wp)
				
						saveFileDialog.SetFileName( "New File "+ strconv.Itoa(count - 1) + ".txt" )
						saveFileDialog.Show()
					})
					
					wp.Resize(fyne.NewSize(800, 600))
					// saveSubBtn.Resize(fyne.NewSize(50, 50))

					wp.SetContent(
						container.NewGridWithRows(2,
							container.NewScroll(viewData),
							container.NewGridWithRows(4,
								container.NewGridWithColumns(4,
									saveSubBtn,),),
						),
					)
					wp.Show()
			}, w)

			openFileDialog.SetFilter(storage.NewExtensionFileFilter( []string{".txt"} ))
			openFileDialog.Show()

	})

	editorContainer := container.NewVBox(
		content,
		input,
		container.NewHBox(saveBtn, openBtn) ,
	)

	w.SetContent( container.NewBorder(DeskBtn, nil, nil, nil, editorContainer), )

	w.Show()
}