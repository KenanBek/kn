package main

import (
	"kn/se/internal/app/crawler"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	log.Println("Application: kn-se-crawler")

	runCount := 1
	runInterval := 5 // seconds
	runCounter := 0

	rc, foundRC := syscall.Getenv("KN_BE_SE_RUN_COUNT")
	if foundRC {
		convRC, err := strconv.Atoi(rc)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if convRC >= 0 {
			runCount = convRC
		} else {
			log.Println("env variable KN_BE_SE_RUN_COUNT ignored, must be greater than -1")
		}
	}

	ri, foundRI := syscall.Getenv("KN_BE_SE_RUN_INTERVAL")
	if foundRI {
		convRI, err := strconv.Atoi(ri)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if convRI > 0 {
			runInterval = convRI
		} else {
			log.Println("env variable KN_BE_SE_RUN_INTERVAL ignored, must be greater than 0")
		}
	}

	ticker := time.NewTicker(time.Duration(runInterval) * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		for {
			if runCount == 0 {
				// do not stop
				return
			}

			if runCounter >= runCount {
				done <- true
			} else {
				time.Sleep(time.Duration(runInterval) * time.Second)
			}
		}
	}()

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-done:
			log.Println("\r- run count reached")
			os.Exit(0)
		case <-interrupt:
			log.Println("\r- program interrupted")
			os.Exit(0)
		case t := <-ticker.C:
			log.Println("Running Crawler.Crawl at", t)

			go func() {
				sl := crawler.NewJSONSourceLoader("assets/initial_sources.json")
				s := crawler.NewCollyScraper()
				r := crawler.NewMongoRepository()
				c := crawler.NewWebCrawler(sl, s, r)
				err := c.Crawl()

				if err != nil {
					log.Println(err)
				}
			}()

			runCounter++
		}
	}
}
