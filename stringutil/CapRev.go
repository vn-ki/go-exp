package stringutil

// CapRev returns it's argument string capitalized and reversed.
// This is soo coool.. It's almost like package stringutil is one big file.
func CapRev(s string) string {
	return Reverse(Capitalize(s))
}
