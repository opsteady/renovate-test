package cicd

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/component"
)

// DockerCicd is a component for creating base Docker image
type DockerCicd struct {
	component.DefaultComponent
}

//sig-storage/csi-node-driver-registrar
// Initialize creates a new DockerCicd struct
func (d *DockerCicd) Initialize(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	buildArgs := map[string]string{
		"FROM_IMAGE": fmt.Sprintf(
			"%s/%s",
			d.GlobalConfig.ManagementDockerRegistry,
			"1.0.0", // renovate: datasource=docker registryUrl=opsteadyos.azurecr.io depName=opsteadyos.azurecr.io/base versioning=semver
		),
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
