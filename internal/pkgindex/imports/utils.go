package imports

// ResolveGoRoot tries to find GOROOT path from environment variables or using "go env" command.
func ResolveGoRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func IsDirIgnored(basename string, ignoreVendor bool) bool { _ = "STUB: not implemented"; return false }

func IsVendorDir(dirName string) bool { _ = "STUB: not implemented"; return false }
