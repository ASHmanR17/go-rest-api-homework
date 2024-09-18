package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// getTasks Обработчик для получения всех задач из конечной точки '/tasks'
// возвращает все задачи, которые хранятся в мапе
func getTasks(w http.ResponseWriter, r *http.Request) {
	// сериализуем данные из слайса tasks
	resp, err := json.Marshal(tasks)
	//При ошибке сервер возвращает статус 500 Internal Server Error.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// в заголовок ответа записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываем сериализованные в JSON данные в тело ответа
	w.Write(resp)
}

// postTask Обработчик для отправки задачи на сервер в конечной точке '/tasks'
// принимает задачу в теле запроса и сохраняет ее в мапе
func postTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer

	// читаем тело запроса в буфер buf
	_, err := buf.ReadFrom(r.Body)
	//При ошибке сервер возвращает статус 400 Bad Request
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// десериализуем данные из буфера в структуру task
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Создаем в мапе новую запись задачи
	tasks[task.ID] = task

	// в заголовок ответа записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус 201 Created
	w.WriteHeader(http.StatusCreated)
}

// getOneTask Обработчик для получения задачи по ID из конечной точки '/tasks/{id}'
// возвращает задачу с указанным в запросе пути ID, если такая есть в мапе
func getOneTask(w http.ResponseWriter, r *http.Request) {
	//chi.URLParam возвращает в виде строки значение параметра из URL '/tasks/{id}'
	id := chi.URLParam(r, "id")

	//проверяем, есть ли ключ в мапе. В случае ошибки возвращаем статус 400 Bad Request.
	task, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	// сериализуем данные из структуры task
	resp, err := json.Marshal(task)
	//При ошибке сервер возвращает статус 400 Bad Request.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// в заголовок ответа записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываем сериализованные в JSON данные в тело ответа
	w.Write(resp)
}

// deleteTask Обработчик удаления задачи по ID из конечной точки '/tasks/{id}'
// удаляет задачу с указанным в запросе пути ID, если такая есть в мапе
func deleteTask(w http.ResponseWriter, r *http.Request) {
	//chi.URLParam возвращает в виде строки значение параметра из URL '/tasks/{id}'
	id := chi.URLParam(r, "id")

	//проверяем, есть ли ключ в мапе. В случае ошибки возвращаем статус 400 Bad Request.
	_, ok := tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	// удаляем ключ из мапы.
	delete(tasks, id)

	// в заголовок ответа записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)

}
