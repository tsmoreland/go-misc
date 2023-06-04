package business

import "fmt"

var (
	kinetecoPrint string = "Kineteco Deal:"
)

// Solar handles all the different energy offers powered by solar.
type Solar struct {
	Name  string
	Netto float64
}

// Wind handles all the different energy offers powered by wind.
type Wind struct {
	Name  string
	Netto float64
}

type Energy interface {
	Solar | Wind // type-set which allows either Solar or Wind to be used in generics
	Cost() float64
}

// Complex is another type-set interface but the use of ~ allows automatic extensibility,
//
//	should additional complex types be added to Go in the future
type Complex interface {
	~complex64 | ~complex128
}

// Print prints the information for a solar product.
// The string is enriched with the required kineteco legal information.
func (s *Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *s)
}

// Print prints the information for a wind product.
// The string is enriched with the required kineteco legal information.
func (w *Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *w)
}

// PrintGeneric returns any type as string.
// The string is enriched with the required Kineteco legal information.
func PrintGeneric[T any](t T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, t)
}

// PrintSlice prints all Energy items in t to stdout, note the generic is
//
//	using T Energy rather than any for a more specific or
//	constrained generic
func PrintSlice[T Energy](t []T) {
	typeOfT := fmt.Sprintf("%T", t)
	for idx, itm := range t {
		fmt.Printf("%d (%v): %v", idx, typeOfT, PrintGeneric(itm))
	}
}

// PrintSlice2 uses a second constraint to replace []T as a slice of all items
//
//	that could be T (the ~ operator).
//	What this means is if we had say type myString string and string then
//	the constraint ~string would also match myString because it approximates
//	string
func PrintSlice2[T Energy, S ~[]T](t S) {
	typeOfT := fmt.Sprintf("%T", t)
	for idx, itm := range t {
		fmt.Printf("%d (%v): %v", idx, typeOfT, PrintGeneric(itm))
	}
}
