//go:build server

package deploy

func runDeploySystem() error {
	return runDeploy("server")
}
