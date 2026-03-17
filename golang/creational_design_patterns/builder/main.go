package builder

import "fmt"

func main() {
	req, err := GetNewHttpRequestBuilder("GET", "/api/outback").SetTimeout(3000).Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(req)
}
