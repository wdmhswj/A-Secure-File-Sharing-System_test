package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	test.Type(in, "Andy")							// 模拟用户输入 "Andy"
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}
}
