package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	var matchRe = regexp.MustCompile(`<div class="m-btn (purple|pink)" data-v-8b1eac0c>([^<]+)</div>`)
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s",contents)
	ms := matchRe.FindAllSubmatch(contents, -1)
	fmt.Printf("%s\n", ms)
	fmt.Printf("%d\n", len(ms))
	for k, v := range ms {
		fmt.Printf("%d:%s\n", k, v[2])
	}
	//fmt.Printf("%s\n",ms[2][1])
}
