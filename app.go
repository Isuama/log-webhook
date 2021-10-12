package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"github.com/oklog/ulid"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	writeToFile(genUlid(), string(body))
}

func genUlid() string {
	//t := time.Now().UTC()
	//log.Println(t.UnixNano())
	//entropy := rand.New(rand.NewSource(t.UnixNano()))
	//id := ulid.MustNew(ulid.Timestamp(t), entropy)
	//id := t.UnixNano().String()
	//log.Println("github.com/oklog/ulid:          %s\n", id.String())
	//return id.String()
	return "log-hunter"
}

func writeToFile(name string, content string) {

	file, err := os.Create("/usr/share/hunter-alert/" + name + ".log")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(content)
	log.Println(content)
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

