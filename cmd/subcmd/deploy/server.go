//go:build server

package deploy

func RunDeploySystem() error {
	return runDeploy("server")
}
