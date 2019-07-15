package main

import ("fmt")

type API struct {a int}

func (api API) passByReference(){
	api.a = 5
	fmt.Println("inside the function api.a =", api.a)
	fmt.Println(&api.a)
	}

func (api *API) passByValue(){api.a = 5}

func main() {
	var api = API{0}
	
	api.passByReference()
	fmt.Println(api.a)
	
	api.passByValue()
	fmt.Println(api.a)
}