package internal

import (
	"Groupie_Tracker/pkg"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func FetchArtists() []pkg.Artists {
	// This function will fetch the artists
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}

	var artists []pkg.Artists
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	return artists
}

func GetArtistById(id int) *pkg.Artists {
	// This function will fetch the artist by id
	url := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	var artist pkg.Artists
	err = json.Unmarshal(body, &artist)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	return &artist
}

func FetchLocations(id int) *pkg.Locations {
	// This function will fetch the locations of the artist
	url := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error(), err)
		return nil
	}

	var locations pkg.Locations
	err = json.Unmarshal(body, &locations)

	return &locations
}
