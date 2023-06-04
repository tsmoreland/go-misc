package main

import (
	"fmt"
	"github.com/tsmoreland/go-misc/generics/business"
)

// main is our simple "playground" for the course.
// Note, that in production code, it is a good practice to keep the main function short.
func main() {
	// Create three different energy offers of kineteco
	solar2k := business.Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := business.Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := business.Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Printf(solar2k.Print())
	fmt.Printf(solar3k.Print())
	fmt.Printf(windwest.Print())
	fmt.Printf("our first generic function: %s\n", business.PrintGeneric[business.Solar](solar2k))
	fmt.Printf("our first generic function with wind: %s\n", business.PrintGeneric[business.Wind](windwest))

	fmt.Println("Challenge Chapter 1:")
	ss := []business.Solar{solar2k, solar3k}
	business.PrintSlice[business.Solar](ss) // generic type [business.Solar] is optional as it can be inferred from ss
	business.PrintSlice[business.Wind]([]business.Wind{windwest, windwest})

	cost := business.Cost(10, solar2k.Netto)
	fmt.Printf("solar2k cost for 10 uses: %v", cost)

	fmt.Println("Chapter 3")
	// cost := business.Cost(0.45, 10) -- compiler error because T can't be inferred float64 != int
	cost = business.Cost(0.45, float64(10)) // work-around #1 cast to the intended type, so it can be inferred
	cost = business.Cost[float64](0.45, 10) // work-around #2 be explicit about generic type
	fmt.Printf("cost for 0.45 uses at 10 per unit: %v", cost)

	business.PrintSlice2(ss)
}
