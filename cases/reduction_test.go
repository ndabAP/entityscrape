package cases

import (
	"bytes"
	"testing"
)

func Test_reduce(t *testing.T) {
	tests := []struct {
		name   string
		text   []byte
		entity string
		want   []byte
	}{
		{
			name:   "dot terminal",
			text:   []byte("Lorem ipsum is a dummy or placeholder text commonly used in graphic design, publishing, and web development."),
			entity: "Lorem ipsum",
			want:   []byte("Lorem ipsum is a dummy or placeholder text commonly used in graphic design, publishing, and web development."),
		},
		{
			name:   "exclamation mark terminal",
			text:   []byte("Hello, World!"),
			entity: "World",
			want:   []byte("Hello, World!"),
		},
		{
			name:   "no entity",
			text:   []byte("Hello, World!"),
			entity: "Triangle",
			want:   []byte(""),
		},
		{
			name:   "two sentences",
			text:   []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat."),
			entity: "Lorem ipsum",
			want:   []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua."),
		},
		{
			name:   "two sentences, entity in second",
			text:   []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat."),
			entity: "veniam",
			want:   []byte("Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat."),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := study[any, any]{}
			got, _ := s.reduce(tc.text, tc.entity)
			if !bytes.Equal(got, tc.want) {
				t.Errorf("reduce() = %s, want %s", got, tc.want)
			}
		})
	}
}
