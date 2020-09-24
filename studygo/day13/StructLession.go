package main 
import(
	"fmt"
)

type Book struct {
	name string
	subject string
	author string
	no int
}

func main(){
	fmt.Println(Book{"go语言入门到精通","go 语言","Ruby",1})
	book :=Book{name:"go语言入门到精通"}
	fmt.Println(book.name)
}