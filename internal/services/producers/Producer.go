package producers

import "fmt"

type Producer interface {
	GetGorutines() int
}

type Generator struct {
}

func GetPorts() []string {
	return []string{
		"80", "280", "443", "488", "591", "593", "623", "664", "777", "832",
		"1128", "1129", "1183", "1184", "5000", "5001", "8008", "8080", "11371",
	}
}
func (g Generator) GetGorutines() int {
	var grt int

	fmt.Print("Выберите количество потоков: ")
	_, err := fmt.Scanln(&grt)
	if err != nil {
		panic(err)
	}
	return grt
}
