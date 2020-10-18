package main

import (
	"fmt"
	"net/url"
)

func main() {

	//s := "postgres://user:pass@host.com:5432/path?k=v#f"
	s := "?k=20&l=30"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// fmt.Println(u.Scheme)

	// fmt.Println(u.User)
	// fmt.Println(u.User.Username())
	// p, _ := u.User.Password()
	// fmt.Println(p)

	// fmt.Println(u.Host)
	// host, port, _ := net.SplitHostPort(u.Host)
	// fmt.Println(host)
	// fmt.Println(port)

	// fmt.Println(u.Path)
	// fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
	fmt.Println(m["l"][0])
}
