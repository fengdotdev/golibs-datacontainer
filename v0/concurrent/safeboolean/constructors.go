package safebolean

func New() SafeBoolean {
	return &GoBoolean{
		value: false,
	}
}

func NewValue(value bool) SafeBoolean {
	return &GoBoolean{
		value: value,
	}
}
func True() SafeBoolean {
	return &GoBoolean{
		value: true,
	}
}

func False() SafeBoolean {
	return &GoBoolean{
		value: false,
	}
}
