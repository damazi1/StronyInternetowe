package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var sg sync.WaitGroup
var mx sync.Mutex
var pierwsze []int = []int{}
var suma int = 0

func licz() {
	tmp := 0
	for range 1000 {
		tmp += 1
	}
	mx.Lock()
	suma += tmp
	mx.Unlock()
	sg.Done()
}

func getcountry(ip, api, field string, ch chan string) {
	m := map[string]any{}

	res, err := http.Get(api + ip)
	if err != nil {
		ch <- fmt.Sprintf("Błąd HTTP: %v", err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- fmt.Sprintf("Błąd HTTP: %v", err)
		return
	}
	value, _ := m[field]
	if value != nil {
		ch <- fmt.Sprintf("Brak lub pusta wartość dla pola '%s'", field)
		return
	}
	json.Unmarshal([]byte(body), &m)
	ch <- m[field].(string) + api
}

func sprawdz(liczba int) { // sprawdzenie czy liczba jest pierwsza
	for i := 2; i < liczba; i++ {
		// liczba ma jakiś dzielnik różny od 1 i siebie samej
		if liczba%i == 0 {
			goto end
		}
	}
	// liczba nie ma innego dzielnika różnego od 1 i siebie samej
	mx.Lock()
	pierwsze = append(pierwsze, liczba)
	mx.Unlock()
end:
	sg.Done()
}
func main() {
	liczb := 50000
	sg.Add(liczb)
	start := time.Now()
	for i := 0; i < liczb; i++ { // test kolejnych liczb
		go sprawdz(i)
	}
	sg.Wait()
	czas := time.Since(start)
	fmt.Println("Pierwsze:", pierwsze)
	fmt.Println("Czas obliczeń:", czas)

	sg.Add(10)
	for range 10 {
		go licz()
	}
	sg.Wait()
	fmt.Println("Obliczona suma:", suma)

	ip := "212.192.10.10000"
	ch := make(chan string)
	go getcountry(ip, "http://ip-api.com/json/", "country", ch)
	go getcountry(ip, "https://freeipapi.com/api/json/", "countryName", ch)
	go getcountry(ip, "http://www.geoplugin.net/json.gp?ip=", "geoplugin_countryName", ch)
	select {
	case s := <-ch:
		fmt.Println("Pierwszy naleziony wynik =", s)
	}
}
