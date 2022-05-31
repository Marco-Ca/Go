package main

import "fmt"

func main() {
	// Loop through ids
	ids := []int{1, 2, 3, 4, 5}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("SUM: %d\n", sum)

	// Range with map
	emails := map[string]string{"Marco": "marco@ymail.com", "Paulo": "paulo@xmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for _, k := range emails {
		fmt.Printf("%s\n", k)
	}
}
