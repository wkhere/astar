package geo

import "fmt"

func ExampleH() {
	fmt.Println(
		H(
			Pt{AbsDeg(50, 03, 59), AbsDeg(05, 42, 53)},
			Pt{AbsDeg(58, 38, 39), AbsDeg(03, 04, 12)},
		))

	fmt.Println(H(Pt{20.42, 3.5}, Pt{23.5, 111.111}))
	// Output:
	// 968
	// 10779
}
