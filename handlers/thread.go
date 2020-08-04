package handlers

import (
	"chitchat/models"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

func NewThread(writer http.ResponseWriter, request *http.Request)  {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

func CreateThread(writer http.ResponseWriter, request *http.Request)  {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			warning(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		fmt.Println(topic)
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func ReadThread(writer http.ResponseWriter, request *http.Request)  {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	info("Read thread id :" , uuid)
	thread, err := models.ThreadByUUID(uuid)
	info(thread, err)
	msg := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "thread_not_found",
	})
	info(msg)
	if err != nil {
		error_message(writer, request, msg)
	} else {
		if thread.Id == 0 {
			error_message(writer, request, msg)
		}
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
