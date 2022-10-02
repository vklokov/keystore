package utils

func ArrayLast[T comparable](entity []T) T {
	return entity[len(entity)-1]
}
