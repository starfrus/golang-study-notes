package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputData() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {

		return "", fmt.Errorf("! Ошика считывания данных:%w", scanner.Err())
	}
	return scanner.Text(), nil
}

func CheckEmptyInputData(text string) error {
	text = strings.TrimSpace(text)
	if text == "" {
		err := errors.New("data is empty")
		return fmt.Errorf("! Строка состоит из пробелов: %w", err)
	}

	return nil
}

func CheckNotNumberInputData(text string) error {
	if _, err := strconv.Atoi(text); err == nil {
		err = errors.New("data is number")
		return fmt.Errorf("! Строка состоит из чисел: %w", err)
	}
	return nil
}

func CheckNumberInputData(text string) (int, error) {
	if err := CheckEmptyInputData(text); err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("! Ошибка проверки на число: %w", err)
	}
	return num, nil
}

func InputID() (int, error) {
	fmt.Print("> Введите id: ")
	text, err := InputData()
	if err != nil {
		return 0, err
	}
	id, err := CheckNumberInputData(text)
	if err != nil {
		return 0, err
	}
	return id, nil
}
