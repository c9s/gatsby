package sqlfragments
import "fmt"
import "strings"

func createIntSliceWithRange(from, to int) ([]interface{}) {
	var slice []interface{}
	for i := from ; i < to ; i++ {
		slice = append(slice, i)
	}
	return slice
}

type SQLFragments struct {
	fragments []string
	args []interface{}
}

func (s *SQLFragments) Len() (int) {
	return len(s.fragments)
}

func (s *SQLFragments) Append(frag string) {
	s.fragments = append( s.fragments, frag )
}

func (s *SQLFragments) AppendQuery(frag string, args ...interface{}) {
	// replace "?" to $%d
	cnt := strings.Count(frag, "?")
	frag = strings.Replace(frag, "?", "$%d", -1)

	var varStartFrom = len(s.args) + 1
	var varNumbers = createIntSliceWithRange(varStartFrom, varStartFrom + cnt)

	s.fragments = append( s.fragments, fmt.Sprintf(frag, varNumbers... ) )
	for _, a := range args {
		s.args = append(s.args, a)
	}
}

func (s *SQLFragments) Args() []interface{} {
	return s.args
}

func (s *SQLFragments) Join(sep string) (string) {
	return strings.Join(s.fragments, sep)
}

func (s *SQLFragments) String() (string) {
	return strings.Join(s.fragments, " ")
}

func (s *SQLFragments) Like(columnName string, value interface{}) (*SQLFragments) {
	s.AppendQuery(fmt.Sprintf("%s = ?", columnName), value)
	return s
}

func New() (*SQLFragments) {
	return new(SQLFragments)
}


