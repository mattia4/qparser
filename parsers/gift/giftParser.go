package giftParser

import (
	"bufio"
	"fmt"
	"os"
	gift "qparser/parsers/gift/utils"
	"strings"
)

/*type GiftParser struct{}

func NewGiftParser() *GiftParser {
	return &GiftParser{}
}*/

func ExtractCodeBlock(scanner *bufio.Scanner) string {

	var codeLines []string

	codeLines = append(codeLines, `"""`)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.TrimSpace(line) == "```" {

			break

		}
		// TODO change
		if strings.Contains(line, "{") {

			line = strings.ReplaceAll(line, "{", "\\{")

		}

		if strings.Contains(line, "}") {

			line = strings.ReplaceAll(line, "}", "\\}")

		}

		if strings.Contains(line, "=") {

			line = strings.ReplaceAll(line, "=", "\\=")

		}

		if strings.Contains(line, ":") {

			line = strings.ReplaceAll(line, ":", "\\:")

		}

		codeLines = append(codeLines, line)
	}

	codeLines = append(codeLines, `"""`)

	return strings.Join(codeLines, "\n")

}

func ExtractAnswerBlock(scanner *bufio.Scanner, firstLine string) (string, error) {

	var codeLines []string

	line, err := gift.ParseMdFileAnswerRow(firstLine)

	if err != nil {

		return "", nil

	}

	codeLines = append(codeLines, line.Text)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.TrimSpace(line) == "" {

			break

		}

		ans, err := gift.ParseMdFileAnswerRow(line)

		if err != nil {

			return "", err

		}

		codeLines = append(codeLines, ans.Text)

	}

	return strings.Join(codeLines, "\n"), nil

}

func ExtractQuestionBlock(line string) string {

	var codeLines []string

	question, err := "", error(nil)

	question, err = gift.MdFileGetQuestion(line)

	question = strings.ReplaceAll(question, "$", "")

	if err != nil {

		fmt.Fprintln(os.Stderr, "err: ", err.Error())

	}

	if strings.Contains(question, "*") {

		question = strings.Trim(question, "*")

	}

	if strings.Contains(question, "_{") {

		question = gift.ParseMathSymbol(question)

	}

	codeLines = append(codeLines, question)

	return strings.Join(codeLines, "\n")
}
