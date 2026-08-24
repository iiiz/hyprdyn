package hyprdyn_ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"go.deanishe.net/fuzzy"

	hd "hyprdyn/lib"
)

/**
* CompletionItem
**/
type CompletionItem struct {
	Label     string
	Highlight bool
	NewEntry  bool
	AutoEntry bool

	// fuzzy sort
	Score *float64
	Match *bool

	// fyne
	text *canvas.Text
}

func (item *CompletionItem) textReset() {
	theme := UseTheme()
	item.text.Color = theme.Text
	item.text.TextStyle = fyne.TextStyle{}
	item.text.TextSize = 14
}

func (item *CompletionItem) GetStyledText() fyne.CanvasObject {
	theme := UseTheme()
	if item.text == nil {
		item.text = canvas.NewText(item.Label, theme.Text)
	} else {
		item.textReset()
	}

	if item.NewEntry {
		item.text.Text = fmt.Sprintf("+> %s", item.Label)
		item.text.TextStyle = fyne.TextStyle{Bold: true}

		if !hd.ValidWorkspaceName(item.Label) {
			item.text.Color = theme.DisabledText
		} else {
			if item.Highlight {
				item.text.Color = theme.NewHighLight
			} else {
				item.text.Color = theme.NewText
			}
		}

		return item.text
	}

	if item.AutoEntry {
		item.text.Color = theme.Suggestion
	}

	if item.Highlight {
		item.text.Text = item.Label
		item.text.TextStyle = fyne.TextStyle{Bold: true}
		item.text.Color = theme.Highlight

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
