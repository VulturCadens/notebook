package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p><b>There was %d (ExecStart=) arguments.</b></p>", len(os.Args))

		w.Write([]byte("<ul>"))

		for i, arg := range os.Args {
			fmt.Fprintf(w, "<li>%d. argument: %s</li>", i, arg)
		}

		w.Write([]byte("</ul>"))
	})

	http.ListenAndServe("127.0.0.1:8000", nil)
}
