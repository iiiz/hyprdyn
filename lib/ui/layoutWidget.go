package hyprdyn_ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/charmbracelet/log"
)

type LayoutWidget struct {
	widget.BaseWidget

	onDismiss func()
	onSubmit  func(l string)

	tabSelectionIndex *int
	dismissTimer      *time.Timer
	layoutList        LayoutList
	layoutBinding     binding.List[*Layout]
	listWidget        *widget.List
	dd                desktop.Driver
}

func NewLayoutWidget(layouts LayoutList, OnSubmit func(l string), OnDismiss func()) *LayoutWidget {
	layoutWidget := &LayoutWidget{
		layoutList: layouts,
		onSubmit:   OnSubmit,
		onDismiss:  OnDismiss,
	}

	for i, l := range layouts {
		if l.Highlight {
			layoutWidget.tabSelectionIndex = &i
		}
	}

	if layoutWidget.tabSelectionIndex == nil {
		i := 0
		layoutWidget.tabSelectionIndex = &i
	}

	if driver, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		layoutWidget.dd = driver
	}

	layoutWidget.layoutBinding = binding.NewList(func(ll1, ll2 *Layout) bool {
		return ll1.Label == ll2.Label
	})

	layoutWidget.layoutBinding.Set(layouts)

	layoutWidget.listWidget = widget.NewListWithData(layoutWidget.layoutBinding, func() fyne.CanvasObject {
		c := container.New(&fixedRowLayout{})

		return c
	}, func(di binding.DataItem, co fyne.CanvasObject) {
		item, _ := di.(binding.Item[*Layout]).Get()
		container := co.(*fyne.Container)

		container.RemoveAll()
		container.Add(item.GetStyledText())
		container.Refresh()
	})

	layoutWidget.ExtendBaseWidget(layoutWidget)

	return layoutWidget
}

func (lw *LayoutWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(lw.listWidget)
}

// INFO: Implements fyne.Focusable
func (lw *LayoutWidget) FocusGained() {
	if lw.dismissTimer != nil {
		lw.dismissTimer.Stop()
		lw.dismissTimer = nil
	}
}

func (lw *LayoutWidget) TypedRune(r rune) {}

func (lw *LayoutWidget) FocusLost() {
	if lw.dismissTimer != nil {
		lw.dismissTimer.Stop()
	}

	lw.dismissTimer = time.AfterFunc(dismissGrace, lw.onDismiss)
}

func (lw *LayoutWidget) AcceptsTab() bool {
	return true
}

func (lw *LayoutWidget) TypedKey(key *fyne.KeyEvent) {
	mods := lw.dd.CurrentKeyModifiers()

	switch key.Name {
	case fyne.KeyTab:
		currentLayoutList, err := lw.layoutBinding.Get()
		if err != nil {
			log.Fatal(err)
		}

		if (mods & fyne.KeyModifierShift) != 0 {
			if *lw.tabSelectionIndex == 0 {
				*lw.tabSelectionIndex = (len(currentLayoutList) - 1)
			} else {
				*lw.tabSelectionIndex--
			}
		} else {
			if *lw.tabSelectionIndex == (len(currentLayoutList) - 1) {
				*lw.tabSelectionIndex = 0
			} else {
				*lw.tabSelectionIndex++
			}
		}

		for i, layout := range currentLayoutList {
			if i == *lw.tabSelectionIndex {
				layout.Highlight = true
			} else {
				layout.Highlight = false
			}
		}

		lw.listWidget.Refresh()
		break
	case fyne.KeyEnter, fyne.KeyReturn:
		currentLayoutList, err := lw.layoutBinding.Get()
		if err != nil {
			log.Fatal(err)
		}

		selectedLayout := currentLayoutList[*lw.tabSelectionIndex]

		if selectedLayout != nil {
			lw.onSubmit(selectedLayout.Label)
		}
		break
	case fyne.KeyEscape:
		lw.onDismiss()
		break
	}
}
