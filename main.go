package main

import (
	"bufio"
	"fmt"
	"os"
	"qparser/models/file"
	giftParser "qparser/parsers/gift"
	giftUtils "qparser/parsers/gift/utils"
	fileUtil "qparser/utils/file"
	"strings"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Fprintln(os.Stderr, file.GetErrStdInMustProvideAtLeastOneArg())

		return

	}

	err := os.MkdirAll(giftUtils.GetOutputDirName(), os.ModePerm)

	if err != nil {

		fmt.Println("err -> ", err)

		return

	}

	for _, inputFileName := range os.Args[1:] {

		if err := convertFile(inputFileName); err != nil {

			fmt.Fprintf(os.Stderr, "failed to convert %s: %v\n", inputFileName, err)

		}

	}

}

func convertFile(inputFileName string) error {

	inputFile, err := os.Open(inputFileName)

	if err != nil {

		return fmt.Errorf("cannot open input file: %w", err)

	}

	defer inputFile.Close()

	outputFile, outputPath, err := fileUtil.CreateOutputFile(inputFileName, giftUtils.GetGiftExt(), giftUtils.GetOutputDirName())

	if err != nil {

		return fmt.Errorf("cannot create output file: %w", err)

	}

	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)

	if err := processInput(scanner, outputFile); err != nil {

		return err

	}

	fmt.Printf("done. converted %s -> %s\n", inputFileName, outputPath)

	return nil
}

func processInput(scanner *bufio.Scanner, file *os.File) error {

	codeBlock := ""

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, "```") {

			codeBlock = giftParser.ExtractCodeBlock(scanner)

			fmt.Fprintf(file, "%s\n{\n", codeBlock)

			continue
		}

		if giftUtils.IsQuestion(line) {

			questionBlock := giftParser.ExtractQuestionBlock(line)

			fmt.Fprintf(file, "%s \n", questionBlock)

		}

		if giftUtils.IsAnswer(line) {

			answerBlock := giftParser.ExtractAnswerBlock(scanner)

			if codeBlock == "" {

				fmt.Fprintf(file, "{\n%s\n", answerBlock)

			} else {

				fmt.Fprintf(file, "%s\n", answerBlock)

			}

			fmt.Fprintln(file, "}")

			continue
		}

	}

	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "\n}")

	}

	return nil

}
