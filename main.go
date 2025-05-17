package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	giftParser "qparser/parsers/gift"
	giftUtils "qparser/parsers/gift/utils"
	fileUtil "qparser/utils/file"
	"strings"
)

func main() {

	dir := flag.String("inputdir", "./input", "directory with input files")

	format := flag.String("format", "md", "file format (no dot)")

	outDir := flag.String("outdir", "results", "output directory")

	verbose := flag.Bool("verbose", false, "detailed output")

	flag.Parse()

	files, err := filepath.Glob(filepath.Join(*dir, "*."+*format))

	if err != nil {

		log.Fatal("Error while seeking for files:", err.Error())

	}

	if len(files) == 0 {

		log.Fatalf("❌ No file. %s found in %s", *format, *dir)

	}

	for _, inputFileName := range files {

		if *verbose {

			fmt.Printf("📄 Elaboro file: %s\n", inputFileName)

		}

		if err := convertFile(inputFileName, *format, *outDir, *verbose); err != nil {

			fmt.Fprintf(os.Stderr, "failed to convert %s: %v\n", inputFileName, err)

		}
	}
}

func convertFile(inputFileName string, format string, outDir string, verbose bool) error {

	inputFile, err := os.Open(inputFileName)

	if err != nil {

		return fmt.Errorf("cannot open input file: %w", err)

	}

	defer inputFile.Close()

	outputFile, outputPath, err := fileUtil.CreateOutputFile(inputFileName, format, outDir)

	if err != nil {

		return fmt.Errorf("cannot create output file: %w", err)

	}

	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)

	if err := processInput(scanner, outputFile); err != nil {

		return fmt.Errorf("error during input processing: %w", err)

	}

	if verbose {

		fmt.Printf("✅ %s -> %s\n", inputFileName, outputPath)

	} else {

		fmt.Printf("converted %s -> %s\n", inputFileName, outputPath)

	}

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
