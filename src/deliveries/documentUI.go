package deliveries

import (
	"constants"
	"entitites"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (c *ConsoleUI) addDocumentPage() (flex *tview.Flex) {

	var documentNo, documentTitle, notes string
	f := tview.NewForm().
		AddInputField(c.requiredLabel(constants.DocNoLabel[c.defLang]), "", 30, nil, func(text string) {
			documentNo = text
		}).
		AddInputField(c.requiredLabel(constants.DocTitleLabel[c.defLang]), "", 30, nil, func(text string) {
			documentTitle = text
		}).
		AddInputField(constants.NotesLabel[c.defLang], "", 50, nil, func(text string) {
			notes = text
		}).
		AddButton(constants.SaveLabel[c.defLang], func() {
			a := entitites.Document{
				DocumentNo:    documentNo,
				DocumentTitle: documentTitle,
				Notes:         notes,
			}
			err := a.Validate()
			if err != nil {
				c.changePage(constants.ErrorNotifPageId, c.failedNotification())
			} else {
				c.docUseCase.CreateDocument(a)
				c.changePage(constants.SuccessNotifPageId, c.successNotification())
			}

		}).
		AddButton(constants.CancelLabel[c.defLang], func() {
			c.changePage(constants.MainPageId, c.mainPage())
		})
	f.SetBorder(true).SetTitle(constants.AddDocLabel[c.defLang]).SetTitleAlign(tview.AlignLeft)

	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f, 0, 4, true)
	return
}

func (c *ConsoleUI) listDocumentPage(list []*entitites.Document) (flex *tview.Flex) {
	t := tview.NewTable().SetBorders(true)
	t.SetBorder(true).SetTitle("Daftar Arsip")
	t.SetCell(0, 0, tview.NewTableCell(constants.DocNoLabel[c.defLang]))
	t.SetCell(0, 1, tview.NewTableCell(constants.DocTitleLabel[c.defLang]))
	t.SetCell(0, 2, tview.NewTableCell(constants.NotesLabel[c.defLang]))

	for r := 1; r <= len(list); r++ {
		t.SetCell(r, 0,
			tview.NewTableCell(list[r-1].DocumentNo))
		t.SetCell(r, 1,
			tview.NewTableCell(list[r-1].DocumentTitle))
		t.SetCell(r, 2,
			tview.NewTableCell(list[r-1].Notes))
	}
	nav := c.navigation("")

	var searchField *tview.InputField
	searchField = tview.NewInputField().
		SetLabel("Enter a number: ").
		SetPlaceholder("E.g. 1234").
		SetFieldWidth(20).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				list := c.docUseCase.FindDocumentByNumber(searchField.GetText())
				c.changePage(constants.ViewListDocPageId, c.listDocumentPage(list))
			}
			if key == tcell.KeyTAB {
				c.app.SetFocus(nav)
			}
		})
	searchField.SetBorder(true).SetTitle("Cari Arsip").SetTitleAlign(tview.AlignLeft)
	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(searchField, 0, 1, true).
		AddItem(t, 0, 4, false).
		AddItem(nav, 0, 1, true)

	return
}

func (c *ConsoleUI) successNotification() (modal *tview.Modal) {
	return c.notificationModal(constants.SuccessLabel[c.defLang], constants.SuccessMessageType, []string{constants.AddMoreLabel[c.defLang], constants.FinishLabel[c.defLang]}, func(buttonIndex int, buttonLabel string) {
		if buttonLabel == constants.FinishLabel[c.defLang] {
			c.changePage(constants.MainPageId, c.mainPage())
		} else {
			c.pages.RemovePage(constants.SuccessNotifPageId)
			c.changePage(constants.AddDocPageId, c.addDocumentPage())
		}
	})
}
func (c *ConsoleUI) failedNotification() (modal *tview.Modal) {
	return c.notificationModal(constants.MandatoryFailedLabel[c.defLang], constants.FailedMessageType, []string{constants.OkLabel[c.defLang]}, func(buttonIndex int, buttonLabel string) {
		if buttonLabel == constants.OkLabel[c.defLang] {
			c.pages.RemovePage(constants.ErrorNotifPageId)
			c.changePage(constants.AddDocPageId, c.addDocumentPage())
		}
	})
}
