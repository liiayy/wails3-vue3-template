package system

// SetAutostart 被各平台实现
func SetAutostart(appName string, enabled bool) error {
	return setAutostart(appName, enabled)
}

// IsAutostartEnabled 被各平台实现
func IsAutostartEnabled(appName string) (bool, error) {
	return isAutostartEnabled(appName)
}
