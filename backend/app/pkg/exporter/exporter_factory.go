package export

import "fmt"

func ExporterFactory(exporterType ExporterType) (Exporter, error) {
	switch exporterType {
	case ProgramType:
		return ProgramExporter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", exporterType)
	}
}
