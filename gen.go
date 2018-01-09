package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

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
	fmt.Printf("%x-%x-%x-%x-%x\n", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func date(cmd *cobra.Command, args []string) {
	fmt.Println(time.Now().UTC().Format(time.RFC3339Nano))
}
