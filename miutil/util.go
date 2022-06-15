package miutil

import (
	"fmt"
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

func NewError(cause, err error) error {
	return fmt.Errorf("err: %w cause: %v", err, cause)
}

func NilString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func NotProvidedString(s *string) string {
	if s != nil {
		if *s == "" {
			return "not provided"
		}
		return *s
	}
	return "not provided"
}

func NilBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func NilInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

func NilFloat(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

func NilStringSlice(ss *[]string) []string {
	if ss != nil && len(*ss) > 0 {
		return *ss
	}
	return []string{}
}

func NilDateTime(d *time.Time) time.Time {
	if !d.IsZero() {
		return *d
	}
	return time.Time{}
}
