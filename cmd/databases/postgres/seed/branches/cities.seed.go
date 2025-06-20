package branches

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal/models"
	"main/internal/storage"
	"os"
	"strconv"
)

type CitiesEntry struct {
	Data []models.Cities `json:"data"`
}

func Cities() {
	jsonFile, err := os.Open("cmd/databases/postgres/seed/branches/cities.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var CitySeed map[string]interface{}
	json.Unmarshal(byteValue, &CitySeed)
	// fmt.Print("RawRow", CitySeed)

	for _, raw := range CitySeed["data"].([]interface{}) {
		RawRow := raw.(map[string]interface{})

		Lat := ""
		Lng := ""
		Streets_count := ""

		if RawRow["lat"] != nil {
			Lat = fmt.Sprintf("%f", RawRow["lat"].(float64))
		}
		if RawRow["lng"] != nil {
			Lng = fmt.Sprintf("%f", RawRow["lng"].(float64))
		}
		if RawRow["streets_count"] != nil {
			Streets_count = fmt.Sprintf("%f", RawRow["streets_count"].(float64))
		}
		Streets, _ := strconv.Atoi(Streets_count)

		row := models.Cities{
			Display_name_in: RawRow["display_name_in"].(string),
			Display_name:    RawRow["display_name"].(string),
			Streets_count:   Streets,
			Lng:             Lng,
			Lat:             Lat,
		}
		// fmt.Print(RawRow)
		storage.DB.Create(&row)

		for _, distraw := range RawRow["districts"].([]interface{}) {
			dist := distraw.(map[string]interface{})
			distLat := ""
			distLng := ""
			distStreets_count := ""

			if dist["lat"] != nil {
				distLat = fmt.Sprintf("%f", dist["lat"].(float64))
			}
			if dist["lng"] != nil {
				distLng = fmt.Sprintf("%f", dist["lng"].(float64))
			}
			if dist["streets_count"] != nil {
				distStreets_count = fmt.Sprintf("%f", dist["streets_count"].(float64))
			}
			distStreets, _ := strconv.Atoi(distStreets_count)
			distrow := models.Districts{
				Display_name_in: dist["display_name_in"].(string),
				Display_name:    dist["display_name"].(string),
				Streets_count:   distStreets,
				Lng:             distLng,
				Lat:             distLat,
				CityID:          int(row.ID),
			}
			storage.DB.Create(&distrow)
		}

	}
}
