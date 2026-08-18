package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/shopspring/decimal"
)

var ErrProductNotFound = errors.New("product not found")

type ProductPackConfig struct {
	Packs []ProductPackSize `json:"packs"`
}

type ProductPackSize struct {
	Size int             `json:"size"`
	Cost decimal.Decimal `json:"cost"`
}

// Represents a product in the system
// Wouldn't usually directly expose DB types through an API - but I think a mapping layer (and added complexity) is outside the scope here
type Product struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	PackSizes []ProductPackSize `json:"packSizes"`
}

type ProductGetter interface { // online seems to lean towards an interface per method
	get() (*Product, error)
}

type ProductRepository struct {
}

func (p *ProductRepository) Get(id int) (*Product, error) {

	if id == 404 {
		return nil, ErrProductNotFound
	}

	// this would really be in the Db somewhere, per product, but i'm satisfying the requirement "Keep your application flexible so that pack sizes can be changed and added and removed without having to change the code"
	var packs, _ = getPacksFromFile() // discard the possibility of an error, but would error handle here in PROD

	return &Product{ID: 1, Name: "Super Fast Sonic Socks", PackSizes: packs}, nil
}

func getPacksFromFile() ([]ProductPackSize, error) {

	jsonFile, err := os.Open("packs.json") // TODO: Need to close the file, but I'm not 100% sure where this should happen
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// we initialize our Conf object
	var conf ProductPackConfig

	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return conf.Packs, nil
}
