package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read the CSV file.
	file_system := os.DirFS("./")
	in, err := fs.ReadFile(file_system, "problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the bytes into string slices
	r := csv.NewReader(strings.NewReader(string(in)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Validate the records against the struct.
	quiz := []*QuizType{}
	for _, record := range records {
		a, err := strconv.ParseInt(record[1], 10, 0)
		if err != nil {
			log.Fatal(err)
		}

		quiz = append(quiz, &QuizType{
			Question:       record[0],
			ExpectedAnswer: a,
		})
	}

	// Run the quiz!
	results := []*ResultType{}
	var correct_answers int64
	for i, q := range quiz {
		fmt.Printf("Question %s: %s ", strconv.FormatInt(int64(i), 10), q.Question)

		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		answer = strings.TrimRight(answer, "\r\n")
		answer = strings.Trim(answer, " ")

		actual_answer, err := strconv.ParseInt(answer, 10, 0)
		if err != nil {
			log.Println("Invalid answer, expected a number.")
		}

		results = append(results, &ResultType{
			Quiz:         *q,
			ActualAnswer: actual_answer,
			Correct:      actual_answer == q.ExpectedAnswer,
		})

		if actual_answer == q.ExpectedAnswer {
			correct_answers += 1
		}

	}

	correct_answers_str := strconv.FormatInt(correct_answers, 10)
	quiz_len_str := strconv.FormatInt(int64(len(quiz)), 10)

	// Print the quiz results
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("You got %s correct answers out of %s questions!\n", correct_answers_str, quiz_len_str)
	fmt.Println("")

	if correct_answers < int64(len(quiz)) {
		fmt.Println("The questions you got incorrect were:")
		fmt.Println("")

		for _, res := range results {
			if res.Correct {
				continue
			}
			fmt.Println("-------")
			fmt.Println("")
			fmt.Printf("Question: %s\n", res.Quiz.Question)
			fmt.Printf("Your Answer: %s\n", strconv.FormatInt(res.ActualAnswer, 10))
			fmt.Printf("Correct Answer: %s\n", strconv.FormatInt(res.Quiz.ExpectedAnswer, 10))
			fmt.Println("")
			fmt.Println("-------")
			fmt.Println("")
			fmt.Println("")
		}
	}

}

type QuizType struct {
	Question       string
	ExpectedAnswer int64
}

type ResultType struct {
	Quiz         QuizType
	ActualAnswer int64
	Correct      bool
}
