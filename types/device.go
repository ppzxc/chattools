package types

const (
	OsMac        = "mac"
	OsIOS        = "ios"
	OsWindows    = "windows"
	OsUnix       = "unix"
	OsLinux      = "linux"
	OsAndroid    = "android"
	OsBlackBerry = "blackberry"
	OsETC        = "etc"
)

const (
	PlatformDesktop = "desktop"
	PlatformMobile  = "mobile"
)

func IsOsContains(os string) error {
	switch os {
	case OsMac, OsIOS, OsWindows, OsUnix, OsLinux, OsAndroid, OsBlackBerry, OsETC:
		return nil
	default:
		return ErrInvalidRequestUserDeviceInfo
	}
}

func IsPlatformContains(os string) error {
	switch os {
	case PlatformDesktop, PlatformMobile:
		return nil
	default:
		return ErrInvalidRequestUserDeviceInfo
	}
}
