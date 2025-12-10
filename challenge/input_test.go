package challenge

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRaw(t *testing.T) {
	input := `a
b
c`

	require.Equal(t, input, Raw(strings.NewReader(input)))
}

func TestLines(t *testing.T) {
	t.Run("without trailing", func(t *testing.T) {
		input := `a
b
c`

		got := slices.Collect(Lines(strings.NewReader(input)))
		require.EqualValues(t, []string{"a", "b", "c"}, got)
	})

	t.Run("with trailing", func(t *testing.T) {
		input := `a
b
c
`

		got := slices.Collect(Lines(strings.NewReader(input)))
		require.EqualValues(t, []string{"a", "b", "c"}, got)
	})
}

func TestSectionsOf(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  []string
	}{
		{"1aa2aa3", []string{"1", "2", "3"}},
		{"1aa2aa3aa", []string{"1", "2", "3"}},
		{"aa1aa2aa3", []string{"1", "2", "3"}},
		{"aa1aa2aa3aa", []string{"1", "2", "3"}},
		{"aa1aaaa2aa3aa", []string{"1", "2", "3"}},
		{"aa1aaaaaa2aaaa3aa", []string{"1", "2", "3"}},
	} {
		t.Run(tt.input, func(t *testing.T) {
			require.EqualValues(t, tt.want, slices.Collect(SectionsOf(strings.NewReader(tt.input), "aa")))
		})
	}
}

func TestSections(t *testing.T) {
	require.EqualValues(
		t,
		[]string{"a\nb\nc", "1\n2\n3", "d\ne\nf"},
		slices.Collect(
			Sections(
				strings.NewReader(`a
b
c

1
2
3


d
e
f`),
			),
		),
	)
}

func TestInts(t *testing.T) {
	require.EqualValues(
		t,
		[]int{1, 2, 3},
		slices.Collect(
			Ints(
				strings.NewReader(`1
2
3`),
			),
		),
	)
}
