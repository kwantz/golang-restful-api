package web

type CategoryUpdateRequest struct {
	ID   int64  `validate:"required,numeric,min=1"`
	Name string `validate:"required,max=200,min=1"`
}
