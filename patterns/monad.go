package patterns

// Option type to handle missing values
type Option[T any] struct {
	value *T
}

// Some Create a new Option
func Some[T any](v T) Option[T] {
	return Option[T]{value: &v}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// Map Apply a function inside the monad
func (o Option[T]) Map(f func(T) T) Option[T] {
	if o.value == nil {
		return None[T]()
	}
	v := f(*o.value)
	return Some(v)
}

// GetOrElse Get the value or return default
func (o Option[T]) GetOrElse(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}
