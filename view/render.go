package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"git.urantiatech.com/auth/auth/constant"
	"git.urantiatech.com/auth/auth/model"
	"git.urantiatech.com/auth/auth/template/layout"
	"git.urantiatech.com/auth/auth/template/page"
)

// Templates with functions available to them
var base = template.New("auth")
var pages map[string]*template.Template

func init() {

	pages = make(map[string]*template.Template)

	_, err := base.Parse(layout.Default)

	if err != nil {
		fmt.Println(err.Error())
	}

	if pages[constant.Login], err = template.Must(base.Clone()).Parse(page.Login); err != nil {
		log.Fatal("Login:", err)
	}

	if pages[constant.Register], err = template.Must(base.Clone()).Parse(page.Register); err != nil {
		log.Fatal("Register:", err)
	}

	if pages[constant.Register1], err = template.Must(base.Clone()).Parse(page.Register1); err != nil {
		log.Fatal("Register1:", err)
	}

	if pages[constant.Register2], err = template.Must(base.Clone()).Parse(page.Register2); err != nil {
		log.Fatal("Register2:", err)
	}

	if pages[constant.Register3], err = template.Must(base.Clone()).Parse(page.Register3); err != nil {
		log.Fatal("Register3:", err)
	}

	if pages[constant.Reset], err = template.Must(base.Clone()).Parse(page.Reset); err != nil {
		log.Fatal("Reset:", err)
	}

}

func Render(w http.ResponseWriter, page string, data interface{}) error {
	if err := pages[page].Execute(w, data); err != nil {
		log.Fatal(err)
	}
	return nil
}
