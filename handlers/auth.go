package handlers

import (
	"chitchat/models"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request)  {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(writer, nil)
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

func SignupAccount(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:	request.PostFormValue("name"),
		Email:	request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

func Authenticate(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}

		cookie := http.Cookie{
			Name:	"_cookie",
			Value:	session.Uuid,
			HttpOnly:true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "login", 302)
	}
}

func Logout(writer http.ResponseWriter, request *http.Request)  {
	cookie, err := request.Cookie("_cookie")
	info("coolkie.value:" , cookie.Value)
	if err != http.ErrNoCookie {
		session := models.Session{Uuid:cookie.Value}
		info("session:", session)
		err = session.DeleteByUUID()
		if err != nil {
			danger(err, "Failed")
		}
	} else {
		warning(err, "Failed to get cookie")
	}
	http.Redirect(writer, request, "/", 302)
}





