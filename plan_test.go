package main

import (
	"fmt"
	"os"
	"testing"
)

// TestGetPlan テストファイル読み込み
func TestGetPlan(t *testing.T) {
	err := GetPlan("zgoc.yml")
	if err != nil {
		t.Error(err)
	}
}

// TestNoPlan テストファイル読み込み失敗
func TestNoPlan(t *testing.T) {
	err := GetPlan("noplan")
	if err == nil {
		t.Error("noplan")
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
	fmt.Fprintln(file, "syntax error") //書き込み

	err = GetPlan("err.yml")
	if err == nil {
		t.Error("syntax yaml")
	}
}
