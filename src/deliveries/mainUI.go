package deliveries

import (
	"config"
	"constants"
	"entitites"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"repositories"
	"strconv"
	"useCases"
)

var documents = make([]*entitites.Document, 0, 1)

type ConsoleUI struct {
	defLang    string
	docUseCase useCases.DocumentUseCase
	app        *tview.Application
	pages      *tview.Pages
}

func (c *ConsoleUI) Run() {
	c.pages = tview.NewPages().
		AddPage(constants.MainPageId, c.mainPage(), true, true)

	if err := c.app.SetRoot(c.pages, true).Run(); err != nil {
		panic(err)
	}
}

func InitUi(config *config.Config) *ConsoleUI {
	var repo repositories.IRepository
	if config.FileType == "json" {
		repo = repositories.JsonRepository{
			FileName: config.FileName,
		}
	}

	var docRepo = repositories.DocumentFileRepository{
		Repo: repo,
	}
	return &ConsoleUI{
		defLang: config.DefLang,
		docUseCase: useCases.DocumentUseCase{
			docRepo,
		},
		app: tview.NewApplication(),
	}
}

func (c *ConsoleUI) mainPage() (flex *tview.Flex) {
	list := tview.NewList().
		AddItem(constants.AddDocLabel[c.defLang], constants.AddDoc2ndLabel[c.defLang], 'd', func() {
			c.changePage(constants.AddDocPageId, c.addDocumentPage())
		}).
		AddItem(constants.ViewDocLabel[c.defLang], constants.ViewDoc2ndLabel[c.defLang], 'l', func() {
			list := c.docUseCase.FindAll()
			c.changePage(constants.ViewListDocPageId, c.listDocumentPage(list))
		}).
		AddItem(constants.ExitLabel[c.defLang], constants.Exit2ndLabel[c.defLang], 'q', func() {
			c.app.Stop()
		})
	list.SetBorder(true).SetTitle(constants.AppName[c.defLang])

	t := tview.NewTable()
	t.SetBorder(true).SetTitle("Dashboard")
	t.SetCell(0, 0, tview.NewTableCell(constants.DocTotalLabel[c.defLang]))
	t.SetCell(0, 1, tview.NewTableCell(":"))
	t.SetCell(0, 2, tview.NewTableCell(strconv.Itoa(c.docUseCase.GetTotalDocument())))
	flex = tview.NewFlex().
		AddItem(list, 0, 2, true).
		AddItem(t, 0, 4, false)
	return
}

func (c *ConsoleUI) changePage(idHalaman string, komponen tview.Primitive) {
	c.pages.RemovePage(idHalaman)
	c.pages.AddAndSwitchToPage(idHalaman, komponen, true)
}

func (c *ConsoleUI) navigation(judul string) (nav *tview.List) {
	nav = tview.NewList().
		AddItem(constants.HomeLabel[c.defLang], "", 'h', func() {
			c.pages.SwitchToPage(constants.MainPageId)
		})
	nav.SetBorder(true).SetTitle(judul)
	return
}

func (c *ConsoleUI) notificationModal(message string, messageType int, buttons []string, callback func(int, string)) (modal *tview.Modal) {
	color := tcell.ColorGreen
	if messageType == constants.FailedMessageType {
		color = tcell.ColorRed
	}
	modal = tview.NewModal().
		SetBackgroundColor(color).
		SetText(message).
		AddButtons(buttons).
		SetDoneFunc(callback)
	return
}

func (c *ConsoleUI) requiredLabel(label string) string {
	return fmt.Sprintf("%s %s", label, "*")
}
