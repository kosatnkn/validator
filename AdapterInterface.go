package validator

// AdapterInterface is implemented by all validator adapters.
type AdapterInterface interface {

	// Validate validates bound values of an unpacker struct against
	// validation rules defined in that unpacker struct.
	Validate(data interface{}) map[string]string

	// ValidateField validates a single variable.
	ValidateField(field interface{}, rules string) map[string]string
}
