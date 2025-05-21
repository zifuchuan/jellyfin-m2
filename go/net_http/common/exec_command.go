package common

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os/exec"
)

func ExecCommand01(cmd string) (string, error) {
	ctx := context.Background()
	out, err := exec.CommandContext(ctx, "sh", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand02(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand03(cmd string) (string, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(cmd)
	cmd = buf.String()
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand04(cmd string) (string, error) {
	var cmdStruct struct {
		Cmd string `json:"cmd"`
	}
	err := json.Unmarshal([]byte(cmd), &cmdStruct)
	if err != nil {
		return "", err
	}
	out, err := exec.Command("bash", "-c", cmdStruct.Cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand05(cmd string) (string, error) {
	var cmdStruct struct {
		XMLName xml.Name `xml:"command"`
		Cmd     string   `xml:"cmd,attr"`
	}
	// https://www.cnblogs.com/wanghui-garcia/p/10430013.html
	// <?xml version="1.0" encoding="utf-8"?><command cmd="whoami"/>
	err := xml.Unmarshal([]byte(cmd), &cmdStruct)
	if err != nil {
		return "", err
	}
	out, err := exec.Command("bash", "-c", cmdStruct.Cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand06(cmd string) (string, error) {
	cmd = fmt.Sprintf("w%s", cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand07(cmd string) (string, error) {
	cmd = fmt.Sprint(cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

func ExecCommand08(cmd string) (string, error) {
	cmd = fmt.Sprintln(cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	return string(out), err
}

type execCommand09Struct struct {
	buf string
}

func (s *execCommand09Struct) Write(p []byte) (n int, err error) {
	s.buf = string(p)
	n = len(p)
	return
}

func ExecCommand09(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = fmt.Fprintln(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand10(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand11(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand12(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand13(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand14(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}

func ExecCommand15(cmd string) (string, error) {
	s := &execCommand09Struct{}
	_, _ = io.WriteString(s, cmd)
	out, err := exec.Command("bash", "-c", s.buf).CombinedOutput()
	return string(out), err
}
