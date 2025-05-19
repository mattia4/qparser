package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	log "qparser/logger"
	giftParser "qparser/parsers/gift"
	giftUtils "qparser/parsers/gift/utils"
	fileUtil "qparser/utils/file"
	"strings"
)

func main() {

	dir := flag.String("inputdir", "./input", "directory with input files")

	format := flag.String("format", "md", "file format (no dot)")

	outFormat := flag.String("outFormat", "gift", "output file format (no dot)")

	outDir := flag.String("outdir", "results", "output directory")

	verbose := flag.Bool("verbose", false, "detailed output")

	logger := log.NewLogger(*verbose)

	flag.Parse()

	info, err := os.Stat(*dir)

	if err != nil {

		logger.Fatal("Directory does not exist: %s", err)

	}

	if !info.IsDir() {

		logger.Fatal("Path is not a directory")

	}

	files, err := filepath.Glob(filepath.Join(*dir, "*."+*format))

	if err != nil {

		logger.Fatal("Error while seeking for files: %s", err.Error())

	}

	if len(files) == 0 {

		logger.Fatal("No file. %s found in %s", *format, *dir)

	}

	for _, inputFileName := range files {

		logger.Info("Processing file: %s", inputFileName)

		if err := convertFile(logger, inputFileName, *outFormat, *outDir); err != nil {

			logger.Error("failed to convert %s: %v", inputFileName, err)

		}
	}
}

func convertFile(logger *log.Logger, inputFileName string, outFormat string, outDir string) error {

	inputFile, err := os.Open(inputFileName)

	if err != nil {

		return fmt.Errorf("cannot open input file: %w", err)

	}

	defer inputFile.Close()

	outputFile, outputPath, err := fileUtil.CreateOutputFile(inputFileName, outFormat, outDir)

	if err != nil {

		return fmt.Errorf("cannot create output file: %w", err)

	}

	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)

	if err := processInput(scanner, outputFile); err != nil {

		return fmt.Errorf("error during input processing: %w", err)

	}

	logger.Success("%s -> %s", inputFileName, outputPath)

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

			answerBlock, err := giftParser.ExtractAnswerBlock(scanner, line)

			if err != nil {

				return err

			}

			if codeBlock == "" {

				fmt.Fprintf(file, "{\n%s\n", answerBlock)

			} else {

				fmt.Fprintf(file, "%s\n", answerBlock)

			}

			fmt.Fprint(file, "}\n\n")

			continue
		}

	}

	if err := scanner.Err(); err != nil {

		fmt.Fprintln(os.Stderr, "\n}")

	}

	return nil

}
