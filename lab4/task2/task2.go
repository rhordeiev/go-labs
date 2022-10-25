package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func initGUI() {
	window := ui.NewWindow("Тур на відпочинок", 400, 200, false)
	dayCountEntry := ui.NewEntry()
	countryComboBox := ui.NewCombobox()
	seasonComboBox := ui.NewCombobox()
	marginCheckBox := ui.NewCheckbox("Номер-люкс")
	individualGuideCheckBox := ui.NewCheckbox("Індивідуальний гід")
	calculateButton := ui.NewButton("Розрахувати")
	valueLabel := ui.NewLabel("")
	box := ui.NewHorizontalBox()
	firstColumn := ui.NewVerticalBox()
	secondColumn := ui.NewVerticalBox()
	firstColumnFirstRow := ui.NewHorizontalBox()
	firstColumnSecondRow := ui.NewHorizontalBox()
	firstColumnThirdRow := ui.NewHorizontalBox()
	firstColumnLabelRow := ui.NewHorizontalBox()
	countryComboBox.Append("Болгарія")
	countryComboBox.Append("Германія")
	countryComboBox.Append("Польща")
	countryComboBox.SetSelected(0)
	seasonComboBox.Append("Літо")
	seasonComboBox.Append("Зима")
	seasonComboBox.SetSelected(0)
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
	firstColumnLabelRow.Append(ui.NewLabel(""), false)
	firstColumnFirstRow.Append(ui.NewLabel("Кількість днів:"), false)
	firstColumnSecondRow.Append(ui.NewLabel("             Країна:"), false)
	firstColumnThirdRow.Append(ui.NewLabel("          Вартість:  "), false)
	firstColumnLabelRow.Append(ui.NewLabel(""), false)
	firstColumnFirstRow.Append(dayCountEntry, false)
	firstColumnFirstRow.SetPadded(true)
	firstColumnSecondRow.Append(countryComboBox, false)
	firstColumnSecondRow.SetPadded(true)
	firstColumnThirdRow.Append(valueLabel, false)
	secondColumn.Append(ui.NewLabel("\tПора року"), false)
	secondColumn.Append(ui.NewLabel(""), false)
	secondColumn.Append(seasonComboBox, false)
	secondColumn.Append(individualGuideCheckBox, false)
	secondColumn.Append(marginCheckBox, false)
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
		dayCount, err := strconv.ParseFloat(dayCountEntry.Text(), 64)
		if seasonComboBox.Selected() == 0 && countryComboBox.Selected() == 0 {
			if err == nil {
				result = dayCount * 100
			} else {
				fmt.Println(err)
			}
		} else if seasonComboBox.Selected() == 1 && countryComboBox.Selected() == 0 {
			if err == nil {
				result = dayCount * 150
			} else {
				fmt.Println(err)
			}
		} else if seasonComboBox.Selected() == 0 && countryComboBox.Selected() == 1 {
			if err == nil {
				result = dayCount * 160
			} else {
				fmt.Println(err)
			}
		} else if seasonComboBox.Selected() == 1 && countryComboBox.Selected() == 1 {
			if err == nil {
				result = dayCount * 200
			} else {
				fmt.Println(err)
			}
		} else if seasonComboBox.Selected() == 0 && countryComboBox.Selected() == 2 {
			if err == nil {
				result = dayCount * 120
			} else {
				fmt.Println(err)
			}
		} else if seasonComboBox.Selected() == 1 && countryComboBox.Selected() == 2 {
			if err == nil {
				result = dayCount * 180
			} else {
				fmt.Println(err)
			}
		}
		if individualGuideCheckBox.Checked() {
			result += 50 * dayCount
		}
		if marginCheckBox.Checked() {
			result += result * 0.2
		}
		valueLabel.SetText(strconv.FormatFloat(result, 'f', 3, 64) + " $")
	})
	window.Show()
}
func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}
