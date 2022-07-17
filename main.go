package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

//var flagFile = flag.String("file", "", "file used for input")

func main() {
	flag.Parse()
	printMD5(os.Stdin, os.Stdout)

}

func printMD5(r io.Reader, w io.Writer) {
	h := md5.New()

	_, err := io.Copy(h, r)
	if err != nil {
		fmt.Println("error copying md5: ", err)
		os.Exit(1)

	}
	fmt.Fprintf(w, "%x", h.Sum(nil))
}
