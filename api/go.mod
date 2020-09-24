module github.com/ks6088ts/spidey/api

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/google/uuid v1.1.2
	github.com/ks6088ts/spidey/todo v0.0.0-20200923090548-942d402c1ad1
	github.com/vektah/gqlparser/v2 v2.1.0
)

replace github.com/ks6088ts/spidey/todo v0.0.0-20200923090548-942d402c1ad1 => ../todo
