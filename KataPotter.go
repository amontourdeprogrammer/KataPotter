package KataPotter

func BookSale(books [5]int) int {
    count := 0
    for i := 0; i < 5; i++ {
        count += books[i]
    }
    return count * 8
}
