package main

import (
	"github.com/Phyrenos/U-Mangal/cmd"
	"github.com/Phyrenos/U-Mangal/config"
	"github.com/Phyrenos/U-Mangal/log"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	cmd.Execute()
}
