package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"time"
	"math/rand"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page,error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if(err != nil) {
		return nil, err
	}
	return &Page{Title : title, Body : body}, nil
}

func writeAndReadFile() {
	p1 := &Page{Title : "TestFile", Body : []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestFile")
	fmt.Println(string(p2.Body))
}

const lenPath = len("/view/")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func viewHandler(writer http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[lenPath:]
	p, _ := loadPage(title)
	renderTemplate(writer, "view", p)
}

func editHandler(writer http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[lenPath:]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(writer, "edit", p)
}

func _main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil);
}

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}
func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
