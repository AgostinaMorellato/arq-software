import (
	"fmt"
	"errors"
)

func main (){
	
	div, err := division(12, 6)

	if err != nil {
		
		fmt.Println("Error:", err.Error())
		
		return 
	}
   fmt.Println("Division:", div)
}

func division (a , b int )(int, error){
	if b == 0 {
		return -1, errors.New("No se puede dividir por 0")
	}
     
	 return  a / b, nil

}