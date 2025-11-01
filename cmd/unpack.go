package cmd

import (
	"Archive/lib/compression"
	"Archive/lib/compression/vlc"

	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file ",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown method compression")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	defer func() {
		if err := r.Close(); err != nil {
			handleErr(err)
		}
	}()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := decoder.Decode(data)

	fmt.Println(string(data))

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	ext := filepath.Ext(fileName)

	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method:vlc")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
