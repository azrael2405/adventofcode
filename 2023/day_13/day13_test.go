package main

import (
	"day13/day13"
	"testing"
)

func Test_part1(t *testing.T) {
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_13\test1.txt`)
	expected_answer := 405
	answer := day13.Parse_answer_one(test_data)
	if answer != expected_answer {
		t.Fatalf("Part1 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)

	}
}

func Test_part2(t *testing.T) {
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_13\test1.txt`)
	expected_answer := 400
	answer := day13.Parse_answer_two(test_data)
	if answer != expected_answer {
		t.Fatalf("Part2 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)

	}
}
