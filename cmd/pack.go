package cmd

import (
	"Archive/lib/compression"
	"Archive/lib/compression/vlc"

	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// independent command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file ",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {

	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
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

	packed := encoder.Encode(string(data))

	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}

}

// a function that will generate a name for a compressed file based on the path to the current file
func packedFileName(path string) string {

	fileName := filepath.Base(path)

	ext := filepath.Ext(fileName)

	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + packedExtension
}

// child command for the root
func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method:vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
