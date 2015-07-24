package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type Recipe struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

var (
	recipeStorage = make(map[string]*Recipe)
)

func main() {
	goji.Get("/recipes/:id", getRecipes)
	goji.Get("/recipes", getAllRecipes)
	goji.Delete("/recipes/:id", deleteRecipe)
	goji.Put("/recipes", putRecipe)

	goji.Serve()
}

func getRecipes(c web.C, w http.ResponseWriter, req *http.Request) {
	rec := recipeStorage[c.URLParams["id"]]
	data, _ := json.Marshal(rec)
	w.Write(data)
}

func getAllRecipes(c web.C, w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(recipeStorage)
	w.Write(data)
}

func deleteRecipe(c web.C, w http.ResponseWriter, req *http.Request) {
	recipeStorage[c.URLParams["id"]] = nil
	data, _ := json.Marshal(recipeStorage)
	w.Write(data)
}

func putRecipe(c web.C, w http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	rec := new(Recipe)
	json.Unmarshal(data, rec)
	recipeStorage[rec.Id] = rec
	data, _ = json.Marshal(rec)
	w.Write(data)
}
