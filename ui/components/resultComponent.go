package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kr/text"
	"golang.design/x/clipboard"
	"strings"
	"time"
)

func ResultWindow(results []string, application fyne.App) fyne.Window {
	var wind = application.NewWindow(fmt.Sprintf("Цитата - %s", time.Time.Format(time.Now(), "15:04")))
	var resultsWrapped []string
	for _, res := range results {
		resultsWrapped = append(resultsWrapped, text.Wrap(res, 120))
	}
	var resultWrapped = strings.Join(resultsWrapped, "\n")
	var resultLabel = widget.NewLabel(resultWrapped)

	var copyButton = widget.NewButton("Копировать", func() {
		clipboard.Write(clipboard.FmtText, []byte(strings.Join(results, "\n")))
	})

	var windowContent = container.NewVBox(resultLabel, copyButton)
	wind.SetContent(windowContent)
	var s = fyne.NewSize(200, windowContent.Size().Height)
	wind.Resize(s)
	return wind
}
