package main

import (
	"fmt"
)

// var python, java bool

type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number: %v", float64(e))
}

func main() {
	// hello()
	// fmt.Println(add(1, 2))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(split(17))
	// var i int
	// fmt.Println(i, python, java)
	// fmt.Println(Sqrt(2), math.Sqrt(2))
	// pic.Show(Pic)
	// wc.Test(WordCount)
	/* f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	} */
	/* hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	} */
	/* fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2)) */
	// reader.Validate(MyReader{})
	/* s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) */
	/* m := Image{}
	pic.ShowImage(m) */
	/* ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	} */
	// fmt.Println(Same(tree.New(1), tree.New(2)))

	Crawl("https://golang.org/", 4, fetcher)

}
