package update

import (
	"os"
	"path/filepath"

	"github.com/Phyrenos/U-Mangal/constant"
	"github.com/Phyrenos/U-Mangal/filesystem"
	"github.com/Phyrenos/U-Mangal/log"
)

type downloadedChapter struct {
	path   string
	format string
}

func getChapters(manga string) ([]*downloadedChapter, error) {
	log.Infof("getting chapters for %s", manga)
	var chapters []*downloadedChapter

	err := filesystem.Api().Walk(manga, func(path string, info os.FileInfo, err error) error {
		// we will ignore plain chapter (aka folder ones) for the sake of simplicity
		if info.IsDir() {
			return nil
		}

		name := info.Name()
		switch filepath.Ext(name)[1:] {
		case constant.FormatCBZ:
			chapters = append(chapters, &downloadedChapter{path: path, format: constant.FormatCBZ})
		case constant.FormatPDF:
			chapters = append(chapters, &downloadedChapter{path: path, format: constant.FormatPDF})
		case constant.FormatZIP:
			chapters = append(chapters, &downloadedChapter{path: path, format: constant.FormatZIP})
		}

		return nil
	})

	return chapters, err
}
