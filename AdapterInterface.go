package validator

// AdapterInterface is implemented by all validator adapters.
type AdapterInterface interface {
	// Validate validates fields of a struct.
	Validate(data interface{}) map[string]string

	// ValidateField validates a single variable.
	ValidateField(name string, value interface{}, rules string) map[string]string
}
