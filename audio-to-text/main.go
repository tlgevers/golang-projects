package main

import (
	"fmt"
	speech "github.com/tlgevers/golang-projects/audio-to-text/pkg/speech"
)

const (
	//file1 = "gs://athens-software-engineering-lectures/week1_lecture1.flac"
	//output1 = "./output/week1_lecture1.txt"
	file2 = "gs://athens-software-engineering-lectures/week1_lecture2.flac"
	output2 = "./output/week1_lecture2.txt"
	file3 = "gs://athens-software-engineering-lectures/week3_lecture1.flac"
	output3 = "./output/week3_lecture1.txt"
	file4 = "gs://athens-software-engineering-lectures/week3_lecture2_part1.flac"
	output4 = "./output/week3_lecture2_part1.txt"
	file5 = "gs://athens-software-engineering-lectures/week3_lecture2_part2.flac"
	output5 = "./output/week3_lecture2_part2.txt"
	file6 = "gs://athens-software-engineering-lectures/week5_lecture1.flac"
	output6 = "./output/week5_lecture1.txt"
	file7 = "gs://athens-software-engineering-lectures/week5_lecture2.flac"
	output7 = "./output/week5_lecture2.txt"
)

func main() {
	fmt.Println("Video to text")
	speech.Convert(file2, output2)
	speech.Convert(file3, output3)
	speech.Convert(file4, output4)
	speech.Convert(file5, output5)
	speech.Convert(file6, output6)
	speech.Convert(file7, output7)
}
