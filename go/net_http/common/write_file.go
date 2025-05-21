package common

import (
	"bufio"
	"io/ioutil"
	"os"
)

func WriteFile01(path string, text string) (string, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()
	w1 := bufio.NewWriter(file)
	_, err = w1.WriteString(text)
	if err != nil {
		return "", err
	}
	err = w1.Flush()
	if err != nil {
		return "", err
	}
	return text, nil
}

func WriteFile02(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile03(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile04(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile05(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile06(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile07(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile08(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile09(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}

func WriteFile10(path string, text string) (string, error) {
	err := ioutil.WriteFile(path, []byte(text), os.ModePerm)
	return text, err
}
