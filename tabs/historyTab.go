package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
	"gostCituations/models"
)

// HistoryWindow creates window that contains the list of all
// citations, saved into history file.
func HistoryWindow() *fyne.Container {
	var records = models.GetHistoryRecords()
	var historyList = widget.NewList(
		func() int { return len(records.Records) },
		func() fyne.CanvasObject {
			var label = widget.NewLabel("")
			label.Truncation = fyne.TextTruncateEllipsis
			var button = widget.NewButton("", func() {
				clipboard.Write(clipboard.FmtText, []byte(label.Text))
			})
			button.SetIcon(theme.ContentCopyIcon())
			var block = container.New(layout.NewFormLayout(), button, label)
			return block
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			contBlock, _ := object.(*fyne.Container)
			label := contBlock.Objects[1].(*widget.Label)
			label.SetText(records.Records[id].Content)
			label.Refresh()
		},
	)
	var refreshButton = widget.NewButton("", func() {
		currentlyShowing := len(records.Records)
		records = models.GetHistoryRecords()
		if currentlyShowing != len(records.Records) {
			historyList.Refresh()
		}
	})
	refreshButton.SetIcon(theme.ViewRefreshIcon())

	var historyContent = container.NewBorder(refreshButton, nil, nil, nil, historyList)
	return historyContent
}
