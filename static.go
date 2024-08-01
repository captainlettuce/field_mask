package field_mask

import "errors"

const (
	tagName = "field_mask"
)

var (
	ErrReceiverNotPointerToStruct = errors.New("field is not pointer to struct")
	ErrInvalidFieldMask           = errors.New("invalid field mask")
	ErrInvalidReceiverField       = errors.New("invalid receiver field")
	ErrUnsettableReceiver         = errors.New("Unsettable receiver")
)
