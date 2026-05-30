// Package tokens provides token persistence and adapters for QBO clients.
package tokens

// Source is a placeholder adapter that will wrap Service for refresh-aware tokens.
type Source struct{}

// AccessToken is a stub implementation to satisfy the interface scaffold.
func (s *Source) AccessToken() (string, error) {
	return "", nil
}
