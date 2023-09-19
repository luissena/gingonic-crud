package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s is required and must be %s", name, typ)
}

type CreatePostRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (r *CreatePostRequest) Validate() error {

	if r.Author == "" && r.Title == "" && r.Body == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Author == "" {
		return errParamIsRequired("author", "string")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Body == "" {
		return errParamIsRequired("body", "string")
	}
	return nil
}

type UpdatePostRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (r *UpdatePostRequest) Validate() error {

	if r.Author != "" || r.Title != "" || r.Body != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
