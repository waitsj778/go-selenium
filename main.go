package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func main() {
	// chrome起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	// defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	// googleにアクセス
	if err := page.Navigate("https://www.google.com/"); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	// qというnameを持つ要素を取得する
	q := page.FindByName("q")
	// 検索文字列を入力
	q.Fill("")

	// 検索を実行
	if err := q.Submit(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
}
