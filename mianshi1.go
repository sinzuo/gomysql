package main

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
