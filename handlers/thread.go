package handlers

import (
	"chitchat/models"
	"fmt"
	"net/http"
)

func NewThread(writer http.ResponseWriter, request *http.Request)  {
	_, err := session(writer, request)
	if (err != nil) {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layput", "auth.navbar", "new.thread")
	}
}

func CreateThread(writer http.ResponseWriter, request *http.Request)  {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			fmt.Println("Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func ReadThread(writer http.ResponseWriter, request *http.Request)  {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println("Cannot read thread")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
