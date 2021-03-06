package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/otiai10/copy"
)

func main() {
	startTime := time.Now()

	cpFolders := map[string]string{
		"/data/sftp/1차_구축사업/unit5/07랜드마크_이미지_AI데이터":      "/data/ai_data_2020_01/07",
		"/data/sftp/1차_구축사업/unit5/15한국인_재식별_이미지":         "/data/ai_data_2020_01/15",
		"/data/sftp/1차_구축사업/unit5/19위성영상_객체판독":           "/data/ai_data_2020_01/19",
		"/data/sftp/1차_구축사업/unit1/10사람_인체자세_3D_AI데이터":    "/data/ai_data_2020_01/10",
		"/data/sftp/1차_구축사업/unit1/13피트니스_자세_이미지":         "/data/ai_data_2020_01/13",
		"/data/sftp/1차_구축사업/unit2/09한국인_대화음성_AI데이터":      "/data/ai_data_2020_01/09",
		"/data/sftp/1차_구축사업/unit2/18감성_대화_말뭉치":           "/data/ai_data_2020_01/18",
		"/data/sftp/1차_구축사업/unit3/12도로환경_파노라마_이미지_AI데이터": "/data/ai_data_2020_01/12",
		"/data/sftp/1차_구축사업/unit3/16도로주행영상":              "/data/ai_data_2020_01/16",
		"/data/sftp/1차_구축사업/unit5/03딥페이크_변조영상_AI데이터":     "/data/ai_data_2020_01/03",
	}
	for origin := range cpFolders {
		if _, err := os.Stat(origin); err != nil {
			// path/to/whatever exists
			log.Fatal("non exist path")
		} else {
			log.Println(origin, "exist!!")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(len(cpFolders))
	opt := copy.Options{
		AddPermission: os.FileMode(0777),
	}
	log.Printf("---------------\n\n starting copy .... \n\n---------------\n")
	for origin, target := range cpFolders {
		go func(o, t string) {
			err := copy.Copy(o, t, opt)
			if err != nil {
				log.Println(err)
			}
			log.Println(o, "complete!")
			wg.Done()
		}(origin, target)
	}
	wg.Wait()
	log.Println("ALL COMPLETE!!!!!!!")
	log.Println(time.Now().Sub(startTime))
}
