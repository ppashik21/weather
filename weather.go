package main

const (
	basic_url = "http://api.weatherapi.com/v1/current.json?key=9e526769b1e54fd087d163307210106"
)

type request_data struct {
	city string
	air  bool
}

func create_url(x request_data) string {
	var url string
	if x.air == true {
		url = basic_url + "&q=" + x.city + "&aqi=yes"
	} else {
		url = basic_url + "&q=" + x.city + "&aqi=no"
	}

	// new_url := basic_url + x.city + x.air
	return url

}

func main() {
	var r request_data
	r.city = "Minsk"
	r.air = true
	q := create_url(r)
	println(q)
}
