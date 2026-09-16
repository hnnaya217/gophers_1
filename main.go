package main

import (
	"fmt"
	"os"

	firsttask "task-gdgoc-unj/first-task"
)

type route struct {
	name        string
	description string
	run         func() error
}

func routes() []route {
	return []route{
		{
			name:        "first-task",
			description: "Kasir toko kelontong: validasi pembayaran & hitung kembalian",
			run:         func() error { return firsttask.Run(os.Stdin, os.Stdout) },
		},
	}
}

func main() {
	available := routes()

	if len(os.Args) < 2 {
		printUsage(available)
		return
	}

	selected := os.Args[1]
	for _, r := range available {
		if r.name != selected {
			continue
		}

		if err := r.run(); err != nil {
			fmt.Fprintln(os.Stderr, "gagal menjalankan task:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Fprintf(os.Stderr, "task %q tidak ditemukan\n\n", selected)
	printUsage(available)
	os.Exit(1)
}

func printUsage(available []route) {
	fmt.Println("Pemakaian: go run . <nama-task>")
	fmt.Println()
	fmt.Println("Task yang tersedia:")
	for _, r := range available {
		fmt.Printf("  %-12s  %s\n", r.name, r.description)
	}
}
