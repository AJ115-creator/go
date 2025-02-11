_1.)_  TO run go file:
go build __FileName__.go
Just to see out of go file:
*cmd* - go run __Filename__.go

_2.)_ When we Declare a Variable in _GO_ we've to use it and if
unused , delete it .
_2.1.)_ Data Types:
       i.) int 
       ii.)float32
       iii.)float64
       iv.) string
       v.) bool

_3.)_ No need to define datatype while declaring a variable, it can
be automatically inferred in the prog.
different ways of assigning values:
       i) var name string = "hello"
       ii) var name = "goLang"
       iii) name := "Ayush"

_4.)_ Use "const" keyword to declare a constant
_5.)_ Outside the main func we can decalare variable only by 2 ways
     i.) const name = "hello"
     ii.) const name string = "hello"
     _NOT this_ **name:="hello"**  

_6.)_ Grouping Constants
	Eg: - const (
		host = "localhost"
		port = 8080
	)

_7.)_ 