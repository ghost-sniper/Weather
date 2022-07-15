package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	os.Setenv("FYNE_FONT", "c:\\WINDOWS\\FONTS\\msyh.ttc")
	App := app.New()
	MainWindow := App.NewWindow("Weather")
	MainWindow.Resize(fyne.Size{
		Width:  500,
		Height: 100,
	})

	CityEntry := widget.NewEntry()
	ProvinceEntry := widget.NewEntry()
	CountryEntry := widget.NewEntry()

	ProvinceEntry.SetText("省份")
	CityEntry.SetText("城市")
	CountryEntry.SetText("乡镇")

	ProvinceLabel := widget.NewLabel("省份:")
	CityLabel := widget.NewLabel("城市:")
	CountryLabel := widget.NewLabel("象征:")

	QueryButton := widget.NewButton("查询", func() {
		fmt.Println("查询天气中")
	})

	content := container.New(layout.NewGridLayout(7), ProvinceLabel, ProvinceEntry, CityLabel, CityEntry, CountryLabel, CountryEntry, QueryButton)
	MainWindow.SetContent(content)
	MainWindow.ShowAndRun()
}
