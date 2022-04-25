package util

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
