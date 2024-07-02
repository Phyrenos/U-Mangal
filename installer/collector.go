package installer

import (
	"path/filepath"

	"github.com/Phyrenos/U-Mangal/key"
	"github.com/Phyrenos/U-Mangal/util"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

var collector *githubFilesCollector

// Scrapers gets available scrapers from GitHub repo.
// See https://github.com/Phyrenos/U-Mangal-scrapers
func Scrapers() ([]*Scraper, error) {
	if collector == nil {
		setupCollector()
	}

	err := collector.collect()
	if err != nil {
		return nil, err
	}

	return lo.FilterMap(collector.Files, func(f *GithubFile, _ int) (*Scraper, bool) {
		if filepath.Ext(f.Path) != ".lua" {
			return nil, false
		}

		return &Scraper{
			Name: util.FileStem(filepath.Base(f.Path)),
			URL:  f.Url,
		}, true
	}), nil
}

func setupCollector() {
	collector = &githubFilesCollector{
		user:   viper.GetString(key.InstallerUser),
		repo:   viper.GetString(key.InstallerRepo),
		branch: viper.GetString(key.InstallerBranch),
	}
}
