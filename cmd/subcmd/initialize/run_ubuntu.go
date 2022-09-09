//go:build linux

package initialize

func RunInitCmd() error {
	if err := runInitCommon(); err != nil {
		return err
	}
	if err := runInitUbuntuCommon(); err != nil {
		return err
	}
	if err := runInitUbuntuSystem(); err != nil {
		return err
	}
	return nil
}
