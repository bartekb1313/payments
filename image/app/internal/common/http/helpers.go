package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"net/http"
)

func Populate(form interface{}, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	decoder := schema.NewDecoder()
	decoder.Decode(form, r.Form)
}

func Validate(form interface{}) map[string]string {
	listErrors := make(map[string]string)
	validate := validator.New()

	err := validate.Struct(form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for i := 0; i < len(errors); i++ {
			listErrors[errors[i].StructField()] = "Invalid data"
		}
	}

	return listErrors
}
