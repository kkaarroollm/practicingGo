package main

import (
	"fmt"
	b "github.com/kkaarroollm/practicingGo/internal/book"
	"time"
)

func main() {

	book1 := b.Book{
		ID:        1,
		Title:     "W dupie piszczy",
		Author:    "Ania co ma kota",
		Published: false,
		Date:      time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	book2 := b.Book{
		ID:        2,
		Title:     "Alice in Wonderland",
		Author:    "Potter Carroll",
		Published: true,
		Date:      time.Date(1988, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	books := b.Books{&book1, &book2}
	books.Add(&b.Book{
		ID:        3,
		Title:     "The Alchemist",
		Author:    "Lola Coelho",
		Published: true,
		Date:      time.Date(1865, time.November, 26, 0, 0, 0, 0, time.UTC),
	})

	fmt.Println("\nBooks:")
	for _, book := range books {
		fmt.Println(book)
	}

	fmt.Println("\nHard Sort by Title:")
	if _, err := books.Sort(b.Title, true); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, book := range books {
			fmt.Println(book)
		}
	}

	fmt.Println("\nSoft Sort by Author:")
	if copied, err := books.Sort(b.Author, false); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, book := range copied {
			fmt.Println(book)
		}
	}

	fmt.Println("\nBooks after Soft Sort (unchanged):")
	for _, book := range books {
		fmt.Println(book)
	}

	fmt.Println("\nFilter by Published: (true)")
	if filtered, err := books.Filter(b.Published, "true"); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, book := range filtered {
			fmt.Println(book)
		}
	}

	fmt.Println("\nPublishing book 1:")
	if err := book1.Publish(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(book1.String())
	}

	fmt.Println("\nPublishing book 2:")
	if err := book2.Publish(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(book2)
	}

	fmt.Println("\nFilter by Published: (true)")
	if filtered, err := books.Filter(b.Published, "true"); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, book := range filtered {
			fmt.Println(book)
		}
	}
}
