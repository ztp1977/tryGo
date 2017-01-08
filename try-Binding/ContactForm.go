package main

import (
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/mholt/binding"
)

type ContactForm struct {
	User struct {
		ID int
	}
	Email   string
	Message string
}

func (cf *ContactForm) FieldMap(req *http.Request) binding.FieldMap {
	pp.Println(req)
	return binding.FieldMap{
		&cf.User.ID: "user_id",
		&cf.Email:   "email",
		&cf.Message: binding.Field{
			Form:     "message",
			Required: true,
		},
	}
}

func (cf ContactForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if cf.Message == "Go needs generics" {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"message"},
			Classification: "ComplaintError",
			Message:        "Go has generics. They're called interfaces.",
		})
	}
	return errs
}

func handler(resp http.ResponseWriter, req *http.Request) {
	contactForm := new(ContactForm)
	errs := binding.Bind(req, contactForm)
	if errs.Handle(resp) {
		return
	}
	pp.Println(contactForm)
	fmt.Fprintf(resp, "From: %d\n", contactForm.User.ID)
	fmt.Fprintf(resp, "Message: %s\n", contactForm.Message)
}

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/contact", handler)
	fmt.Println("listen to 3000")
	http.ListenAndServe(":3000", nil)
}
