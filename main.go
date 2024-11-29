package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func randomStringWithErrorCheck(length int) string {
	str, err := randomString(length)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		os.Exit(1)
	}
	return str
}

func createCSVFile(dir, fileName string, fileSize int64, numColumns int) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	filePath := filepath.Join(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := make([]string, numColumns)
	for i := 0; i < numColumns; i++ {
		header[i] = fmt.Sprintf("column%d", i+1)
	}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header: %w", err)
	}

	currentSize := int64(0)
	rowNumber := 1

	for currentSize < fileSize {
		row := make([]string, numColumns)
		row[0] = strconv.Itoa(rowNumber)
		for i := 1; i < numColumns; i++ {
			row[i] = randomStringWithErrorCheck(10)
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row: %w", err)
		}

		rowSize := int64(len(row[0]) + (len(row)-1)*10 + (len(row) - 1)) // 10文字 + カンマ
		currentSize += rowSize
		rowNumber++
	}

	fmt.Println("CSV file created successfully:", filePath)
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "csv-generator",
	Short: "CSV Generator generates CSV files with random data",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		fileNamePrefix, _ := cmd.Flags().GetString("prefix")
		fileSize, _ := cmd.Flags().GetInt64("file-size")
		numColumns, _ := cmd.Flags().GetInt("num-columns")
		numFiles, _ := cmd.Flags().GetInt("num-files")

		eg, _ := errgroup.WithContext(cmd.Context())
		eg.SetLimit(10)

		for fileIndex := 1; fileIndex <= numFiles; fileIndex++ {
			outputFile := fmt.Sprintf("%s_%d.csv", fileNamePrefix, fileIndex)
			eg.Go(func() error {
				return createCSVFile(dir, outputFile, fileSize, numColumns)
			})

			if err := eg.Wait(); err != nil {
				fmt.Println("Error creating CSV file:", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.Flags().StringP("dir", "d", "./out", "Output directory")
	rootCmd.Flags().Int64P("file-size", "s", 1024, "File size in bytes")
	rootCmd.Flags().IntP("num-columns", "c", 5, "Number of columns")
	rootCmd.Flags().IntP("num-files", "f", 1, "Number of files")
	rootCmd.Flags().StringP("prefix", "p", "out", "File name prefix")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
