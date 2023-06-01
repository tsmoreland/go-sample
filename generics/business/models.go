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

func PrintSlice[T any](t []T) {
	for idx, itm := range t {
		fmt.Printf("%d - %v", idx, PrintGeneric(itm))
	}
}
