package runner

import "fmt"

func showBanner() {
	fmt.Printf(`
%s  v%s

by @%s
`, banner, version, author)
}
