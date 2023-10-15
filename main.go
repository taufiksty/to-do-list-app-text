package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/taufiksty/to-do-list-app-text/database"
	"github.com/taufiksty/to-do-list-app-text/entity"
	"github.com/taufiksty/to-do-list-app-text/repository/mysql"
	"github.com/taufiksty/to-do-list-app-text/service"
)

func clearTerminal() {
	var cmd *exec.Cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func listToDo(tasks []entity.Task) {
	fmt.Println("Daftar Tugas : ")
	for _, task := range tasks {
		status := "Belum selesai"

		if task.Done {
			status = "Selesai"
		}

		fmt.Printf("%d. %s - %s [%s]\n", task.Id, task.Title, task.Description, status)
	}
}

func inputToDo() (string, string, bool) {
	var title, description string
	var done string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title : ")
	if scanner.Scan() {
		title = scanner.Text()
	}

	fmt.Print("Description : ")
	if scanner.Scan() {
		description = scanner.Text()
	}

	fmt.Print("Done [y/n]: ")
	fmt.Scan(&done)

	var doneBool bool

	if strings.ToLower(done) == "y" {
		doneBool = true
	} else {
		doneBool = false
	}

	return title, description, doneBool
}

func menu() {
	fmt.Println("\n------- MENU -------")
	fmt.Println("[1] Tambah To Do")
	fmt.Println("[2] Ubah To Do")
	fmt.Println("[3] Hapus To Do")
	fmt.Println("[4] Keluar")
}

func main() {
	repo := mysql.NewTaskRepository(database.GetConnection())

	taskService := service.TaskService{Repository: repo}

	fmt.Println("------- TO DO LIST APP CLI -------")
	fmt.Println("Press enter to continue")

	fmt.Scanln()
	clearTerminal()

	i := "n"
	for i == "n" {
		// Daftar To Do
		all, _ := taskService.GetAll()
		listToDo(all)

		// Menu
		var menuOption int
		menu()
		fmt.Println("Pilih opsi Anda?")
		fmt.Scanln(&menuOption)
		if menuOption != 2 && menuOption != 3 {
			clearTerminal()
		}

		if menuOption == 1 {
			title, description, done := inputToDo()

			newTask := entity.Task{
				Title:       title,
				Description: description,
				Done:        done,
			}

			_, err := taskService.AddTask(newTask)

			if err != nil {
				fmt.Println("Error :", err)
			}

			clearTerminal()

		} else if menuOption == 2 {

			var no int
			fmt.Print("No. berapa yang ingin diubah? ")
			fmt.Scan(&no)

			title, description, done := inputToDo()

			updatedTask := entity.Task{
				Title:       title,
				Description: description,
				Done:        done,
			}

			_, err := taskService.UpdateTask(no, updatedTask)

			if err != nil {
				fmt.Println("Error :", err)
			}

			clearTerminal()

		} else if menuOption == 3 {
			var no int
			fmt.Print("No. berapa yang ingin dihapus? ")
			fmt.Scan(&no)

			_, err := taskService.DestroyTask(no)

			if err != nil {
				fmt.Println("Err :", err)
			}

			clearTerminal()

		} else if menuOption == 4 {
			var option string
			fmt.Println("Anda yakin ingin keluar [y/n]?")
			fmt.Scanln(&option)
			i = strings.ToLower(option)
		}

	}

}
