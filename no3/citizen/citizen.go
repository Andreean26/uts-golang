package citizen

import (
    "encoding/json"
    "net/http"
    "strings"
)

type Person struct {
    Nama    string `json:"Nama"`
    Nik     string `json:"Nik"`
    Alamat  string `json:"alamat"`
}

var people =[]Person{
    {Nama: "Budi", Nik: "123456789", Alamat: "Jl. Sudirman"},
    {Nama: "Andi", Nik: "987654321", Alamat: "Jl. Thamrin"},
}

func FindPerson(w http.ResponseWriter, r *http.Request) {
	count := 0

	if r.Method == "POST" {
		getNamePerson := Person{
			Nama: r.FormValue("Nama"),
		}

		for _, value := range people {
			if strings.EqualFold(getNamePerson.Nama, value.Nama) {
				response, err := json.Marshal(Person{
					Nama:   value.Nama,
					Nik:    value.Nik,
					Alamat: value.Alamat,
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				count++

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}

		if count == 0 {
			http.Error(w, "Name not found", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func GetPeople(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}

