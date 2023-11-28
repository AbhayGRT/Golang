package main
import "fmt"
func main() {

	codes := map[string]string{"en": "English","fr": "French"}
	
	fmt.Println(codes)
	fmt.Println(len(codes))

	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])

	// Adding a key value pair
	codes["it"] = "Italian"
	fmt.Println(codes)

	// update key-value pair
	codes["en"] = "English Language"

	// delete key-value pair
	delete(codes, "en")

	// iterate
	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}
	
	// Empty
	code := make(map[string]int)
	fmt.Println(code)

	// found and value
	cod := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := cod["en"]
	fmt.Println(found, value)
	value, found = cod["hh"]
	fmt.Println(found, value)
}