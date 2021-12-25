package ds_string

type Pattern interface {
	Match(str string) int
	MatchAll(str string) []int
}
