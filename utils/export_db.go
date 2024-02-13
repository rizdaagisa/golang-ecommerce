package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

func Export_DB(data interface{}, sheet_name string, filename string, destination string, file_extension string) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}

	// Ensure slice is not empty
	if val.Len() == 0 {
		return fmt.Errorf("no data to export")
	}

	// Get type of elements in slice
	elementType := val.Index(0).Type()

	file := excelize.NewFile()
	file.SetSheetName(file.GetSheetName(0), sheet_name)

	// Write header row
	for i := 0; i < elementType.NumField(); i++ {
		field := elementType.Field(i)
		col := 'A' + i
		file.SetCellValue(sheet_name, fmt.Sprintf("%c1", col), field.Name)
	}

	// Write data rows
	for rowIndex := 0; rowIndex < val.Len(); rowIndex++ {
		elem := val.Index(rowIndex)
		for colIndex := 0; colIndex < elementType.NumField(); colIndex++ {
			field := elem.Field(colIndex)
			col := 'A' + colIndex
			file.SetCellValue(sheet_name, fmt.Sprintf("%c%d", col, rowIndex+2), field.Interface())
		}
	}

	errs := file.SaveAs(destination + filename + "." + file_extension)
	if errs != nil {
		// Handle error
		fmt.Println("Failed to save Excel file:", errs)
		return errs
	}

	return nil
}
