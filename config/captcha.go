package config

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Captcha() (string, string) {
	files, err := ioutil.ReadDir("./samples")
	var captcha_str []string

	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		captcha_str = append(captcha_str, f.Name())
	}
	rand.Seed(time.Now().UnixNano())
	//min := 0
	//max := 1069
	random_number := (rand.Intn(1070))
	result := strings.Split(captcha_str[random_number], ".png")[0]
	result = strings.Split(result, ".jpg")[0]
	path := string("./samples/" + result + ".png")

	return result, path
}
