package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"html/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
	//return ioutil.WriteFile(filename, p.Body, os.ModeType.Perm())
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	p, err := loadPage(title)
	if err != nil {
		//w.Write([]byte(string(err.Error())))
		//fmt.Fprintf(w, "%s", err.Error())
		p = &Page{Title: title}
	}
	//fmt.Fprintf(w, "%s", p.Body)
	renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/edit/"):]
	//p, err := loadPage(title)
	//if err != nil {
	//	p = &Page{Title: title}
	//}
	//fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//	"<form action=\"/save/%s\" method=\"POST\">"+
	//	"<textarea name=\"body\">%s</textarea><br>"+
	//	"<input type=\"submit\" value=\"Save\">"+
	//	"</form>",
	//	p.Title, p.Title, p.Body)

	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//t, _ := template.ParseFiles("edit.html")
	//t.Execute(w, p)
	renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	//body_bytes := make([]byte, 1024)
	//_, err := r.Body.Read(body_bytes)
	//body := string(body_bytes)
	//if err != nil{
	//	fmt.Sprint(w, "wo read Request Body error: %s", err.Error())
	//
	//}
	//fmt.Sprint(w, "wo read Request.Body: %s", body)

	/* version 1 */
	//title := r.URL.Path[len("/save/"):]
	//body := r.FormValue("body")
	//p := &Page{Title: title, Body: []byte(body)}
	//p.save()
	//http.Redirect(w, r, "/view/"+title, http.StatusFound)


	/* version 2 */
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	/* version 1 */
	//t, _ := template.ParseFiles(tmpl + ".html")
	//t.Execute(w, p)

	/* version 2 */
	//t, err := template.ParseFiles(tmpl + ".html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//err = t.Execute(w, p)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	/* version 3 */
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func test1() {
	//p1 := &Page{"Test Page", []byte("This is a test page body")}
	//p1.save()
	//
	//p2, err := loadPage("Test Page")
	//if err != nil {
	//	fmt.Printf("error: %r", err)
	//	return
	//}
	//fmt.Println(string(p2.Body))
	//fmt.Println(p2.Body)
}
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
