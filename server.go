package main


import (

	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

)


type Page struct {

	Title string
	Body  []byte

}


func (p *Page) save() error {

	filename := p.Title + ".txt"
	fileContent, _ := ioutil.ReadFile(filename) 
	fmt.Println("Previous file state: ",  fileContent) 
	return ioutil.WriteFile(filename, p.Body, 0600)

}


func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("***************Error!", err)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println("***************Error!", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w,p)
	if err != nil {
		fmt.Println("***************Error!", err.Error())
        	http.Error(w, err.Error(), http.StatusInternalServerError)
    	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Println("***************Error!", err.Error())
        	http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
	}
	renderTemplate(w, "view", p)
}
/*
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	fmt.Println(body)

	err := p.save()
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("***************Error!", err.Error())
		return
	} else {
		fmt.Println("Changes saved: ", err)
	}
	http.Redirect(w, r, "view"+title, http.StatusFound)
}
*/
func main() {

	http.HandleFunc("/view/", viewHandler)
/*	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", editHandler)
*/	
	log.Fatal(http.ListenAndServe(":8080", nil))

}



