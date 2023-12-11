package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type CourseData struct {
	PublicName string
	Icon       string
	CourseCode string
	DocX       string
	PDF        string
}

const templateFilePath = "web/templates/sales_hub_template.html"
const remoteFilePath = "/opt/enchilada/run/static/outlines/sales_hub"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/courses", http.StatusFound)
}

func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch live courses data from the server
	liveCourses := getLiveCourses()

	// Render the "Courses Table" section template
	renderSectionTemplate(w, "Courses Table", liveCourses)
}

func SalesEnablementHandler(w http.ResponseWriter, r *http.Request) {
	// Render the "Sales Enablement" section template
	renderSectionTemplate(w, "Sales Enablement", nil)
}

func ProposalTemplatesHandler(w http.ResponseWriter, r *http.Request) {
	// Render the "Proposal Templates" section template
	renderSectionTemplate(w, "Proposal Templates", nil)
}

func PricingHandler(w http.ResponseWriter, r *http.Request) {
	// Render the "Pricing" section template
	renderSectionTemplate(w, "Pricing", nil)
}

func getLiveCourses() []CourseData {
	// Define the directory path
	coursesDir := "../labs/courses"

	// List all directories within "courses"
	courseDirs, err := os.ReadDir(coursesDir)
	if err != nil {
		fmt.Println("Error reading course directories:", err)
		return nil
	}

	// Create a slice to store information about live courses
	var liveCourses []CourseData

	// Iterate over each course directory
	for _, courseDir := range courseDirs {
		// Check if it's a directory
		if courseDir.IsDir() {
			courseName := courseDir.Name()

			// Read the content of the course.yml file
			ymlPath := filepath.Join(coursesDir, courseName, "course.yml")
			ymlContent, err := os.ReadFile(ymlPath)
			if err != nil {
				fmt.Println("Error reading course.yml for", courseName, ":", err)
				continue
			}

			// Extract relevant information from course.yml
			courseState := getValueFromYAML(ymlContent, "course_state")
			if courseState == "live" {
				marketingName := getValueFromYAML(ymlContent, "marketing_name")

				// Create CourseData struct
				courseData := CourseData{
					PublicName: marketingName,
					Icon:       fmt.Sprintf("https://static.alta3.com/outlines/%s/%s.png", courseName, courseName),
					CourseCode: courseName,
					DocX:       fmt.Sprintf("https://static.alta3.com/outlines/%s/%s.docx", courseName, courseName),
					PDF:        fmt.Sprintf("https://static.alta3.com/outlines/%s/%s.pdf", courseName, courseName),
				}

				// Append to the list of live courses
				liveCourses = append(liveCourses, courseData)
			}
		}
	}

	return liveCourses
}

func getValueFromYAML(ymlContent []byte, key string) string {
	lines := strings.Split(string(ymlContent), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key+":") {
			// Extract the value after the colon
			return strings.TrimSpace(strings.TrimPrefix(line, key+":"))
		}
	}
	return ""
}

func renderSectionTemplate(w http.ResponseWriter, sectionTitle string, courses []CourseData) {
	// Read the HTML template file
	templateFile, err := os.ReadFile(templateFilePath)
	if err != nil {
		fmt.Println("Error reading HTML template file:", err)
		http.Error(w, "Internal ServerError", http.StatusInternalServerError)
		return
	}

	// Create a new template and parse the HTML template
	tmpl, err := template.New("sales_hub").Parse(string(templateFile))
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a data structure to pass to the template
	data := struct {
		SectionTitle string
		Courses      []CourseData
	}{
		SectionTitle: sectionTitle,
		Courses:      courses,
	}

	// Execute the template with the provided data
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing HTML template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
