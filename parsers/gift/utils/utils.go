package giftUtils

import (
	"errors"
	"fmt"
	giftAnswer "qparser/parsers/gift/models/answer"
	"qparser/utils"
	"regexp"
	"strings"
)

const ext = ".gift"
const outputDirName = "converted"

func MdFileGetQuestion(line string) (string, error) {

	r, err := parseMdFileQuestionRow(line)

	if err != nil {

		return "", err

	}

	return r[1], nil
}

func MdFileGetTopic(line string) (string, error) {

	r, err := parseMdFileQuestionRow(line)

	if err != nil {

		return "", err

	}

	fmt.Println("Topic ID:", r[2])

	return r[2], nil
}

func ParseMdFileAnswerRow(line string) (giftAnswer.Answer, error) {

	isAnswerCorrect := strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]")

	isAnswerIncorrect := strings.Contains(line, "- [ ]")

	checkAnswer := (isAnswerCorrect && !isAnswerIncorrect)

	answerText := strings.SplitN(line, "]", 2)[1]

	if len(answerText) < 1 {
		return giftAnswer.Answer{}, errors.New("parse error")
	}

	return giftAnswer.Answer{
		Text:      answerText,
		IsCorrect: checkAnswer,
		Valid:     true,
	}, nil

}

func GetGiftExt() string { return ext }

func GetOutputDirName() string { return outputDirName }

func IsNotCorrectAnswer(line string) bool {
	return !strings.Contains(line, "- [x]") && !strings.Contains(line, "- [X]")
}

func IsNotIncorrectAnswer(line string) bool { return !strings.Contains(line, "- [ ]") }

func IsNotTitle(line string) bool { return !(strings.Count(line, "#") == 1) }

func IsNotQuestion(line string) bool { return !strings.Contains(line, "##") }

func IsCorrectAnswer(line string) bool {
	return strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]")
}

func IsAnswer(line string) bool {
	return strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") || strings.Contains(line, "- [ ]")
}

func IsIncorrectAnswer(line string) bool { return strings.Contains(line, "- [ ]") }

func IsQuestion(line string) bool { return strings.Contains(line, "##") }

func ParseMathSymbol(line string) string {

	re := regexp.MustCompile(`_\{([^}]+)\}`)

	match := re.FindStringSubmatch(line)

	if len(match) <= 1 {

		return line

	}

	newPedice := utils.ToSubscript(match[1])

	updatedText := re.ReplaceAllString(line, newPedice)

	return updatedText
}

func parseMdFileQuestionRow(line string) ([]string, error) {

	var re *regexp.Regexp

	if strings.Contains(line, "topic") {

		re = regexp.MustCompile(`## (.*?)\s*\{topic:#([^\}]+)\}`)

	} else {

		re = regexp.MustCompile(`## (.*)`)

	}

	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {

		return nil, fmt.Errorf("nessun match valido trovato nella riga: %q", line)
	}

	qre := regexp.MustCompile(`\d+\.\s*`)

	parts := qre.Split(matches[1], -1)

	if len(parts) < 2 {

		parts = matches

	}

	return parts, nil
}
