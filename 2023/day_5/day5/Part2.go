package day5

import (
	"fmt"
	"helper"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parse_seeds_multiple(seeds_line string) [][]int{
	seeds := [][]int{}
	re := regexp.MustCompile(`\d+ \d+`)
	for _, seed_string := range re.FindAllString(seeds_line, -1){
		seed_split := strings.Split(strings.TrimSpace(seed_string), " ")
		seed_start, conversion_error := strconv.Atoi(seed_split[0])
		helper.Check_error(conversion_error)
		seed_range, conversion_error := strconv.Atoi(seed_split[1])
		helper.Check_error(conversion_error)
		seeds = append(seeds, []int{seed_start, seed_start + seed_range})
	}
	return seeds
}


func Parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 2")
	seeds := parse_seeds_multiple(_data[0])
	answer_map := parse_data_to_almanac_with_seed_ranges(_data[1:], seeds)
	locations := []int{}
	for _, value := range answer_map["location"]{
		locations = append(locations, value[0])
	}
	slices.Sort(locations)
	answer := locations[0]
	fmt.Println("Answer 2:", answer)
}