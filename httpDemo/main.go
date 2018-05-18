package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/* Old Method
	//Initilaied this way because most 'read' functions
	//are not set up to handle resizing the slice if
	//it isn't large enough.
	bs := make([]byte, 99999)

	//Has returns, but reads directly into bs
	resp.Body.Read(bs)

	fmt.Println(string(bs))
	*/

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

//Be careful to actually do the things an interface
//is supposed to do. Go will let you fulfill an interface
//without fulfilling its logical requirements.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %v bytes\n", len(bs))
	return len(bs), nil
}
