package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"

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
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	out := fmt.Sprintf("%x-%x-%x-%x-%x\n", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
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
