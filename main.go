package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello from Go! Pod: %s\n", hostname)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	fmt.Println("Server running on :8087")
	http.ListenAndServe(":8087", nil)
}

/*```

**`go.mod`**
```
module github.com/<your-username>/go-app

go 1.22*/
