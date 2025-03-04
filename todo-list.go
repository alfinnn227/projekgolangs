package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Description string
	Done        bool
}

var tasks []Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Todo List ===")
		fmt.Println("1. Tambah tugas")
		fmt.Println("2. Lihat tugas")
		fmt.Println("3. Tandai selesai")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			showTasks()
		case "3":
			markDone(scanner)
		case "4":
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Opsi tidak valid!")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan tugas: ")
	scanner.Scan()
	taskDesc := scanner.Text()
	tasks = append(tasks, Task{Description: taskDesc, Done: false})
	fmt.Println("Tugas berhasil ditambahkan!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas!")
		return
	}
	fmt.Println("\nDaftar Tugas:")
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
}

func markDone(scanner *bufio.Scanner) {
	showTasks()
	fmt.Print("Masukkan nomor tugas yang selesai: ")
	scanner.Scan()
	var index int
	_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Nomor tidak valid!")
		return
	}
	tasks[index-1].Done = true
	fmt.Println("Tugas ditandai sebagai selesai!")
}
