package export

import (
	// repo "EQA/backend/app/repository"
	dto "EQA/backend/app/domain/dto"
	"bytes"
	"fmt"

	// "fmt"

	"github.com/xuri/excelize/v2"
)

type ProgramExporter struct{}

func (e ProgramExporter) Export(data interface{}) (*bytes.Buffer, error) {
	file := excelize.NewFile()
	sheet := "Sheet1"
	file.NewSheet(sheet)

	style, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, err
	}

	// file.MergeCell(sheet, "A1", "A2")
	// file.MergeCell(sheet, "B1", "B2")
	// file.MergeCell(sheet, "C1", "C2")
	// file.MergeCell(sheet, "D1", "D2")
	// file.MergeCell(sheet, "E1", "E2")
	// file.MergeCell(sheet, "F1", "F2")
	// file.MergeCell(sheet, "G1", "G2")

	// file.MergeCell(sheet, "I1", "M1")
	// file.MergeCell(sheet, "N1", "O1")

	file.SetCellStyle(sheet, "A1", "I1", style)

	file.SetCellValue(sheet, "A1", "TT")
	file.SetCellValue(sheet, "B1", "Tên chương trình")
	file.SetCellValue(sheet, "C1", "Nhà cung cấp")
	file.SetCellValue(sheet, "D1", "Mã")
	file.SetCellValue(sheet, "E1", "Đợt")
	file.SetCellValue(sheet, "F1", "Ngày nhận mẫu")
	file.SetCellValue(sheet, "G1", "Ngày trả kết quả")
	file.SetCellValue(sheet, "H1", "Ngày nhận kết quả")
	file.SetCellValue(sheet, "I1", "Kết quả")

	records, _ := data.([]dto.ProgramExportResp)
	rowNum := 2
	for index, program := range records {
		for i := 0; i < 9; i++ {
			columnLetter := IndexToColumnLetter(i + 1)
			cell := fmt.Sprintf("%s%d", columnLetter, rowNum)
			switch i {
			case 0:
				file.SetCellValue(sheet, cell, index+1)
			case 1:
				file.SetCellValue(sheet, cell, program.ProgramName)
			case 2:
				file.SetCellValue(sheet, cell, program.ProviderName)
			case 3:
				file.SetCellValue(sheet, cell, program.Code)
			case 4:
				var sample interface{}
				if program.Sample < 10 && program.Sample > 0 {
					sample = fmt.Sprintf("%d%d", 0, program.Sample)
				} else {
					sample = fmt.Sprintf("%d", program.Sample)
				}
				file.SetCellValue(sheet, cell, sample)
			case 5:
				file.SetCellValue(sheet, cell, program.DateOfReceive.Format("02/01/2006"))
			case 6:
				file.SetCellValue(sheet, cell, program.DateOfReturn.Format("02/01/2006"))
			case 7:
				if program.DateOfReceiveResult != nil {
					file.SetCellValue(sheet, cell, program.DateOfReceiveResult.Format("02/01/2006"))
				} else {
					file.SetCellValue(sheet, cell, "")
				}
			case 8:
				var status interface{}
				if program.Status == 1 {
					status = fmt.Sprintf("Đạt %d%%", program.PercentPassed)
				} else {
					status = ""
				}
				file.SetCellValue(sheet, cell, status)
			}
		}
		rowNum += 1
	}

	b, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return b, nil
}
