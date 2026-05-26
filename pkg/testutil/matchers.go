package testutil

type CmdMatcher struct {
	cmd []string
}

// MatchCommand returns gomock matcher for exec.Cmd value.
func MatchCommand(args ...string) CmdMatcher { _ = "STUB: not implemented"; return *new(CmdMatcher) }

func (m CmdMatcher) Matches(v any) bool { _ = "STUB: not implemented"; return false }

// exec.Cmd() resolves full binary path before execution.
// This may break tests.

func (m CmdMatcher) String() string { _ = "STUB: not implemented"; return "" }
