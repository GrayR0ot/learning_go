package task

import (
	"fmt"
	"imaging/filter"
	"imaging/my_image"
	"sync"
	"time"
)

func ProcessImages(input string, output string) {
	images := LoadFiles(input)
	fmt.Printf("Loaded %v my_image(s) from your input folder\n", len(images))

	images = channelProcess(images)
	images = waitGroupProcess(images)

	SaveFiles(images, output)
	fmt.Printf("Saved %v my_image(s) to your output folder\n", len(images))
}

func channelProcess(images []my_image.MyImage) []my_image.MyImage {

	poolSize := 2
	workerPool := make(chan struct{}, poolSize)

	results := make(chan my_image.MyImage, len(images))

	startTime := time.Now()

	for _, img := range images {
		workerPool <- struct{}{}
		go func(img my_image.MyImage) {
			defer func() {
				<-workerPool // Release the worker slot
			}()
			results <- filter.Filter(&img)
		}(img)
	}

	filteredImages := make([]my_image.MyImage, 0)
	for range images {
		filtered := <-results
		filteredImages = append(filteredImages, filtered)
	}
	close(results)

	endTime := time.Now()
	took := endTime.Sub(startTime)
	fmt.Println("[Channel] Time to process:", took)
	return filteredImages
}

func waitGroupProcess(images []my_image.MyImage) []my_image.MyImage {

	var wg sync.WaitGroup
	results := make(chan my_image.MyImage, len(images))

	startTime := time.Now()

	for _, img := range images {
		wg.Add(1)
		go func(img my_image.MyImage) {
			defer wg.Done()
			results <- filter.Filter(&img)
		}(img)
	}

	wg.Wait()
	close(results)

	filteredImages := make([]my_image.MyImage, 0)
	for img := range results {
		filteredImages = append(filteredImages, img)
	}

	endTime := time.Now()
	took := endTime.Sub(startTime)
	fmt.Println("[WaitGroup] Time to process:", took)
	return filteredImages
}
