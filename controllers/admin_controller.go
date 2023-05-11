package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go_warung-laundry/config"
	"go_warung-laundry/models"

	"github.com/gorilla/mux"
)

type AdminController struct{}

//method untuk mengelola data user

// Method untuk menambah data user ke dalam database
func (c *AdminController) AddUser(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

// Method untuk mengupdate data user di dalam database
func (c *AdminController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID user dari parameter request
	params := r.URL.Query()
	id := params.Get("id")

	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	// decode data JSON dari request ke dalam struct User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result = db.Save(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

// Method untuk menghapus data user di dalam database
func (c *AdminController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID user dari parameter request
	params := r.URL.Query()
	id := params.Get("id")

	var user models.User
	result := db.Delete(&user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Method untuk menambah data pelanggan ke dalam database
func (c *AdminController) AddPelanggan(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	var pelanggan models.Pelanggan
	err := json.NewDecoder(r.Body).Decode(&pelanggan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result := db.Create(&pelanggan)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data pelanggan berhasil ditambahkan"))
}

// Method untuk mengupdate data pelanggan di dalam database
func (c *AdminController) UpdatePelanggan(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID pelanggan dari parameter request
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	var pelanggan models.Pelanggan
	result := db.First(&pelanggan, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&pelanggan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result = db.Save(&pelanggan)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data pelanggan berhasil diupdate"))
}

// Method untuk menghapus data pelanggan dari database
func (c *AdminController) DeletePelanggan(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID pelanggan dari parameter request
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	var pelanggan models.Pelanggan
	result := db.Delete(&pelanggan, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data pelanggan berhasil dihapus"))
}

// Method untuk mendapatkan seluruh data pelanggan dari database
func (c *AdminController) GetAllPelanggan(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	var pelanggans []models.Pelanggan
	result := db.Find(&pelanggans)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	response, _ := json.Marshal(pelanggans)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Method untuk menambah data jenis laundry ke dalam database
func (c *AdminController) AddJenisLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	var jenisLaundry models.JenisLaundry
	err := json.NewDecoder(r.Body).Decode(&jenisLaundry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result := db.Create(&jenisLaundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&jenisLaundry)
}

// Method untuk mengupdate data jenis laundry di dalam database
func (c *AdminController) UpdateJenisLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID jenis laundry dari parameter request
	params := r.URL.Query()
	id := params.Get("id")

	var jenisLaundry models.JenisLaundry
	result := db.First(&jenisLaundry, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&jenisLaundry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result = db.Save(&jenisLaundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&jenisLaundry)
}

// Method untuk menghapus data jenis laundry dari database
func (c *AdminController) DeleteJenisLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID jenis laundry dari parameter request
	params := r.URL.Query()
	id := params.Get("id")

	var jenisLaundry models.JenisLaundry
	result := db.First(&jenisLaundry, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
		return
	}

	result = db.Delete(&jenisLaundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&jenisLaundry)
}

// Method untuk menambah data laundry ke dalam database
func (c *AdminController) AddLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	var laundry models.Laundry
	err := json.NewDecoder(r.Body).Decode(&laundry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result := db.Create(&laundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&laundry)
}

// Method untuk mengupdate data laundry di dalam database
func (c *AdminController) UpdateLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID laundry dari parameter
	params := mux.Vars(r)
	laundryID := params["id"]

	var laundry models.Laundry
	result := db.First(&laundry, laundryID)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Laundry Not Found"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&laundry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	result = db.Save(&laundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&laundry)
}

// Method untuk menghapus data laundry dari database
func (c *AdminController) DeleteLaundry(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

	// mendapatkan ID laundry dari parameter
	params := mux.Vars(r)
	laundryID := params["id"]

	var laundry models.Laundry
	result := db.First(&laundry, laundryID)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Laundry Not Found"))
		return
	}

	result = db.Delete(&laundry)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Laundry has been deleted"))
}
