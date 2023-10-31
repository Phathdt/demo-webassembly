package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

type Keygen struct {
	PrivateKey1 string
	PrivateKey2 string
}

func keygen() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 0 {
			return "Invalid no of arguments passed"
		}

		k := Keygen{"asasasas", "sasasasa"}

		// return k.PrivateKey1, k.PrivateKey2
		m, err := StructToMap(k)
		if err != nil {
			return err
		}

		return js.ValueOf(m)
	})

	return jsonFunc
}
func signTransaction() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		return args[0]
	})

	return jsonFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("keygen", keygen())
	js.Global().Set("signTransaction", signTransaction())
	<-make(chan bool)
}

func StructToMap(t interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
