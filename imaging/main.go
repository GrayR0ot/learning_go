package main

import "imaging/task"

func main() {
	const inputFolder string = "input"
	const outputFolder string = "output"

	task.ProcessImages(inputFolder, outputFolder)
}
