package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var file = "./Catalog.json"

func CatalogAdres() string {
	return file
}

func SaveCatalog() {
	rawDataOut, err := json.MarshalIndent(&Ws, "", "  ")
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(file, rawDataOut, 0)
	if err != nil {
		fmt.Println("Cannot write updated catalog file:", err)
	}
	fmt.Println("saved in ", file)
}

// upload in memory catalog from json file
func UploadCatalog(Ws *UW) (*UW, string) {
	ok := "ok"
	catalogFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		// create new if error
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ok = fmt.Sprintln(err)
		}
		if err := f.Close(); err != nil {
			ok = fmt.Sprintln(err)
		}
		catalog := Ws
		return catalog, ok
	}
	defer catalogFile.Close()
	jsonParser := json.NewDecoder(catalogFile)
	if err := jsonParser.Decode(&Ws); err != nil {
		ok = fmt.Sprintln(err)
	}
	fmt.Println("Catalog uploaded from: ", file)
	return Ws, ok
}
