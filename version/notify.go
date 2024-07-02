package version

import (
	"fmt"

	"github.com/Phyrenos/U-Mangal/color"
	"github.com/Phyrenos/U-Mangal/constant"
	"github.com/Phyrenos/U-Mangal/icon"
	"github.com/Phyrenos/U-Mangal/key"
	"github.com/Phyrenos/U-Mangal/style"
	"github.com/Phyrenos/U-Mangal/util"
	"github.com/spf13/viper"
)

func Notify() {
	if !viper.GetBool(key.CliVersionCheck) {
		return
	}

	erase := util.PrintErasable(fmt.Sprintf("%s Checking if new version is available...", icon.Get(icon.Progress)))
	version, err := Latest()
	erase()
	if err == nil {
		comp, err := Compare(version, constant.Version)
		if err == nil && comp <= 0 {
			return
		}
	}

	fmt.Printf(`
%s New version is available %s %s
%s

`,
		style.Fg(color.Green)("▇▇▇"),
		style.Bold(version),
		style.Faint(fmt.Sprintf("(You're on %s)", constant.Version)),
		style.Faint("https://github.com/Phyrenos/U-Mangal/releases/tag/v"+version),
	)

}
