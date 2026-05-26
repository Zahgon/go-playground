package webutil

import (
	"regexp"
)

var gtagRegExp = regexp.MustCompile(`(?i)^[A-Z]{2}-[A-Z0-9\-\+]+$`)

// ValidateGTag validates Google Analytics tag ID
func ValidateGTag(gtag string) error { _ = "STUB: not implemented"; return nil }
