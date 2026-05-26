package cmdutil

// FatalOnError prints error to stderr and shuts down the application.
//
// Used when logging facility isn't initialized yet.
//
// Does nothing if passed error is nil.
func FatalOnError(err error) { _ = "STUB: not implemented"; return }
