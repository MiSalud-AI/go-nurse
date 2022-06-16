package miutil

import (
	"log"
	"regexp"
	"time"
)

func OnlyNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
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
	if len(slice) == 0 {
		return make([][]T, 0)
	}
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

func NilDateTime(d *time.Time) time.Time {
	if d == nil {
		return time.Time{}
	}
	return *d
}
