package main

import (
	"log"
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/paketo-buildpacks/packit/v2/fs"
	"github.com/paketo-buildpacks/packit/v2/pexec"
	"github.com/paketo-buildpacks/packit/v2/sbom"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/paketo-buildpacks/packit/v2/servicebindings"

	yarninstall "github.com/initializ-buildpacks/yarn-install"
)

type SBOMGenerator struct{}

func (s SBOMGenerator) Generate(path string) (sbom.SBOM, error) {
	return sbom.Generate(path)
}

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	installProcess := yarninstall.NewYarnInstallProcess(pexec.NewExecutable("yarn"), fs.NewChecksumCalculator(), logger)
	sbomGenerator := SBOMGenerator{}
	symlinker := yarninstall.NewSymlinker()
	packageManagerConfigurationManager := yarninstall.NewPackageManagerConfigurationManager(servicebindings.NewResolver(), logger)
	entryResolver := draft.NewPlanner()
	home, err := os.UserHomeDir()
	tmpDir := os.TempDir()
	if err != nil {
		// not tested
		log.Fatal(err)
	}

	packit.Run(
		yarninstall.Detect(),
		yarninstall.Build(entryResolver,
			packageManagerConfigurationManager,
			home,
			symlinker,
			installProcess,
			sbomGenerator,
			chronos.DefaultClock,
			logger,
			tmpDir,
		),
	)
}
