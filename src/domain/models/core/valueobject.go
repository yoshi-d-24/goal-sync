package core

import "reflect"

type IValueObject[T any] interface {
	Value() T
	Equals(other IValueObject[T]) bool
}

type ValueObject[T any] struct {
	value T
}

func (v *ValueObject[T]) Value() T {
	return v.value
}

func NewValueObject[T any](value T) IValueObject[T] {
	return &ValueObject[T]{value: value}
}

func (v *ValueObject[T]) Equals(other IValueObject[T]) bool {
	return reflect.DeepEqual(v.Value(), other.Value())
}
