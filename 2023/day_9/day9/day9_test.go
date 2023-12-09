package day9

import (
	"helper"
	"os"
	"strings"
	"testing"
)

func parse_input_from_file(filepath string) []string{
	file_data, file_error := os.ReadFile(filepath)
	helper.Check_error(file_error)
	lines := strings.Split(string(file_data), "\n")
	return lines
}
func Test_part1(t *testing.T){
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_9\test1.txt`)
	expected_answer := 6440
	answer := Parse_answer_one(test_data)
	if answer != expected_answer{
		t.Fatalf("Part1 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
		
	}
}

func Test_part2(t *testing.T){
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_9\test1.txt`)
	expected_answer := 5905
	answer := Parse_answer_two(test_data)
	if answer != expected_answer{
		t.Fatalf("Part2 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
		
	}
}