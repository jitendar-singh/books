package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/redhat-developer/books"
)

var _ = Describe("Book", func() {
	var (
		longBook Book
		shortBook Book
	)
	BeforeEach(func() {
		longBook = Book{
			Title: "Adventures of Tom Sawyer",
			Author: "Mark Twain",
			Pages: 346,
		}

		shortBook = Book{Title: "Animal Farm",Author: "George Orwell",Pages: 275}
	})
	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("Should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})
		Context("With less than 300 pages", func() {
			It("Should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
