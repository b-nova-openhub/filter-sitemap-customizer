package filter

type Matcher interface {
	Match(string, string, string) bool
}

type MatchApplier struct {
	Matcher Matcher
}

func (m *MatchApplier) Apply(s, prefix, pattern string) bool {
	return m.Matcher.Match(s, prefix, pattern)
}
