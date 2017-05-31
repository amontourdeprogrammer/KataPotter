package KataPotter
import "testing"

func TestNoBookReturns0(t *testing.T){
    books := [5]int{0,0,0,0,0}
    if BookSale(books) != 0.0 {
        t.Error("Expected 0")
    }
}

func TestOneHP1Returns8(t *testing.T) {
    books := [5]int{1,0,0,0,0}
    if BookSale(books) != 8 {
        t.Error("Expected 8")
    }
}

