package hyprdyn_ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/charmbracelet/log"
)

type SelectorWidget struct {
	widget.Entry

	TabCompleteListWidget *widget.List

	onSubmit          func(i string)
	onDismiss         func()
	completionBinding binding.List[*CompletionItem]
	completionList    CompletionList
	tabSelectionIndex *int
	dd                desktop.Driver
}

func NewSelectorWidget(workspaceNames []string, OnSubmit func(i string), OnResize func(y float32), OnDismiss func()) (*SelectorWidget, float32) {
	selector := &SelectorWidget{}
	selector.onSubmit = OnSubmit
	selector.onDismiss = OnDismiss
	selector.SetPlaceHolder("Workspace name")
	selector.tabSelectionIndex = nil

	if driver, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		selector.dd = driver
	}

	selector.completionBinding = binding.NewList(func(cli1, cli2 *CompletionItem) bool {
		return cli1.Label == cli2.Label
	})

	for _, name := range workspaceNames {
		selector.completionList = append(selector.completionList, &CompletionItem{Label: name, Highlight: false})
	}

	selector.completionBinding.Set(selector.completionList)

	selector.TabCompleteListWidget = widget.NewListWithData(selector.completionBinding, func() fyne.CanvasObject {
		c := container.New(&fixedRowLayout{})

		return c
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		item, _ := di.(binding.Item[*CompletionItem]).Get()
		container := co.(*fyne.Container)

		container.RemoveAll()
		container.Add(item.GetStyledText())
		container.Refresh()
	})

	selector.OnChanged = func(input string) {
		if selector.tabSelectionIndex != nil {
			selector.tabSelectionIndex = nil

			for _, li := range selector.completionList {
				li.Highlight = false
			}
		}

		nextList := selector.completionList.FuzzySort(input, true)
		nextLen := nextList.Len()

		if (nextLen == 1 && *nextList[0].Match == false) || nextLen == 0 {
			nextList = append(nextList, &CompletionItem{
				Label: input, Highlight: false, NewEntry: true,
			})

			if nextLen == 1 {
				nextList.Swap(0, 1)
			}
		}

		selector.completionBinding.Set(nextList)
		selector.TabCompleteListWidget.Refresh()

		// Would love to have a proper list item height here but I can't find a const or a layout that's suitable in this framework for the use case.
		OnResize((50 * float32(nextList.Len())) + selector.Entry.Size().Height)
	}

	selector.ExtendBaseWidget(selector)

	return selector, (48*float32(selector.completionList.Len()) + 42)
}

func (s *SelectorWidget) FocusLost() {
	s.onDismiss()
}

func (s *SelectorWidget) AcceptsTab() bool {
	return true
}

func (s *SelectorWidget) TypedKey(key *fyne.KeyEvent) {
	mods := s.dd.CurrentKeyModifiers()

	switch key.Name {
	case fyne.KeyTab:
		currentCompletionList, err := s.completionBinding.Get()
		if err != nil {
			log.Fatal(err)
		}

		if s.tabSelectionIndex == nil {
			var i int = 0
			s.tabSelectionIndex = &i
		} else {
			if (mods & fyne.KeyModifierShift) != 0 {
				if *s.tabSelectionIndex == 0 {
					*s.tabSelectionIndex = (len(currentCompletionList) - 1)
				} else {
					*s.tabSelectionIndex--
				}
			} else {
				if *s.tabSelectionIndex == (len(currentCompletionList) - 1) {
					*s.tabSelectionIndex = 0
				} else {
					*s.tabSelectionIndex++
				}
			}
		}

		for i, ci := range currentCompletionList {
			if i == *s.tabSelectionIndex {
				ci.Highlight = true
			} else {
				ci.Highlight = false
			}
		}

		s.TabCompleteListWidget.Refresh()
	case fyne.KeyReturn, fyne.KeyEnter:
		var input string

		if s.tabSelectionIndex != nil {
			currentCompletionList, err := s.completionBinding.Get()
			if err != nil {
				log.Fatal(err)
			}

			input = currentCompletionList[*s.tabSelectionIndex].Label
		} else {
			input = s.Entry.Text

			if input == "" {
				break
			}
		}

		s.onSubmit(input)
	default:
		s.Entry.TypedKey(key)
	}
}

func (s *SelectorWidget) KeyDown(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyEscape {
		s.onDismiss()
	}
}
