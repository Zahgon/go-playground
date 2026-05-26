package osutil

const envVarsDelimiter = "="

// EnvironmentVariables is a key-value pair of environment variable and value.
type EnvironmentVariables map[string]string

// Join returns slice of environment variable and values joined by delimiter (=).
func (s EnvironmentVariables) Join() []string { _ = "STUB: not implemented"; return nil }

// Append appends new values to existing item.
func (s EnvironmentVariables) Append(items EnvironmentVariables) { _ = "STUB: not implemented"; return }

// Concat joins two items together into a new one.
func (s EnvironmentVariables) Concat(newItems EnvironmentVariables) EnvironmentVariables {
	_ = "STUB: not implemented"
	return *new(EnvironmentVariables)
}

// SplitEnvironmentValues splits slice of '='-separated key-value items and returns key-value pair.
//
// Second optional parameter allows filter items.
func SplitEnvironmentValues(vals []string, filterKeys ...string) EnvironmentVariables {
	_ = "STUB: not implemented"
	return *new(EnvironmentVariables)
}

// SelectEnvironmentVariables selects environment variables as key-value pair with whitelist filter.
func SelectEnvironmentVariables(filterKeys ...string) EnvironmentVariables {
	_ = "STUB: not implemented"
	return *new(EnvironmentVariables)
}
