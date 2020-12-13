package book

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetBooks(t *testing.T) {

	loadedResult := LoadBooks().GetBooksFromDataSource()
	defaultApiService := &DefaultApiService{}
	books, _ := defaultApiService.GetBooks()
	if len(books.([]*Book)) != len(loadedResult) {
		t.Errorf("Loaded Data and direct data are diffeerent " +
			"loadedResult lenght: %d and directResult lenght: %d", len(loadedResult), len(books.([]*Book)))
	}

}
func TestSearchBooks(t *testing.T) {
	defaultApiService := &DefaultApiService{}
	for _, searchPayload := range readSearchCriteriaFile() {
		books, _ := defaultApiService.SearchBooks(searchPayload)
		if len(books.([]*Book)) == 0 {
			t.Errorf("Books not Found for %+v\n",searchPayload)
		}
	}
}
func TestSearchUnexistingBooks(t *testing.T) {
	criteria := `{"author": "Victor Hugo"}`
	payload := SearchPayload{}
	json.Unmarshal([]byte(criteria), &payload)
	defaultApiService := &DefaultApiService{}
	books, _ := defaultApiService.SearchBooks(payload)
	if len(books.([]*Book)) != 0 {
		t.Errorf("Books not Found for %+v\n",payload)
	}
}


func readSearchCriteriaFile() []SearchPayload{
	jsonFile, _ := os.Open("./data/search-books-test.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := []SearchPayload{}
	json.Unmarshal(byteValue, &result)
	return result
}