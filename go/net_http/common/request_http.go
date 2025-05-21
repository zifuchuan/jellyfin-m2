package common

import (
	"net/http"
	"strconv"
)

func RequestHttp01(url string) (string, error) {
	// Head Post PostForm
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp02(url string) (string, error) {
	client := http.Client{}
	// Head Post PostForm
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp03(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Token", "123")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp04(url string) (string, error) {
	return "", nil
}

func RequestHttp05(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp06(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp07(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp08(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp09(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}

func RequestHttp10(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(resp.StatusCode), nil
}
