package main 
import (
	"fmt"
	"errors"
	"net/http"
	"ioutil"
	"encoding/json"
)
//PKG.GO.DEV/NET/HTTP Implementacion del paquete

type Categories []Category

type Category struct{
	
     ID string `json:"id"`
	 Name string `json: "name"`
}

func main (){
	
	cats, err := GetCategories("MLA")
	if err != nil {
		//validar
	}
	fmt.Println("Las categorias son ...")
}

func GetCategories(siteID string)(Categories, error) {
   
	response := http.Get //Completar

	bytes := ioutil.Readall(response.Bytes)//completar agregar el error y hacer el if 

	var  cats Categories
	json.Unmarshal(bytes, &cats)

	return cats, nil


}