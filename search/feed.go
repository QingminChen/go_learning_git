package search

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

//const dataFile = "C:\\Codes\\GO_PATH\\src\\go_learning\\data\\data.json"
const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	var wd, err = os.Getwd()
	//filepath.Dir(dataFile)
	//filepath.Abs(wd)
	var absPath, _ = filepath.Abs(wd+"\\"+dataFile)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}

func init() {
	fmt.Println("Qingmin+++++++++++++++++++++++++++++++++++++++++++search feed init 1")

}

func init(){
	fmt.Println("Qingmin+++++++++++++++++++++++++++++++++++++++++++search feed init 2")
}
