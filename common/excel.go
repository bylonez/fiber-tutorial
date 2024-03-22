package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

type excelTagType struct {
	head string
}

func parseExcelTag(obj any) []excelTagType {
	t := reflect.TypeOf(obj)
	var titles []excelTagType
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := field.Tag.Get("excel")
		excelTag := excelTagType{}
		for _, i2 := range strings.Split(tags, ";") {
			kv := strings.Split(i2, ":")
			if len(kv) != 2 {
				panic("invalid excel tag")
			}
			tagKey := kv[0]
			tagValue := kv[1]
			switch tagKey {
			case "head":
				excelTag.head = tagValue
			default:
				panic("invalid excel tag")
			}
		}
		titles = append(titles, excelTag)
	}
	return titles
}

func WriteResponse[T any](objs []T, c *fiber.Ctx) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	// headline
	for i, v := range parseExcelTag(objs[0]) {
		cell, err := excelize.CoordinatesToCellName(i+1, 1)
		if err != nil {
			panic(err)
		}
		f.SetCellValue("Sheet1", cell, v.head)
	}
	// body
	for rowNum, obj := range objs {
		valueS := reflect.ValueOf(obj)
		for i := 0; i < valueS.NumField(); i++ {
			cell, err := excelize.CoordinatesToCellName(i+1, rowNum+2)
			if err != nil {
				panic(err)
			}
			f.SetCellValue("Sheet1", cell, valueS.Field(i))
		}
	}

	f.SetActiveSheet(index)

	// write response
	err = f.Write(c.Response().BodyWriter())
	if err != nil {
		panic(err)
	}

	// 设置响应头信息
	filename := "export.xlsx"
	c.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header.Set("Content-Disposition", "attachment; filename="+filename)
}
