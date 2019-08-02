package cli

import (
	"fmt"

	"github.com/ndabAP/entityscrape/pkg/api"
)

// Do does
func Do() error {
	news, err := api.Get("Donald Trump")
	if err != nil {
		return err
	}

	for _, n := range news {
		assocentities, err := assoc(n.Text, []string{"Donald Trump", "Donald John Trump", "Donald J. Trump", "Donald T.", "D. Trump", "Trump"})
		if err != nil {
			return err
		}

		fmt.Println(assocentities)
	}

	return nil
}
