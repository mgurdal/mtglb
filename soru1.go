// Read fonksiyonu içerisinde dosyanın kapalı olup olmadığını
// kontrol edin. Dosya kapalı ise "file: file is closed"
// mesajı ile birlikte yeni bir hata döndürün.

package main

import "fmt"

type File struct {
	closed bool
}

func Read(file File) {}

func main() {
	f := File{closed: true}
	err := Read(f)
	fmt.Println(err)
}
