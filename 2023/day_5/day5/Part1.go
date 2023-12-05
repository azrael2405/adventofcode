package day5

import (
	"fmt"
	"helper"
	"os"
	"slices"
	"time"
)


func Parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	answer_map := parse_data_to_almanac(_data)
	file, _ := os.Create("almanac.out")
	defer file.Close()
	for key, values := range answer_map{
		file.WriteString(fmt.Sprintln(key))
		for int_key, value := range values{
			file.WriteString(fmt.Sprintln(int_key, value))
		}
	}
	locations := []int{}
	for key := range answer_map["location"]{
		locations = append(locations, key)
	}
	slices.Sort(locations)
	answer := locations[0]
	fmt.Println("Answer 1:", answer)
}