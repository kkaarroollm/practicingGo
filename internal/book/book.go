package book

import (
	"fmt"
	"time"
)

type Book struct {
	ID        int
	Title     string
	Author    string
	Published bool
	Date      time.Time
}

func (b *Book) String() string {
	return fmt.Sprintf("%d: %s by %s (Published: %v)", b.ID, b.Title, b.Author, b.Published)
}
