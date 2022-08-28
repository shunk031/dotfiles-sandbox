//go:build server

package deploy

func RunDeploy() error {
	return runDeploy("server")
}
