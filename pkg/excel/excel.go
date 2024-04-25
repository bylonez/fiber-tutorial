package excel

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type exportTag struct {
	Head string `json:"head"`
}

type importTag struct {
	Head string `json:"head"`
}

func parseTag[tagT, objT any]() []*tagT {
	objType := reflect.TypeFor[objT]()
	var tags []*tagT
	for iObj := 0; iObj < objType.NumField(); iObj++ {
		objField := objType.Field(iObj)
		tagsStr := objField.Tag.Get("excel")
		tagMap := make(map[string]string)
		for _, i2 := range strings.Split(tagsStr, ";") {
			kv := strings.Split(i2, ":")
			if len(kv) != 2 {
				panic("invalid tag")
			}
			tagMap[kv[0]] = kv[1]
		}
		excelTag := new(tagT)
		tagType := reflect.TypeFor[tagT]()
		tagValue := reflect.ValueOf(excelTag).Elem()
		for iTag := 0; iTag < tagType.NumField(); iTag++ {
			field := tagType.Field(iTag)
			strValue := tagMap[field.Tag.Get("json")]
			if strValue != "" {
				switch field.Type.Kind() {
				case reflect.String:
					tagValue.FieldByName(field.Name).SetString(strValue)
				case reflect.Int:
					intValue, err := strconv.ParseInt(strValue, 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					tagValue.FieldByName(field.Name).SetInt(intValue)
				default:
					panic("unknown type")
				}
			}
		}
		tags = append(tags, excelTag)
	}
	return tags
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
	for i, v := range parseTag[exportTag, T]() {
		cell, err := excelize.CoordinatesToCellName(i+1, 1)
		if err != nil {
			panic(err)
		}
		f.SetCellValue("Sheet1", cell, v.Head)
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

func ParseExcel[T any](c *fiber.Ctx, key string) []*T {
	formFile, err := c.FormFile(key)
	if err != nil {
		log.Fatal(err)
	}

	file, err := formFile.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}

	var dtos []*T
	for _, sheetName := range xlsx.GetSheetList() {
		rows, err := xlsx.GetRows(sheetName)
		if err != nil {
			log.Fatal(err)
		}

		// parse headline
		headerMap := make(map[string]int)
		for i, cell := range rows[0] {
			headerMap[cell] = i
		}

		m := make(map[int]int)
		for i, v := range parseTag[importTag, T]() {
			headI, ok := headerMap[v.Head]
			if ok {
				m[i] = headI
			}
		}

		dtoTypeR := reflect.TypeFor[T]()
		for _, row := range rows[1:] {
			dto := new(T)
			dtoValue := reflect.ValueOf(dto).Elem()
			for i := 0; i < dtoTypeR.NumField(); i++ {
				field := dtoTypeR.Field(i)
				colIndex, ok := m[i]
				if ok {
					if colIndex >= len(row) {
						continue
					}
					cellValue := row[colIndex]
					switch field.Type.Kind() {
					case reflect.String:
						dtoValue.Field(i).SetString(cellValue)
					case reflect.Int:
						intValue := 0
						fmt.Sscanf(cellValue, "%d", &intValue)
						dtoValue.Field(i).SetInt(int64(intValue))
					case reflect.Uint:
						var intValue uint
						fmt.Sscanf(cellValue, "%d", &intValue)
						dtoValue.Field(i).SetUint(uint64(intValue))
					default:
						panic("unhandled default case")
					}
				}
			}
			dtos = append(dtos, dto)
		}
	}
	return dtos
}
