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
	Fragments []string
	Args []interface{}
}

func (s * SQLFragments) Len() (int) {
	return len(s.Fragments)
}

func (s * SQLFragments) Append(frag string) {
	s.Fragments = append( s.Fragments, frag )
}

func (s * SQLFragments) AppendQuery(frag string, args ...interface{}) {
	// replace "?" to $%d
	cnt := strings.Count(frag, "?")
	frag = strings.Replace(frag, "?", "$%d", -1)

	var varStartFrom = len(s.Args) + 1
	var varNumbers = createIntSliceWithRange(varStartFrom, varStartFrom + cnt)

	s.Fragments = append( s.Fragments, fmt.Sprintf(frag, varNumbers... ) )
	for _, a := range args {
		s.Args = append(s.Args, a)
	}
}

func (s * SQLFragments) Join(sep string) (string) {
	return strings.Join(s.Fragments, sep)
}

func (s * SQLFragments) String() (string) {
	return strings.Join(s.Fragments, " ")
}

func New() (*SQLFragments) {
	return new(SQLFragments)
}

