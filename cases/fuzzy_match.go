package cases

import (
	"bytes"
	"strings"
)

// fuzzyMatch reads sentences from channel c, applies fuzzy matching against
// entity, and writes matching sentences into buf. When done, it signals
// completion by closing the done channel.
func (study study[samples, aggregated]) fuzzyMatch(
	c chan string,
	entity []string,
	buf *bytes.Buffer,
	done chan struct{},
) {
	go func() {
		defer close(done)

		for s := range c {
			// Size
			if len(s) < 3 {
				continue
			}
			// Filters
			switch {
			case strings.Contains(s, "http"), strings.Contains(s, ".com"), strings.Contains(s, "www"):
				continue
			default:
			}

			// Entity
			for _, e := range entity {
				if !strings.Contains(s, e) {
					continue
				}

				// Concatenate sentence.
				buf.WriteString(s)
				break
			}
		}
	}()
}
