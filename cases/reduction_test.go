package cases

import (
	"testing"
)

func Test_reduce(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		text, entity string
		want         string
	}{
		{
			name:   "basic case",
			text:   "This is a sentence about Gemini. This is another sentence. And a third one about Gemini Code Assist.",
			entity: "Gemini",
			want:   "This is a sentence about Gemini. And a third one about Gemini Code Assist.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, _ := reduce(tc.text, tc.entity)
			if got != tc.want {
				t.Errorf("reduce() = %s, want %s", got, tc.want)
			}
		})
	}
}
