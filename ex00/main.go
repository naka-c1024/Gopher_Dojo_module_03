package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"
	// "omikuji/fortune"
)

type Result struct {
	Fortune string `json:"fortune"`
	Wish    string `json:"wish"`
	Health  string `json:"health"`
}

func isOshougatsu() bool {
	t := time.Now()
	// t := time.Date(2001, 1, 3, 23, 59, 59, 0, time.Local) // debug用
	if t.Month() == time.Month(1) {
		if today := t.Day(); 1 <= today && today <= 3 {
			return true
		}
	}
	return false
}

func omikuji() int {
	if isOshougatsu() {
		return 0
	}
	return rand.Intn(7)
}

func rtnHealth(rdm_num int) string {
	switch rdm_num {
	case 0:
		return "Very good. It will improve with exercise."
	case 1, 2:
		return "Good."
	case 3:
		return "No problems."
	case 4:
		return "No problems. However, do not binge eat or drink."
	case 5:
		return "Try to make sure your body does not become cold."
	case 6:
		return "Your worries will continue. Sometimes, you just need to rest your mind."
	default:
		return "??"
	}
}

func rtnWish(rdm_num int) string {
	switch rdm_num {
	case 0:
		return "It will definitely come true."
	case 1:
		return "It shall come true without a hitch."
	case 2:
		return "Teamwork is the key to success. If you work together, it shall come true."
	case 3:
		return "If you take action, it will come true."
	case 4:
		return "It will take time, but it will happen."
	case 5:
		return "If you desire it, it will come true. However, this will come with much difficulty."
	case 6:
		return "It will be difficult for it to come true."
	default:
		return "??"
	}
}

func rtnFortune(rdm_num int) string {
	switch rdm_num {
	case 0:
		return "Dai-kichi"
	case 1:
		return "Kichi"
	case 2:
		return "Chuu-kichi"
	case 3:
		return "Sho-kichi"
	case 4:
		return "Sue-kichi"
	case 5:
		return "Kyo"
	case 6:
		return "Dai-kyo"
	default:
		return "??"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	rdm_num := omikuji()
	jsonDate := Result{}
	jsonDate.Fortune = rtnFortune(rdm_num)
	jsonDate.Wish = rtnWish(rdm_num)
	jsonDate.Health = rtnHealth(rdm_num)
	err := json.NewEncoder(w).Encode(jsonDate)
	if err != nil {
		log.Println(err)
	}
}

func Run() {
	flag.Parse()

	if flag.Arg(0) == "" {
		log.Fatalln("error: empty argument")
	}
	http.HandleFunc("/", handler)
	rand.Seed(time.Now().UnixNano())
	if err := http.ListenAndServe(":"+flag.Arg(0), nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	// flag.Parse()

	// if flag.Arg(0) == "" {
	// 	log.Fatalln("error: empty argument")
	// }
	// http.HandleFunc("/", handler)
	// rand.Seed(time.Now().UnixNano())
	// if err := http.ListenAndServe(":"+flag.Arg(0), nil); err != nil {
	// 	log.Fatalln("ListenAndServe: ", err)
	// }
	Run()
}
