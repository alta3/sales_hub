package app

import (
	"html/template"
	"net/http"
)

// Define a struct for passing data to templates.
// This example is basic and can be expanded to include any data you want to display on your pages.
type PageData struct {
	Title string
}

// HomeHandler redirects to the courses page as per your original setup.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", "Courses")
}

// CoursesHandler for rendering the courses page.
func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "courses.html", "Courses")
}

// EventsHandler for rendering the events page.
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "events.html", "Events")
}

// AboutUsHandler for rendering the about us page.
func AboutUsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", "About Us")
}

// BlogHandler for rendering the blog page.
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "blog.html", "Blog")
}

// ContactUsHandler for rendering the contact page.
func ContactUsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html", "Contact Us")
}

// renderTemplate is a helper function to parse and execute templates with the base template.
func renderTemplate(w http.ResponseWriter, templateName string, title string) {
	tmpl, err := template.ParseFiles("web/templates/base.html", "web/templates/"+templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: title,
	}

	// Changed to Execute, which automatically executes the first parsed template.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
