// Package zmap implements functions for maps.
package zmap

// Reverse the keys and values of a map.
func Reverse(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
