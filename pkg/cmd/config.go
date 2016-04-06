package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	// Used to marshal the configuration into YAML
	UserName string `mapstructure:"user.name" yaml:"user.name"`
	UserEmail string `mapstructure:"user.email" yaml:"user.email"`
}

func isValidKey(key string) bool {
	switch key {
	case
		"user.name",
		"user.email":
		return true
	}
	return false
}

func NewCmdConfig(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config name [value]",
		Short: "Get or set global options",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dispatchConfig(args, os.Stdout)
		},
	}
	return cmd
}

func dispatchConfig(args []string, out io.Writer) error {
	if len(args) == 0 {
		return errors.New("Not enough arguments")
	} else if len(args) == 1 {
		return errors.New("Any operation other than setting a value is not implemented yet")
	} else if len(args) == 2 {
		return setConfigValue(args[0], args[1])
	} else {
		return errors.New("Too many arguments")
	}
}

func setConfigValue(key, value string) error {
	if !isValidKey(key) {
		return fmt.Errorf("'%s' is not a valid configuration key", key)
	}

	viper.Set(key, value)
	return nil
}

func unmarshalConfig() (Config, error) {
	var C Config
	err := viper.Unmarshal(&C)
	return C, err
}
