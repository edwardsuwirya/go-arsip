package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"os"
	"strconv"
	"strings"
)

type LangConfig struct {
	DefLang string
}
type FileConfig struct {
	FileName string
	FileType string
}
type Config struct {
	LangConfig
	FileConfig
	ConsoleGui int
}

func NewConfig() *Config {
	if len(os.Args) == 4 {
		langArgs := os.Args[1]
		fileArgs := os.Args[2]
		consoleBasedArgs := os.Args[3]

		errLang := validation.Validate(langArgs, validation.Required, validation.Length(2, 2))
		errFile := validation.Validate(fileArgs, validation.Required, validation.Length(2, 100))
		errConsole := validation.Validate(consoleBasedArgs, validation.Required, validation.Length(1, 1), is.Digit)

		if errLang != nil || errFile != nil || errConsole != nil {
			panic("Failed configuration")
		}
		c, _ := strconv.Atoi(consoleBasedArgs)
		fileType := strings.Split(fileArgs, ".")
		return &Config{
			LangConfig: LangConfig{
				DefLang: strings.ToUpper(langArgs),
			},
			FileConfig: FileConfig{
				FileName: fileArgs,
				FileType: fileType[1],
			},
			ConsoleGui: c,
		}
	} else {
		panic("No configuration default is defined")
	}

}
