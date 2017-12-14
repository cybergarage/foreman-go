![foreman_logo](./img/icon.png)

# Coding Guidelines

To develop `foreman-go`, you must follow the following coding guidelines basically.

- [Effective Go](https://golang.org/doc/effective_go.html#interface-names)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Static Analyzer

In addition to the above standard guideline, you must use the following static analysers and fix the all warnings 

- [go vet](https://golang.org/cmd/vet/)
- [golint](https://github.com/golang/lint)

## Extra Coding Guidelines

In addition to the above standard guideline, you must follow the following extra rules too.

### Interface names

[Effective Go](https://golang.org/doc/effective_go.html#interface-names) specifies only an interface naming rule when the interface has only a method. 
Multiple-method interfaces are named by the class name plus an `-ing` suffix.
