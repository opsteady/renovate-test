package cicd

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/component"
)

// DockerCicd is a component for creating base Docker image
type DockerCicd struct {
	component.DefaultComponent
}

// Initialize creates a new DockerCicd struct
func (d *DockerCicd) Initialize(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	// renovate: datasource=docker depName=base versioning=semver
	BASE_VERSION = "1.0.0"
	buildArgs := map[string]string{
		"FROM_IMAGE":               fmt.Sprintf("%s/%s", d.GlobalConfig.ManagementDockerRegistry, "base:2.0.0"),
		"VAULT_CA_STORAGE_ACCOUNT": d.GlobalConfig.VaultCaStorageAccountName,
	}
	d.SetDockerBuildInfo("cicd", "2.0.0", buildArgs)
}

func (c *DockerCicd) Deploy() {
	c.Logger.Info().Msg("Deploy not supported for this component")
}

func (c *DockerCicd) Destroy() {
	c.Logger.Info().Msg("Destroy not supported for this component")
}
