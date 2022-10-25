package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func initGUI() {
	window := ui.NewWindow("Склопакет", 400, 200, false)
	widthEntry := ui.NewEntry()
	heightEntry := ui.NewEntry()
	materialComboBox := ui.NewCombobox()
	glazingComboBox := ui.NewCombobox()
	windowsillCheckBox := ui.NewCheckbox("Підвіконня")
	calculateButton := ui.NewButton("Розрахувати")
	valueLabel := ui.NewLabel("")
	box := ui.NewHorizontalBox()
	firstColumn := ui.NewVerticalBox()
	secondColumn := ui.NewVerticalBox()
	firstColumnFirstRow := ui.NewHorizontalBox()
	firstColumnSecondRow := ui.NewHorizontalBox()
	firstColumnThirdRow := ui.NewHorizontalBox()
	firstColumnFourthRow := ui.NewHorizontalBox()
	firstColumnLabelRow := ui.NewHorizontalBox()
	materialComboBox.Append("Дерево")
	materialComboBox.Append("Метал")
	materialComboBox.Append("Металопластик")
	materialComboBox.SetSelected(0)
	glazingComboBox.Append("Однокамерний")
	glazingComboBox.Append("Двокамерний")
	glazingComboBox.SetSelected(0)
	box.SetPadded(true)
	box.Append(firstColumn, false)
	box.Append(secondColumn, false)
	firstColumn.Append(firstColumnLabelRow, false)
	firstColumn.Append(ui.NewLabel(""), false)
	firstColumn.Append(firstColumnFirstRow, false)
	firstColumn.Append(ui.NewLabel(""), false)
	firstColumn.Append(firstColumnSecondRow, false)
	firstColumn.Append(ui.NewLabel(""), false)
	firstColumn.Append(firstColumnThirdRow, false)
	firstColumn.Append(ui.NewLabel(""), false)
	firstColumn.Append(firstColumnFourthRow, false)
	firstColumnLabelRow.Append(ui.NewLabel("Розміри вікна"), false)
	firstColumnFirstRow.Append(ui.NewLabel("Ширина(см):"), false)
	firstColumnSecondRow.Append(ui.NewLabel("  Висота(см):"), false)
	firstColumnThirdRow.Append(ui.NewLabel("     Матеріал:"), false)
	firstColumnFourthRow.Append(ui.NewLabel("       Вартість:  "), false)
	firstColumnLabelRow.Append(ui.NewLabel(""), false)
	firstColumnFirstRow.Append(widthEntry, false)
	firstColumnFirstRow.SetPadded(true)
	firstColumnSecondRow.Append(heightEntry, false)
	firstColumnSecondRow.SetPadded(true)
	firstColumnThirdRow.Append(materialComboBox, false)
	firstColumnThirdRow.SetPadded(true)
	firstColumnFourthRow.Append(valueLabel, false)
	secondColumn.Append(ui.NewLabel("\tСклопакет"), false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(glazingComboBox, false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(windowsillCheckBox, false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(calculateButton, false)
	window.SetChild(box)
	window.SetMargined(true)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	calculateButton.OnClicked(func(*ui.Button) {
		var result float64
		width, errWidth := strconv.ParseFloat(widthEntry.Text(), 64)
		height, errHeight := strconv.ParseFloat(heightEntry.Text(), 64)
		if glazingComboBox.Selected() == 0 && materialComboBox.Selected() == 0 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.25
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		} else if glazingComboBox.Selected() == 1 && materialComboBox.Selected() == 0 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.30
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		} else if glazingComboBox.Selected() == 0 && materialComboBox.Selected() == 1 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.05
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		} else if glazingComboBox.Selected() == 1 && materialComboBox.Selected() == 1 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.10
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		} else if glazingComboBox.Selected() == 0 && materialComboBox.Selected() == 2 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.15
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		} else if glazingComboBox.Selected() == 1 && materialComboBox.Selected() == 2 {
			if errWidth == nil && errHeight == nil {
				result = width * height * 0.20
			} else if errWidth != nil {
				fmt.Println(errWidth)
			} else {
				fmt.Println(errHeight)
			}
		}
		if windowsillCheckBox.Checked() {
			result += 35
		}
		valueLabel.SetText(strconv.FormatFloat(result, 'f', 3, 64) + " грн.")
	})
	window.Show()
}
func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}
