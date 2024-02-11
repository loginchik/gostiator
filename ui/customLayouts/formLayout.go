package customLayouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* form block */

type FormBlock struct {
	Title    *widget.Label
	Elements *fyne.Container
}

func NewFormBlock(titleText string, object *fyne.Container) *fyne.Container {
	var titleLabel = widget.NewLabel(titleText)
	return container.NewVBox(titleLabel, object)
}

/* form layout */

var _ fyne.Layout = (*CustomFormLayout)(nil)

func NewFormLayout() fyne.Layout {
	return &CustomFormLayout{innerPadding: theme.Padding() * 2}
}

type CustomFormLayout struct {
	innerPadding float32
}

func (f *CustomFormLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	var x1, y1 float32 = 0.0, 0.0
	var objectsWidth = size.Width
	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		var childHeight = child.MinSize().Height
		child.Move(fyne.NewPos(x1, y1))
		child.Resize(fyne.NewSize(objectsWidth, childHeight))
		y1 += childHeight + f.innerPadding
	}
}

func (f *CustomFormLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var totalHeight float32 = 0
	var maxWidth float32 = 0
	for _, child := range objects {
		var childSize = child.MinSize()
		totalHeight += childSize.Height
		maxWidth = fyne.Max(childSize.Width, maxWidth)
	}
	totalHeight += f.innerPadding * float32(len(objects)-1)

	return fyne.NewSize(maxWidth, totalHeight)
}
