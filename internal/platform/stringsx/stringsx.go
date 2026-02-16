package stringsx

import "strings"

// NormalizeSpace collapses all whitespace (spaces/tabs/newlines) into single spaces.
// WHY: input normalization prevents subtle validation bugs and inconsistent behavior.
func NormalizeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// IsBlank returns true if the string is empty or whitespace-only.
// WHY: avoid scattered TrimSpace checks across business logic.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
