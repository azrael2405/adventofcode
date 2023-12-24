package main

import (
	"day18/day18"
	"testing"
)

func Test_part1(t *testing.T) {
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_18\test1.txt`)
	expected_answer := 6440
	answer := day18.Parse_answer_one(test_data)
	if answer != expected_answer {
		t.Fatalf("Part1 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
	}
}

func Test_part2(t *testing.T) {
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_18\test1.txt`)
	expected_answer := 5905
	answer := day18.Parse_answer_two(test_data)
	if answer != expected_answer {
		t.Fatalf("Part2 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
	}
}
