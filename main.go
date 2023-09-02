// main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("secret-key"))
	users = make(map[string]string) // Simple in-memory user store (username:password)
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/signup", signupHandler)
	r.HandleFunc("/dashboard", dashboardHandler)
	r.HandleFunc("/logout", logoutHandler)

	// Serve static files under the "/static/" URL path
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"].(string)

	if ok && username != "" {
		// User is already logged in, redirect to dashboard
		http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		// In a real application, you would validate the username and password here.
		// For simplicity, we're using a basic in-memory user store.

		if storedPassword, ok := users[username]; ok && storedPassword == password {
			session.Values["username"] = username
			session.Save(r, w)

			// Redirect to the dashboard with the username as a query parameter
			http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
			return
		}
	}

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"].(string)

	if ok && username != "" {
		http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		// In a real application, you would validate the username and password here.
		// For simplicity, we're using a basic in-memory user store.
		// Ensure the username is unique before creating the account.

		if _, exists := users[username]; exists {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		users[username] = password
		session.Values["username"] = username
		session.Save(r, w)

		// Redirect to the dashboard with the username as a query parameter
		http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["username"] = ""
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
