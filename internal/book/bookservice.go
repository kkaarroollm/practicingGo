package book

import (
	"fmt"
	"github.com/kkaarroollm/practicingGo/pkg/utils"
	"reflect"
	"sort"
	"time"
)

type Books []*Book

func (b *Books) Add(book *Book) {
	*b = append(*b, book)
}

func (b *Book) Publish() error {
	if b.Published {
		return fmt.Errorf("book is already published")
	}
	b.Published = true
	b.Date = time.Now()
	return nil
}

func (b *Books) Sort(field BookField, hard bool) (Books, error) {
	getSliceToSort := func() Books {
		if hard {
			return *b
		}
		copied := make(Books, len(*b))
		copy(copied, *b)
		return copied
	}

	sorted := getSliceToSort()

	functions := map[BookField]func(i, j int) bool{
		ID:        func(i, j int) bool { return sorted[i].ID < sorted[j].ID },
		Title:     func(i, j int) bool { return sorted[i].Title < sorted[j].Title },
		Author:    func(i, j int) bool { return sorted[i].Author < sorted[j].Author },
		Date:      func(i, j int) bool { return sorted[i].Date.Before(sorted[j].Date) },
		Published: func(i, j int) bool { return sorted[i].Published && !sorted[j].Published },
	}

	sortFunc, exists := functions[field]
	if !exists {
		return nil, fmt.Errorf("invalid field: %v", field)
	}

	sort.Slice(sorted, sortFunc)

	if hard {
		*b = sorted
	}
	return sorted, nil
}

func (b *Books) Filter(field BookField, value string) ([]*Book, error) {
	var result []*Book

	for _, book := range *b {
		fieldValue := reflect.ValueOf(book).Elem().FieldByName(string(field))

		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("invalid field: %s", field)
		}

		parsedValue, err := utils.ParseValue(value, fieldValue.Kind())

		if err != nil {
			return nil, fmt.Errorf("failed to parse value: %v", err)
		}

		if utils.CompareValues(fieldValue.Interface(), parsedValue) {
			result = append(result, book)
		}
	}

	return result, nil
}

// / just for practice
func sortByPublished(b Books) Books {
	published := make(Books, 0)
	unpublished := make(Books, 0)

	for _, book := range b {
		if book.Published {
			published = append(published, book)
		} else {
			unpublished = append(unpublished, book)
		}
	}

	return append(published, unpublished...)
}
