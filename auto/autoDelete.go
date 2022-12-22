package auto

import (
	"fmt"
	"log"
	"os"
	"time"
)

// files(images) older than provided number of days will be deleted
func DeleteOld(numberOfDays int64) {
	for {
		fileInfo, err := os.ReadDir("./static/images")
		if err != nil {
			log.Fatal(err.Error())
		}
		now := time.Now()
		for _, info := range fileInfo {
			i, _ := info.Info()
			if diff := now.Sub(i.ModTime()); diff > time.Hour*24*time.Duration(numberOfDays) {
				log.Printf("Deleting %s which is %s old\n", info.Name(), diff)
				file := "./static/images/" + info.Name()
				fmt.Println(file)
				os.Remove(file)
			}
		}
		time.Sleep(time.Hour * 24)
	}
}
