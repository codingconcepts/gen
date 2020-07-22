package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	uuidCmd := &cobra.Command{Use: "uuid", Short: "Generate a UUID", Run: guid}
	dateCmd := &cobra.Command{Use: "date", Short: "Gets the current date", Run: date}
	unixCmd := &cobra.Command{Use: "unix", Short: "Gets the current date in epoch seconds", Run: unix}

	rootCmd.AddCommand(uuidCmd, dateCmd, unixCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error running root command: %v", err)
	}
}

func guid(cmd *cobra.Command, args []string) {
	out := uuid.New().String()
	fmt.Println(out)

	mustCopy(out)
}

func date(cmd *cobra.Command, args []string) {
	out := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Println(out)

	mustCopy(out)
}

func unix(cmd *cobra.Command, args []string) {
	out := time.Now().UTC().Unix()
	fmt.Println(out)

	mustCopy(strconv.FormatInt(out, 10))
}

func mustCopy(s string) {
	if err := clipboard.WriteAll(s); err != nil {
		log.Fatalf("error copying %q to clipboard: %v", s, err)
	}
}
