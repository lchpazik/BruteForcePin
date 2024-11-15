package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	targetPassword = "1111"
)

func Menu() {
	fmt.Println("\n\n\n\n\t\t\t██████  ██████  ██    ██ ████████ ███████  ██████  ██████   ██████ ███████")
	fmt.Println("\t\t\t██   ██ ██   ██ ██    ██    ██    ██      ██    ██ ██   ██ ██      ██      ")
	fmt.Println("\t\t\t██████  ██████  ██    ██    ██    █████   ██    ██ ██████  ██      █████   ")
	fmt.Println("\t\t\t██   ██ ██   ██ ██    ██    ██    ██      ██    ██ ██   ██ ██      ██      ")
	fmt.Println("\t\t\t██████  ██   ██  ██████     ██    ██       ██████  ██   ██  ██████ ███████")

	fmt.Println("\n\n\t\t\t\t═══════════════════════════════════════════════")
	fmt.Println("\t\t\t\t\t\t    lchpazik   ")
	fmt.Println("\t\t\t\t═══════════════════════════════════════════════")
	fmt.Println("\t\t\t\t      1. Начать подбирать 4-х значный пинкод")
	fmt.Println("\t\t\t\t\t     2. Очистить консоль")
	fmt.Println("\t\t\t\t\t\t  3. Выйти")
	fmt.Println("\t\t\t\t-----------------------------------------------")
	fmt.Print("\t\t   Выберите действие (1 - Начало, 2 - Очистка консоли, 3 - Выход из программы): ")
}

func clearConsole() {
	var cmd *exec.Cmd
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GenerateRandomPinCode() string {
	password := ""
	for i := 0; i < 4; i++ {
		password += strconv.Itoa(rand.Intn(10))
	}
	return password
}

func checkPassword(password, target string) bool {
	return password == target
}

func readPasswordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passwords []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			passwords = append(passwords, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	isRunning := false

	for {
		Menu()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:

			PopularPincode, err := readPasswordsFromFile("popularpincode.txt")

			if err != nil {
				fmt.Println("Ошибка чтения файла:", err)
				return
			}

			startTime := time.Now()

			for _, password := range PopularPincode {
				clearConsole()
				fmt.Println("-> Пробую популярный пароль:", password)

				if checkPassword(password, targetPassword) {
					elapsedTime := time.Since(startTime)
					seconds := int(elapsedTime.Seconds()) % 60
					milliseconds := int(elapsedTime.Milliseconds()) % 1000

					fmt.Printf("✅ Пароль найден: %s\n", password)
					fmt.Printf("-> Время работы программы: %02d:%03d\n", seconds, milliseconds)
					isRunning = true
					break
				}
			}

			if isRunning == false {
				for {
					password := GenerateRandomPinCode()
					fmt.Println("Пробую случайный пароль:", password)

					if checkPassword(password, targetPassword) {
						elapsedTime := time.Since(startTime)
						seconds := int(elapsedTime.Seconds()) % 60
						milliseconds := int(elapsedTime.Milliseconds()) % 1000

						fmt.Printf("Пароль найден: %s\n", password)
						fmt.Printf("Время работы программы: %02d:%03d\n", seconds, milliseconds)
						break
					}
				}
			}

		case 2:
			clearConsole()
		case 3:
			fmt.Print("\n-> Вы уверены, что хотите выйти? (Y/N): ")

			var confirm string
			fmt.Scan(&confirm)

			if strings.ToUpper(confirm) == "Y" {

				fmt.Println("-> Завершаю программу...")
				return

			} else if strings.ToUpper(confirm) == "N" {

				fmt.Println("-> Возвращаюсь в главное меню...")
				time.Sleep(time.Duration(2) + time.Second)
				clearConsole()

			} else {

				clearConsole()
				fmt.Println("Неверный ввод. Возвращаюсь в главное меню...")

			}

		default:

			fmt.Println("\n-> Неверный выбор. Попробуйте снова.")
			time.Sleep(time.Duration(1) + time.Second)
			clearConsole()
		}
	}
}
