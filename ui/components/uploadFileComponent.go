package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
)

func UploadFileBlock(parentWindow fyne.Window, citationType string) *fyne.Container {
	var fileEntry = widget.NewEntry()
	fileEntry.Disable()

	var generateButton = widget.NewButton("Цитировать", func() {
		switch citationType {
		case "article":
			log.Println("Для статей это будет чуть проще, но потом")
		default:
			log.Println("Надо бы сгенерировать по файлу")
		}
	})
	generateButton.Hide()

	var uploadFile = widget.NewButton("Загрузить файл", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, parentWindow)
				return
			}
			if reader == nil {
				// means file upload was cancelled
				return
			} else {
				fileEntry.SetText(reader.URI().Path())
				generateButton.Show()
			}

		}, parentWindow)
		fd.Resize(fyne.NewSize(700, 500))
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		fd.Show()
	})
	return container.NewVBox(uploadFile, fileEntry, generateButton)
}
