//go:build ubuntu

package initialize

func RunInitCmd() error {
	// runInitCommon()

	return runInitUbuntu()
}
