package storage

const (
	// ExtWasm is wasm file extension
	ExtWasm = "wasm"
	// ExtGo is go file extension
	ExtGo = "go"
)

// ArtifactID represents artifact ID
type ArtifactID string

// Ext returns string with artifact ID and extension
func (a ArtifactID) Ext(ext string) string { _ = "STUB: not implemented"; return "" }

// String returns string
func (a ArtifactID) String() string {
	_ = "STUB: not implemented"

	// GetArtifactID generates new artifact ID from contents
	return ""
}

func GetArtifactID(entries map[string][]byte) (ArtifactID, error) {
	_ = "STUB: not implemented"

	// Keys have to be sorted for constant hashing
	return *new(ArtifactID), nil
}
