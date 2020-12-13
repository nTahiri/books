// Book list is base on https://raw.githubusercontent.com/benoitvallon/100-best-books/master/books.json
package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)
const DAT_URL = "https://raw.githubusercontent.com/nTahiri/books/master/data/books.json"
type DataSource interface {
	GetBooksFromDataSource() []*Book
}
type bookDataSource struct {
	books []*Book
	url string
}

var once sync.Once
var instance DataSource

func readDataFromFile(db *bookDataSource) (result []*Book, err error) {
	res, err := http.Get(db.Url())
	if err != nil {
		                                                                                   
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(message, &result)
	return result, err
}
func (b *bookDataSource) SetUrl(url string){b.url = url}
func (b *bookDataSource) Url() string{return b.url}
func (b *bookDataSource) GetBooksFromDataSource() []*Book {
	return b.books
}
func LoadBooks() DataSource {
	once.Do(func() {
		db := bookDataSource{}
		db.SetUrl(DAT_URL)
		books, err := readDataFromFile(&db)
		if err == nil {
			db.books = books
		}
		instance = &db
	})
	return instance
}