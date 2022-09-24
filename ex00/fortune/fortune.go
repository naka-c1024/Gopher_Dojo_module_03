package fortune

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type result struct {
	Fortune string `json:"fortune"`
	Wish    string `json:"wish"`
	Health  string `json:"health"`
}

var TimeNow = time.Now()

func IsOshougatsu() bool {
	t := TimeNow
	if t.Month() == time.Month(1) {
		if today := t.Day(); 1 <= today && today <= 3 {
			return true
		}
	}
	return false
}

func Omikuji() int {
	if IsOshougatsu() {
		return 0
	}
	return rand.Intn(7)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	rdmNum := Omikuji()
	jsonDate := result{}
	jsonDate.Fortune = RtnFortune(rdmNum)
	jsonDate.Wish = RtnWish(rdmNum)
	jsonDate.Health = RtnHealth(rdmNum)
	if err := json.NewEncoder(w).Encode(jsonDate); err != nil {
		log.Println(err)
	}
}

func Run() {
	flag.Parse()

	if flag.Arg(0) == "" {
		log.Fatalln("error: empty argument")
	} else if flag.Arg(1) != "" {
		log.Fatalln("error: multiple arguments")
	}
	http.HandleFunc("/", Handler)
	rand.Seed(time.Now().UnixNano())
	if err := http.ListenAndServe(":"+flag.Arg(0), nil); err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
