package requests

import (
	_ "fmt"
	"log"
	_ "log"

	"github.com/imroc/req"
	_ "github.com/imroc/req"
)

func GetDeployment() {
	r, err := req.Get("https://gorest.co.in/public/v1/posts")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", r)
}
func ListDeployment() {
}
func PatchDeployment() {
}
