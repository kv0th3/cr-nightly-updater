package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"os/exec"
)

func main() {
	res, err := http.Get("https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux/LAST_CHANGE")
	if err != nil {
		log.Fatal(err)
	}
	last_change, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latest Chromium snapshot build number is %s\n", last_change)
	url := "https://commondatastorage.googleapis.com/chromium-browser-snapshots/Linux/" + string(last_change) + "/chrome-linux.zip"

	fmt.Printf("url: %s\n", url)

	res_download, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res_download.Body)

	if err != nil {
		log.Fatal(err)
	}

	res_download.Body.Close()

	ioutil.WriteFile("chrome-linux.zip", data, 0666)

	fmt.Printf("Downloaded the latest Chromium build %s\n", last_change)

	cmd := exec.Command("unzip", "-o", "chrome-linux.zip")
	err3 := cmd.Run()
	if err3 != nil {
		log.Fatal(err3)
	}
	println("Unzipped chrome-linux.zip")
}
