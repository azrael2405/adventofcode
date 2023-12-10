package main

import (
	"day10/day10"
	"testing"
)

func Test_part1_1(t *testing.T){
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_10\test1.txt`)
	expected_answer := 4
	answer := day10.Parse_answer_one(test_data)
	if answer != expected_answer{
		t.Fatalf("Part1 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
		
	}
}

func Test_part1_2(t *testing.T){
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_10\test2.txt`)
	expected_answer := 8
	answer := day10.Parse_answer_one(test_data)
	if answer != expected_answer{
		t.Fatalf("Part1 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
		
	}
}


func Test_part2(t *testing.T){
	test_data := parse_input_from_file(`E:\projects\adventofcode2022\2023\day_10\test3.txt`)
	expected_answer := 4
	answer := day10.Parse_answer_two(test_data)
	if answer != expected_answer{
		t.Fatalf("Part2 Test does not conform to expected answer: %d / %d\n", answer, expected_answer)
		
	}
}