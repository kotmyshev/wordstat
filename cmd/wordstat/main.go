package main

import (
	"fmt"

	"pkg/rndstr"
)

func main() {

	rndslovo := rndstr.GenerateRandomString(115, 10)

	/*
		http.HandleFunc("/", reqHandler)

		err := http.ListenAndServe("localhost:8080", nil)
		log.Fatal(err)

			req := httptest.NewRequest("GET", "http://localhost:8080///example.com/foo", nil)
			w := httptest.NewRecorder()
			reqHandler(w, req)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Header.Get("Content-Type"))
			fmt.Println(string(body))
	*/

	fmt.Println(rndslovo)
}
