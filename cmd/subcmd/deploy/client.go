//go:build client

package deploy

func runDeploySystem() error {
	return runDeploy("client")
}
