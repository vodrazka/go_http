package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Pass port or ip and port.")
		return
	} else if len(os.Args) == 2 { // LISTEN
		port := os.Args[1]
		http.HandleFunc("/", indexHandler)
		log.Printf("Listening on port %s", port)
		log.Printf("Open http://localhost:%s in the browser", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else if len(os.Args) == 3 { //SEND
		ip := os.Args[1]
		port := os.Args[2]
		log.Printf("Will sending to %s:%s", ip, port)
		for {
			fmt.Println("Message or CTRL+C:")
			in := bufio.NewReader(os.Stdin)
			msg, _ := in.ReadString('\n')
			_, err := http.Post(fmt.Sprintf("http://%s:%s", ip, port), "application/json", bytes.NewBufferString(msg))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
