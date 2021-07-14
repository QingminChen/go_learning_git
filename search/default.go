package search

import "fmt"

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	fmt.Println("Qingmin+++++++++++++++++++++++++++++++++++++++++++search default init 1")
	Register("default", matcher)
}

func init(){
	fmt.Println("Qingmin+++++++++++++++++++++++++++++++++++++++++++search default init 2")
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
