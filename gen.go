package main

import (
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	guidCmd := &cobra.Command{
		Use:   "guid",
		Short: "Generate a GUID",
		Run:   guid,
	}

	dateCmd := &cobra.Command{
		Use:   "date",
		Short: "Gets the current data",
		Run:   date,
	}

	rootCmd.AddCommand(guidCmd, dateCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error running root command: %v", err)
	}
}

func guid(cmd *cobra.Command, args []string) {
	out := uuid.New().String()
	fmt.Println(out)

	if err := clipboard.WriteAll(out); err != nil {
		log.Fatalf("error copying to clipboard: %v", err)
	}
}

func date(cmd *cobra.Command, args []string) {
	out := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Println(out)

	if err := clipboard.WriteAll(out); err != nil {
		log.Fatalf("error copying to clipboard: %v", err)
	}
}
