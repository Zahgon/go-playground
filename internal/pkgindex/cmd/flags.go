package cmd

type globalFlags struct {
	goRoot       string
	heapProfFile string
}

func (f globalFlags) withDefaults() (globalFlags, error) {
	_ = "STUB: not implemented"
	return *new(globalFlags), nil
}

type importsFlags struct {
	*globalFlags
	verbose     bool
	prettyPrint bool
	stdout      bool
	outFile     string
}

func (f importsFlags) validate() error { _ = "STUB: not implemented"; return nil }
