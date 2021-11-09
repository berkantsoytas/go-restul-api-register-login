package main

import (
	"fmt"
	"net/http"

	. "./helper"
)

func main() {
	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()
	
	// Sign up
	/*
		URI: http://localhost:8080/signup
		Method : POST
		Body: x-www-form-urlencoded 
		Data example : username = berkantsoytas | password = 123456 | email = mailtoberkant@gmail.com
	*/
	mux.HandleFunc("/signup", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirmPassword")

		uNameCheck := IsEmpty(uName)
		uMailCheck := IsEmpty(email)
		uPwdCheck := IsEmpty(pwd)
		uPwdConfirmCheck := IsEmpty(pwdConfirm)

		if uNameCheck && uMailCheck && uPwdCheck && uPwdConfirmCheck {
			fmt.Fprintf(rw, "ErrorCode is -10 : There is empty data.")
			return
		}

		if pwd == pwdConfirm {
			// Database operation
			fmt.Fprintf(rw, "Successful sign up.")
		} else {
			fmt.Fprintf(rw, "ErrorCode is -11 : Password and confirm password is not same.")
		}

	})

	// Login
	mux.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		
		email = r.FormValue("email")
		pwd = r.FormValue("password")

		uMailCheck := IsEmpty(email)
		uPwdCheck := IsEmpty(pwd)

		if uMailCheck && uPwdCheck {
			fmt.Fprintf(rw, "ErrorCode is -10 : There is empty data.")
			return
		}

		dbPwd := "12345!*."
		dbEmail := "mailtoberkant@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(rw, "Successful login.")
		} else {
			fmt.Fprintln(rw, "ErrorCode is -11 : Email or password is not correct.")
		}
	})

	http.ListenAndServe(":8080", mux)
}