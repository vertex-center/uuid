package uuid

// Type and Format allows supporting https://github.com/wI2L/fizz

func (NullUUID) Type() string   { return "string" }
func (NullUUID) Format() string { return "uuid" }

func (UUID) Type() string   { return "string" }
func (UUID) Format() string { return "uuid" }
