package provider

import (
	"github.com/Phyrenos/U-Mangal/provider/generic"
	"github.com/Phyrenos/U-Mangal/provider/mangadex"
	"github.com/Phyrenos/U-Mangal/provider/manganato"
	"github.com/Phyrenos/U-Mangal/provider/manganelo"
	"github.com/Phyrenos/U-Mangal/provider/mangapill"
	"github.com/Phyrenos/U-Mangal/source"
)

const CustomProviderExtension = ".lua"

var builtinProviders = []*Provider{
	{
		ID:   mangadex.ID,
		Name: mangadex.Name,
		CreateSource: func() (source.Source, error) {
			return mangadex.New(), nil
		},
	},
}

func init() {
	for _, conf := range []*generic.Configuration{
		manganelo.Config,
		manganato.Config,
		mangapill.Config,
	} {
		conf := conf
		builtinProviders = append(builtinProviders, &Provider{
			ID:   conf.ID(),
			Name: conf.Name,
			CreateSource: func() (source.Source, error) {
				return generic.New(conf), nil
			},
		})
	}
}
