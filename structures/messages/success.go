package messages

const (
	// SuccessMessageSignature is the signature byte for the SUCCESS message
	SuccessMessageSignature = 0x70
)

// SuccessMessage Represents an SUCCESS message
type SuccessMessage struct {
	metadata map[string]interface{}
}

// NewSuccessMessage Gets a new SuccessMessage struct
func NewSuccessMessage(metadata map[string]interface{}) SuccessMessage {
	return SuccessMessage{
		metadata: metadata,
	}
}

// Signature gets the signature byte for the struct
func (i SuccessMessage) Signature() int {
	return SuccessMessageSignature
}

// Fields gets the fields to encode for the struct
func (i SuccessMessage) Fields() []interface{} {
	return []interface{}{i.metadata}
}
