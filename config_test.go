package main

import (
	"fmt"
	"os"
	"testing"
)

// TestGetConfig テストファイル読み込み
func TestGetConfig(t *testing.T) {
	err := GetConfig("zgoc.yml")
	if err != nil {
		t.Error(err)
	}
}

// TestNoConfig テストファイル読み込み失敗
func TestNoConfig(t *testing.T) {
	err := GetConfig("noconfig")
	if err == nil {
		t.Error("noconfig")
	}
}

// TestErrYaml テストファイル読み込み失敗
func TestErrYaml(t *testing.T) {
	file, err := os.OpenFile("err.yml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		t.Error(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "test error")

	err = GetConfig("err.yml")
	if err == nil {
		t.Error("syntax yaml")
	}
}
