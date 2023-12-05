package day5

import (
	"fmt"
	"helper"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parse_seeds(seeds_line string) []int{
	seeds := []int{}
	for _, seed := range strings.Split(strings.TrimSpace(strings.Split(seeds_line, ":")[1]), " "){
		seed_int, conversion_error := strconv.Atoi(seed)
		helper.Check_error(conversion_error)
		seeds = append(seeds, seed_int)
	}
	return seeds
}


func Parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	seeds := parse_seeds(_data[0])
	answer_map := parse_data_to_almanac(_data[1:], seeds)
	locations := []int{}
	for key := range answer_map["location"]{
		locations = append(locations, key)
	}
	slices.Sort(locations)
	answer := locations[0]
	fmt.Println("Answer 1:", answer)
}