package miutil

import (
	"fmt"
	"log"
	"regexp"
)

func OnlyNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

func NewError(cause, err error) error {
	return fmt.Errorf("err: %w cause: %v", err, cause)
}

func AvoidNil[T comparable](t *T) T {
	if t != nil {
		return *t
	}
	var r T
	return r
}

func NilSlice[T comparable](ss *[]T) []T {
	if ss != nil && len(*ss) > 0 {
		return *ss
	}
	return []T{}
}

func ChunkSlice[T comparable](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
