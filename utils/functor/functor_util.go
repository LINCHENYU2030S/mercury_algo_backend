package utils

// Generic Map function
func Map[T any, R any](slice []T, f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Ptr[T any](v T) *T {
	return &v
}

func PaginateSlice[T any](elems []T, page, pageSize int) []T {
	if page < 1 || pageSize < 0 {
		return nil
	}
	start := (page - 1) * pageSize
	if start >= len(elems) {
		return nil
	}
	end := start + pageSize
	if end > len(elems) {
		end = len(elems)
	}
	return elems[start:end]
}
