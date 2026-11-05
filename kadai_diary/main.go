package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

// 日記データの保存用関数
func saveDiary(entries []Entry) error {
	fmt.Println("日記データを保存します")

	jsonData, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("変換失敗しました。エラー：", err)
		return err
	}

	err = os.WriteFile("diary.json", jsonData, 0644)
	if err != nil {
		fmt.Println("書き込みに失敗しました：", err)
		return err
	}

	fmt.Println("diary.jsonへの保存が成功しました")
	return nil
}

// 日記データのロード用関数
func loadDiary() ([]Entry, error) {
	fmt.Println("日記データを読み込みます")

	diaryData, err := os.ReadFile("diary.json")
	if os.IsNotExist(err) {
		// 初回はファイルがないため、空のスライスを返す
		return []Entry{}, nil
	}
	if err != nil {
		fmt.Println("日記データの読み込みに失敗しました: ", err)
		return nil, err
	}

	var loadedData []Entry

	err = json.Unmarshal(diaryData, &loadedData)
	if err != nil {
		fmt.Println("jsonからGoスライスへの変換に失敗しました：", err)
		return nil, err
	}

	fmt.Println("日記データの読み込みに成功しました")
	return loadedData, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	// 1番目の引数（コマンド）を取得
	command := os.Args[1]

	// add or listコマンドによって処理を切り替え
	switch command {
	case "add":
		// 引数にタイトル・本文があれば取得する
		var title, content string
		if len(os.Args) > 2 {
			title = os.Args[2]
		}
		if len(os.Args) > 3 {
			content = os.Args[3]
		}

		// loadDiary()関数で既存の日記データを読み込む
		existingEntries, err := loadDiary()
		if err != nil {
			fmt.Println("エラー:", err)
			return
		}

		// 新しい日記エントリーを作成してスライスに追加する
		e := Entry{
			Timestamp: time.Now(),
			Title:     title,
			Content:   content,
		}

		fmt.Println("追加:", e)

		existingEntries = append(existingEntries, e)

		// saveDiary()関数でファイルに保存する
		saveDiary(existingEntries)

	case "list":
		// loadDiary()関数で既存の日記データを読み込む
		lists, err := loadDiary() // 返り値が２つ
		if err != nil {
			fmt.Println("エラー:", err)
			return
		}

		// 日記エントリーのスライスに対して繰り返し処理を行い、1件ずつの日記エントリーを出力する
		for i, Entry := range lists {
			fmt.Println("--- 日記データ ---")
			fmt.Printf("エントリー%d\n", i+1)
			fmt.Printf("日時：%v\n", Entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("タイトル：%s\n", Entry.Title)
			fmt.Printf("本文：\n%s\n", Entry.Content)
		}
	}
}
