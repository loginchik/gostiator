package customLayouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"math"
)

var _ fyne.Layout = (*AdaptiveGridLayoutRatio)(nil)

func NewRatioLayout(ratios ...float32) fyne.Layout {
	return &AdaptiveGridLayoutRatio{ratios: ratios, adapt: true}
}

type AdaptiveGridLayoutRatio struct {
	ratios          []float32
	adapt, vertical bool
}

func (g *AdaptiveGridLayoutRatio) horizontal() bool {
	if g.adapt {
		return fyne.IsHorizontal(fyne.CurrentDevice().Orientation())
	} else {
		return !g.vertical
	}
}

func (g *AdaptiveGridLayoutRatio) countRows(objects []fyne.CanvasObject) int {
	var totalRatio float32 = 0
	for _, rat := range g.ratios {
		totalRatio += rat
	}
	var count = int(math.Ceil(float64(totalRatio / 1)))
	return count
}

func (g *AdaptiveGridLayoutRatio) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	var rows = g.countRows(objects)
	var columns = len(g.ratios)

	var padWidth = float32(columns-1) * theme.Padding()
	var padHeight = float32(rows-1) * theme.Padding()
	var tGap = float64(padWidth)
	var tCellWidth = float64(size.Width) - tGap
	var cellHeight = float64(size.Height-padHeight) / float64(rows)

	if !g.horizontal() {
		padWidth, padHeight = padHeight, padWidth
		tCellWidth = float64(size.Width-padWidth) - tGap
		cellHeight = float64(size.Height-padHeight) / float64(columns)
	}

	var row, col = 0, 0
	var x1, x2, y1, y2 float32 = 0.0, 0.0, 0.0, 0.0
	for i, child := range objects {
		if !child.Visible() {
			continue
		}

		if i == 0 {
			x1 = 0
			y1 = 0
		} else {
			x1 = x2 + theme.Padding()*1
			y1 = y2 - float32(cellHeight)
		}
		x2 = x1 + float32(tCellWidth)*g.ratios[i]
		y2 = float32(cellHeight)

		var childPos = fyne.NewPos(x1, y1)
		child.Move(childPos)
		var childSize = fyne.NewSize(x2-x1, y2-y1)
		child.Resize(childSize)

		if g.horizontal() {
			if (i+1)%columns == 0 {
				row++
				col = 0
			} else {
				col++
			}
		} else {
			if (i+1)%columns == 0 {
				col++
				row = 0
			} else {
				row++
			}
		}
	}
}

func (g *AdaptiveGridLayoutRatio) MinSize(objects []fyne.CanvasObject) fyne.Size {
	rows := g.countRows(objects)
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		minSize = minSize.Max(child.MinSize())
	}
	if g.horizontal() {
		var width = minSize.Width * float32(len(g.ratios))
		var height = minSize.Height * float32(rows)
		minContentSize := fyne.NewSize(width, height)
		return minContentSize.Add(fyne.NewSize(theme.Padding()*fyne.Max(float32(len(g.ratios)-1), 0), theme.Padding()*fyne.Max(float32(rows-1), 0)))
	}

	minContentSize := fyne.NewSize(minSize.Width*float32(rows), minSize.Height*float32(len(g.ratios)))
	return minContentSize.Add(fyne.NewSize(theme.Padding()*fyne.Max(float32(rows-1), 0), theme.Padding()*fyne.Max(float32(len(g.ratios)-1), 0)))
}
