package cicd

import "github.com/opsteady/opsteady/cli/component"

// DockerBase is a component for creating base Docker image
type DockerBase struct {
	component.DefaultComponent
}

// Initialize creates a new DockerBase struct
func (d *DockerBase) Initialize(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	d.SetDockerBuildInfo("base", "2.4.0", nil)
}

func (c *DockerBase) Deploy() {
	c.Logger.Info().Msg("Deploy not supported for this component")
}

func (c *DockerBase) Destroy() {
	c.Logger.Info().Msg("Destroy not supported for this component")
}
