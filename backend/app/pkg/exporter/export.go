package export

import "bytes"

type Exporter interface {
	Export(data interface{}) (*bytes.Buffer, error)
}

type ExporterType string

const (
	ProgramType ExporterType = "program"
)

func IndexToColumnLetter(index int) string {
	var columnLetter string
	for index > 0 {
		index-- // Convert to 0-based index
		columnLetter = string('A'+(index%26)) + columnLetter
		index /= 26
	}
	return columnLetter
}
