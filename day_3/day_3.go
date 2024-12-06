package day_3

import (
	"regexp"
	"strconv"
	"strings"
)

func mul(x, y int) int {
	return x * y
}

func List_map[I any, O any](l []I, mapper func(I) O) []O {
	mapped := make([]O, len(l))

	for i, element := range l {
		mapped[i] = mapper(element)
	}

	return mapped
}

func extract_valid_operations(s string) []string {
	r := regexp.MustCompile(`(?P<operation>[a-zA-Z_][a-zA-Z0-9_\']*)\((?P<args>([1-9][0-9]{0,2})(,[1-9][0-9]{0,2})*)?\)`)

	operations := r.FindAll([]byte(s), -1)

	if operations == nil {
		return []string{}
	}

	return List_map(operations, func(o []byte) string { return string(o) })
}

func interpret_operation(s string) (string, int, []int) {
	r := regexp.MustCompile(`(?P<operation>[a-zA-Z_][a-zA-Z0-9_\']*)\((?P<args>([1-9][0-9]{0,2})(,[1-9][0-9]{0,2})*)?\)`)

	matchs := r.FindStringSubmatch(s)
	result := make(map[string]string)

	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matchs[i]
		}
	}

	argc := 0
	argv := []int{}

	if args, is_in_result := result["args"]; is_in_result {
		argc = strings.Count(args, ",") + 1
		argv = List_map(strings.Split(args, ","), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
	}

	return result["operation"], argc, argv
}

func Evaluate_sequence(sequence string) int {
	result := 0

	enable_evaluation := true
	operations := extract_valid_operations(sequence)
	for _, op := range operations {
		op_name, argc, argv := interpret_operation(op)

		if op_name == "do" {
			enable_evaluation = true
		} else if op_name == "don't" {
			enable_evaluation = false
		} else if enable_evaluation && op_name == "mul" && argc == 2 {
			result += mul(argv[0], argv[1])
		}
	}

	return result
}
