package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Task struct {
	Title, Description string
	Done               bool
}

func clearTerminal() {
	var cmd *exec.Cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func listToDo(tasks []Task) {
	fmt.Println("Daftar Tugas : ")
	for i, task := range tasks {
		status := "Belum selesai"

		if task.Done {
			status = "Selesai"
		}

		fmt.Printf("%d. %s - %s [%s]\n", i+1, task.Title, task.Description, status)
	}
}

func menu() {
	fmt.Println("\n------- MENU -------")
	fmt.Println("[1] Tambah To Do")
	fmt.Println("[2] Ubah To Do")
	fmt.Println("[3] Hapus To Do")
	fmt.Println("[4] Keluar")
}

func main() {
	var tasks []Task

	tasks = append(tasks, Task{
		Title:       "Belajar",
		Description: "Belajar golang dasar",
		Done:        false,
	})
	tasks = append(tasks, Task{
		Title:       "Olahraga",
		Description: "Jogging santai",
		Done:        true,
	})

	fmt.Println("------- TO DO LIST APP CLI -------")
	fmt.Println("Press enter to continue")

	fmt.Scanln()
	clearTerminal()

	i := "n"
	for i == "n" {
		// Daftar To Do
		listToDo(tasks)

		// Menu
		var menuOption int
		menu()
		fmt.Println("Pilih opsi Anda?")
		fmt.Scanln(&menuOption)
		if menuOption != 2 && menuOption != 3 {
			clearTerminal()
		}

		if menuOption == 1 {
			var title, description string
			var done string

			fmt.Print("Title : ")
			fmt.Scan(&title)
			fmt.Print("Description : ")
			fmt.Scan(&description)
			fmt.Print("Done [y/n]: ")
			fmt.Scan(&done)

			var doneBool bool
			if strings.ToLower(done) == "y" {
				doneBool = true
			} else {
				doneBool = false
			}

			tasks = append(tasks, Task{title, description, doneBool})
		} else if menuOption == 2 {
			var no int
			fmt.Print("No. berapa yang ingin diubah? ")
			fmt.Scan(&no)

			var title, description string
			var done string

			fmt.Print("Title : ")
			fmt.Scan(&title)
			fmt.Print("Description : ")
			fmt.Scan(&description)
			fmt.Print("Done [y/n]: ")
			fmt.Scan(&done)

			var doneBool bool
			if strings.ToLower(done) == "y" {
				doneBool = true
			} else {
				doneBool = false
			}

			tasks[no-1].Title = title
			tasks[no-1].Description = description
			tasks[no-1].Done = doneBool
		} else if menuOption == 3 {
			var no int
			fmt.Print("No. berapa yang ingin dihapus? ")
			fmt.Scan(&no)

			index := no - 1

			tasks = append(tasks[:index], tasks[no:]...)
		} else if menuOption == 4 {
			var option string
			fmt.Println("Anda yakin ingin keluar [y/n]?")
			fmt.Scanln(&option)
			i = strings.ToLower(option)
		}

	}

}
