package repository

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gh0st3e/TodoForTechArt/internal/database"
	"github.com/gh0st3e/TodoForTechArt/internal/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func Index(w http.ResponseWriter, r *http.Request) { // Главная страница
	t, err := template.ParseFiles("internal/templates/index.html") // Подключение html файла

	if err != nil {
		errors.Wrap(err, "repository.Index.ParseFiles couldn't load template")
	}

	db, err := database.ConnectToDB()
	if err != nil {
		errors.Wrap(err, "repository.Index.ConnectToDB couldn't connect to database")
	}
	defer db.Close()

	res, err := db.Query("SELECT  * FROM `todo` ORDER BY `Id`") // Запрос к таблице
	if err != nil {
		errors.Wrap(err, "repository.Index.Query couldn't load data from database")
	}
	todos := []entity.Todo{} // массив с элементами из бд
	for res.Next() {
		var post entity.Todo
		err = res.Scan(&post.ID, &post.Todo)
		if err != nil {
			errors.Wrap(err, "repository.Index.Scan couldn't scan data")
		}

		todos = append(todos, post) // Добавление элемента в массив
	}

	t.ExecuteTemplate(w, "index", todos) // Отправка массива в HTML для дальнейшей реализации

}

func DelTodo(w http.ResponseWriter, r *http.Request) { // Для удаления задачи
	todo := r.FormValue("TODO_CHECKBOX")            // Кликнутый элемент для удаления
	todoForUpdate := r.FormValue("UPDATE_CHECKBOX") // Кликнутый элемент для изменения
	NewTodo := r.FormValue("UPDATE_TODO")           // Новый текст для задачи

	db, _ := database.ConnectToDB()

	if todo != "" { // Если кликнут на удаление
		_, err := db.Query(fmt.Sprintf("DELETE FROM `todo` WHERE `Id`='%s'", todo))
		if err != nil {
			errors.Wrap(err, "repository.DelTodo.Query couldn't delete data from database")
		}
		NewID() // Функция для переопределения ID элементов в БД
	} else {
		if NewTodo != "" {
			_, err := db.Query(fmt.Sprintf("UPDATE `todo` SET `todo`='%s' WHERE `Id`='%s'", NewTodo, todoForUpdate)) // Запрос на изменение
			if err != nil {
				errors.Wrap(err, "repository.DelTodo.Query couldn't update data from database")
			}
		}

	}

	http.Redirect(w, r, "/", http.StatusSeeOther) // Переадресация на главную страницу
}

func NewID() { //Функция для переопределения ID в бд
	db, err := database.ConnectToDB()

	var size uint = 1 // переменная для установки новых ID

	res, err := db.Query("SELECT * FROM `todo` ORDER BY `todo`.`id` ASC")
	if err != nil {
		errors.Wrap(err, "repository.NewID.Query couldn't load data from database")
	}

	for res.Next() {

		var post entity.Todo
		err = res.Scan(&post.ID, &post.Todo)
		if err != nil {
			errors.Wrap(err, "repository.NewID.Scan couldn't scan data")
		}

		// Через цикл перебриаем все задачи и меняем их текущий айди на новый начиная с 1, затем инкриментируем переменную size
		db.Query(fmt.Sprintf("UPDATE `todo` SET `id`='%d' WHERE `Id`='%d'", size, post.ID))
		size++
	}

}

func AddTodo(w http.ResponseWriter, r *http.Request) { // Функция добавления задачи в БД

	todo := r.FormValue("NEW_TODO")

	if todo == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, err := database.ConnectToDB()

		var size uint = 0

		res, err := db.Query("SELECT * FROM `todo` ORDER BY `todo`.`id` ASC")
		if err != nil {
			errors.Wrap(err, "repository.AddTodo.Query couldn't load data from database")
		}

		for res.Next() {
			var post entity.Todo
			err = res.Scan(&post.ID, &post.Todo)
			if err != nil {
				errors.Wrap(err, "repository.Index.AddTodo couldn't scan data")
			}
			size = post.ID
			// Через цикл перебираю все задачи, чтобы узнать значение последнего ID
			// В запросе ниже добавляю новую запись с ID равному (последнему найденому+1)

		}

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `todo` (`id`,`todo`) VALUES ('%d','%s')", size+1, todo))

		if err != nil {
			errors.Wrap(err, "repository.AddTodo.Query couldn't add data from database")
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
