//go:build client

package deploy

func RunDeploySystem() error {
	return runDeploy("client")
}
