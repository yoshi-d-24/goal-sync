package core

import (
	"reflect"
)

type ValueObject[T any] struct {
	value T
}

func (v ValueObject[T]) Value() T {
	return v.value
}

func NewValueObject[T any](value T) ValueObject[T] {
	return ValueObject[T]{value: value}
}

func (v *ValueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(v.value, other.value)
}
