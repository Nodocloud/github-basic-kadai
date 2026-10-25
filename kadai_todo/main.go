package main

import "fmt"

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *ToDo) Complete() {
	fmt.Printf("ID:%dのToDoを完了に更新します\n", t.ID)

	t.Completed = true
}

func printTodos(todos []*ToDo) {
	for _, todo := range todos {
		if todo.Completed == true {
			fmt.Printf("[完了](ID:%d)%s\n", todo.ID, todo.Title)
		} else {
			fmt.Printf("[未完了](ID:%d)%s\n", todo.ID, todo.Title)
		}
	}

}

func main() {
	todos := []*ToDo{
		{
			ID:        1,
			Title:     "学習計画",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "環境構築",
			Completed: false,
		},
		{
			ID:        3,
			Title:     "基礎文法",
			Completed: false,
		},
	}

	// todosスライスの初期値
	printTodos(todos)

	// ポインタで完了状態に更新
	todos[0].Complete() // ID1
	todos[1].Complete() // ID2

	// 変更後の値
	printTodos(todos)

}
