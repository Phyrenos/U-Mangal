package converter

import (
	"fmt"
	"strings"

	"github.com/Phyrenos/U-Mangal/constant"
	"github.com/Phyrenos/U-Mangal/converter/cbz"
	"github.com/Phyrenos/U-Mangal/converter/pdf"
	"github.com/Phyrenos/U-Mangal/converter/plain"
	"github.com/Phyrenos/U-Mangal/converter/zip"
	"github.com/Phyrenos/U-Mangal/source"
	"github.com/samber/lo"
)

// Converter is the interface that all converters must implement.
type Converter interface {
	Save(chapter *source.Chapter) (string, error)
	SaveTemp(chapter *source.Chapter) (string, error)
}

var converters = map[string]Converter{
	constant.FormatPlain: plain.New(),
	constant.FormatCBZ:   cbz.New(),
	constant.FormatPDF:   pdf.New(),
	constant.FormatZIP:   zip.New(),
}

// Available returns a list of available converters.
func Available() []string {
	return lo.Keys(converters)
}

// Get returns a converter by name.
// If the converter is not available, an error is returned.
func Get(name string) (Converter, error) {
	if converter, ok := converters[name]; ok {
		return converter, nil
	}

	return nil, fmt.Errorf("unkown format \"%s\", available options are %s", name, strings.Join(Available(), ", "))
}
