package parser

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx"
)

const timeZoneUtc = 8 //irkutsk

func SingleSheetExcelFile(pg *Page) {
	p := *pg
	applicants := p.Applicants
	t := p.Time
	t = t.Add(time.Hour * timeZoneUtc)
	h, mi, _ := t.Clock()
	y, mo, d := t.Date()
	date := fmt.Sprintf("%v.%v.%v", d, mo, y) + "   " + fmt.Sprintf("%02d:%02d", h, mi)
	plan := p.Planned
	text := p.Title
	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("Sheet!")
	if err != nil {
		log.Fatal(err)
	}
	row := sheet.AddRow()
	row.AddCell().SetString("Направление")
	row.AddCell().SetString(text)
	row.AddCell().SetString("Обновление")
	row.AddCell().SetString(date)
	row = sheet.AddRow()
	row.AddCell().SetString("Подали")
	row.AddCell().SetInt(applicants)
	row.AddCell().SetString("Пройдет")
	row.AddCell().SetInt(plan)
	row = sheet.AddRow()
	row.AddCell().SetString("номер ")
	row.AddCell().SetString("приоритет")
	row.AddCell().SetString("согласие")
	row.AddCell().SetString("балл")

	for _, v := range p.List {
		row := sheet.AddRow()
		row.AddCell().SetInt(v.Id)
		row.AddCell().SetInt(v.Priority)
		if v.Acceptance {
			row.AddCell().SetString("+")
		} else {
			row.AddCell().SetString("-")
		}
		row.AddCell().SetInt(v.Sum)
	}
	wb.Save("asd" + ".xlsx")
}

func MultiSheetExcelFile(pg []*Page, name string) {
	wb := xlsx.NewFile()
	i := 0
	for _, v := range pg {
		i++
		p := *v
		applicants := p.Applicants
		t := p.Time
		t = t.Add(time.Hour * timeZoneUtc)
		h, mi, _ := t.Clock()
		y, mo, d := t.Date()
		date := fmt.Sprintf("%v.%v.%v", d, mo, y) + "   " + fmt.Sprintf("%02d:%02d", h, mi)
		plan := p.Planned
		text := p.Title
		sheet, err := wb.AddSheet(fmt.Sprintf("pg %v", i))
		if err != nil {
			log.Fatal(err)
		}
		row := sheet.AddRow()
		row.AddCell().SetString("Направление")
		row.AddCell().SetString(text)
		row.AddCell().SetString("Обновление")
		row.AddCell().SetString(date)
		row = sheet.AddRow()
		row.AddCell().SetString("Подали")
		row.AddCell().SetInt(applicants)
		row.AddCell().SetString("Пройдет")
		row.AddCell().SetInt(plan)
		row = sheet.AddRow()
		row.AddCell().SetString("номер ")
		row.AddCell().SetString("приоритет")
		row.AddCell().SetString("согласие")
		row.AddCell().SetString("балл")

		for _, v := range p.List {
			row := sheet.AddRow()
			row.AddCell().SetInt(v.Id)
			row.AddCell().SetInt(v.Priority)
			if v.Acceptance {
				row.AddCell().SetString("+")
			} else {
				row.AddCell().SetString("-")
			}
			row.AddCell().SetInt(v.Sum)
		}
	}
	wb.Save(name + ".xlsx")
}
