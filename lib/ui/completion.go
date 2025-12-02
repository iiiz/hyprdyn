package hyprdyn_ui

import (
	"fmt"
	"image/color"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"go.deanishe.net/fuzzy"
)

/**
* CompletionItem
**/
type CompletionItem struct {
	Label     string
	Highlight bool
	NewEntry  bool

	// fuzzy sort
	Score *float64
	Match *bool

	// fyne
	text *canvas.Text
}

func (item *CompletionItem) textReset() {
	item.text.Color = color.White
	item.text.TextStyle = fyne.TextStyle{}
	item.text.TextSize = 14
}

func (item *CompletionItem) GetDefaultText() fyne.CanvasObject {
	if item.text == nil {
		item.text = canvas.NewText(item.Label, color.White)
	}

	item.textReset()

	return item.text
}

func (item *CompletionItem) GetStyledText() fyne.CanvasObject {
	if item.text == nil {
		item.text = canvas.NewText(item.Label, color.White)
	} else {
		item.textReset()
	}

	if item.NewEntry {
		specialRegexp := regexp.MustCompile("^special(?:[:]{1}.*)*$")

		item.text.Text = fmt.Sprintf("+> %s", item.Label)
		item.text.TextStyle = fyne.TextStyle{Bold: true}

		if specialRegexp.MatchString(item.Label) {
			item.text.Color = color.RGBA{R: 255, G: 90, B: 90, A: 255}
		} else {
			if item.Highlight {
				item.text.Color = color.RGBA{R: 110, G: 190, B: 255, A: 255}
			} else {
				item.text.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
			}
		}

		return item.text
	}

	if item.Highlight {
		item.text.Text = item.Label
		item.text.TextStyle = fyne.TextStyle{Bold: true}
		item.text.Color = color.RGBA{R: 90, G: 90, B: 255, A: 255}

		return item.text
	}

	return item.text
}

/**
* CompletionList
**/
type CompletionList []*CompletionItem

func (list CompletionList) Len() int {
	return len(list)
}

func (list CompletionList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list CompletionList) Less(i, j int) bool {
	return list[i].Label < list[j].Label
}

func (list CompletionList) Keywords(i int) string {
	return list[i].Label
}

func (list CompletionList) FuzzySort(term string, dropNegativeScored bool) CompletionList {
	if term == "" {
		return list
	}

	sortResult := fuzzy.New(list).Sort(term)

	if dropNegativeScored {
		var nextList CompletionList

		for _, r := range sortResult {
			if r.Score > 0 {
				for _, li := range list {
					if r.SortKey == li.Label {
						li.Highlight = false
						li.Score = &r.Score
						li.Match = &r.Match

						nextList = append(nextList, li)
					}
				}
			}
		}

		return nextList
	}

	return list
}
