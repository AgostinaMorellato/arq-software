package div

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name           string
	A              float64
	B              float64
	ExpectedResult float64
	ExpectedError  bool
}

func TestDivision(t *testing.T) {
	for _, testCase := range []TestCase{
		{
			Name:           "8/4 OK",
			A:              8,
			B:              4,
			ExpectedResult: 2,
			ExpectedError:  false,
		},
		{
			Name:           "9/2 OK",
			A:              9,
			B:              2,
			ExpectedResult: 4.5,
			ExpectedError:  false,
		},
		{
			Name:          "9/0 Error",
			A:             9,
			B:             0,
			ExpectedError: true,
		},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			result, err := Division(testCase.A, testCase.B)
			if testCase.ExpectedError {
				assert.NotNil(t, err)
				return
			}
			assert.EqualValues(t, testCase.ExpectedResult, result)
		})
	}
}

/*result, err := Division (8,4)
	_, err := Division(8, 4)

	if err != nil {
		t.Error(err)
		return
	}
	assert.Nil(t, err)
	// if result !=2 {
	//t.Error(err)
	//return }
	assert.EqualValues(t, result, 2)
	for 1
    for i:= 0; i<= 10; i++{
		fmt.Println(i)
	}
	//for 2 
	values := []int{1,2,3,4,5,6}
	for _, v := range values {
		fmt.Println(v)
	}*/
/*func TestDivisionErrorByZero(t *testing.T) {
  
	_, err := Division(8, 0)

	if err == nil {
		t.Error(err)
		return
	}

}*/

