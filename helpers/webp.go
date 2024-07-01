package helpers

import (
	"fmt"
	"image"
	"os"

	"github.com/chai2010/webp"
)

func ConvertToWebP(pathInput, pathOutput string) error {
	file, errOpenInput := os.Open(pathInput)
	if errOpenInput != nil {
		return fmt.Errorf(
			"failed to open input file: %w",
			errOpenInput,
		)
	}
	defer file.Close()

	imgDecoded, format, errDecodeInput := image.Decode(file) //fails
	if errDecodeInput != nil {
		return fmt.Errorf(
			"failed to decode %s image: %w",
			pathInput,
			errOpenInput,
		)
	}

	fmt.Printf("Input format: %s\n", format)

	outFile, errCreateOutput := os.Create(pathOutput)
	if errCreateOutput != nil {
		return fmt.Errorf(
			"failed to create output file: %w",
			errOpenInput,
		)
	}
	defer outFile.Close()

	if errEncodeInput := webp.Encode(
		outFile,
		imgDecoded,
		&webp.Options{Lossless: true},
	); errEncodeInput != nil {
		return fmt.Errorf(
			"failed to encode image to WebP: %w",
			errEncodeInput,
		)
	}

	return nil
}
