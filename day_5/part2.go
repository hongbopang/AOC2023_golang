package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type mapper struct {
	mapStart int
	mapEnd   int
	offset   int
}

type mapper_wrapper struct {
	mapFrom     string
	mapTo       string
	mappersList []mapper
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, 0)
		} else {
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}
			result = append(result, x)
		}
	}
	result = append(result, 0)
	return result, scanner.Err()
}

func Readstrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func makeStringSliceInts(strings []string) []int {
	var intSlice = []int{}
	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}
	return intSlice
}

func apply_map(input int, wrapper mapper_wrapper) int {

	for _, mapper := range wrapper.mappersList {
		if input <= mapper.mapEnd && input >= mapper.mapStart {
			return input + mapper.offset
		}
	}
	return input
}

func gauntlet(seeds []int, mapper map[string]mapper_wrapper) int {
	names := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	ans := 2616560939000
	ptr := 0
	for ptr < len(seeds) {
		fmt.Println(seeds[ptr])
		for seed_val := seeds[ptr]; seed_val < seeds[ptr]+seeds[ptr+1]; seed_val++ {
			curr_val := seed_val
			for _, name := range names {
				wrapper := mapper[name]
				curr_val = apply_map(curr_val, wrapper)
			}
			if curr_val < ans {
				ans = curr_val
			}
		}
		ptr += 2

	}
	return ans
}

func parse(lines []string) int {
	seeds := strings.Split(lines[0], " ")[1:]
	seeds_int := makeStringSliceInts(seeds)

	mappers := make(map[string]mapper_wrapper)
	mappers["seed-to-soil"] = mapper_wrapper{}
	mappers["soil-to-fertilizer"] = mapper_wrapper{}

	mappers["fertilizer-to-water"] = mapper_wrapper{}
	mappers["water-to-light"] = mapper_wrapper{}
	mappers["light-to-temperature"] = mapper_wrapper{}
	mappers["temperature-to-humidity"] = mapper_wrapper{}
	mappers["humidity-to-location"] = mapper_wrapper{}

	name := "seed-to-soil"

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		_, ok := mappers[strings.Split(line, " ")[0]]
		if ok {
			name = strings.Split(line, " ")[0]

			if entry, ok := mappers[name]; ok {
				names := strings.Split(name, "-")
				// Then we modify the copy
				entry.mapFrom = names[0]
				entry.mapTo = names[2]
				// Then we reassign map entry
				mappers["key"] = entry
			}

		} else {
			values := makeStringSliceInts(strings.Split(line, " "))
			start := values[1]
			end := values[1] + values[2] - 1
			offset := values[0] - values[1]

			tmp_mapper := mapper{start, end, offset}

			if entry, ok := mappers[name]; ok {
				entry.mappersList = append(entry.mappersList, tmp_mapper)
				mappers[name] = entry
			}
		}

	}

	return gauntlet(seeds_int, mappers)

}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("file open fail")
	}
	lines, err := Readstrings(readFile)
	fmt.Println(parse(lines))

}
