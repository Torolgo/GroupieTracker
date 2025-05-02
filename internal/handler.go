package internal

import (
	"Groupie_Tracker/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// This function was made to put the information of the structure on the home page
	artists := FetchArtists()

	if artists == nil {
		// If the artist is not found, it will return an error 503
		log.Println("Error 503 : Service Unavailable")
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusServiceUnavailable, Message: "Service Unavailable"}
		templErr.Execute(w, dataErr)
		return
	}

	pageData := pkg.ArtistPageData{
		Artists: artists,
	}

	templ, err := template.ParseFiles("./frontend/template/homePage.gohtml")
	if err != nil {
		// If the template is not found, it will return an error 404
		log.Println("Error 404 : Not Found", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusNotFound, Message: "Not Found"}
		templErr.Execute(w, dataErr)
		return
	}

	err = templ.Execute(w, pageData)
}

func CardPagehandler(w http.ResponseWriter, r *http.Request) {
	// This function takes the id of the url and put the artist who have this one on the page.
	artistID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(artistID)
	if err != nil {
		log.Println("Error 500 : Internal Server Error", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		templErr.Execute(w, dataErr)
		return
	}

	artists := GetArtistById(id)
	locations := FetchLocations(id)

	if artists == nil || locations == nil {
		// If the artist or the location is not found, it will return an error 503
		log.Println("Error 503 : Service Unavailable")
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusServiceUnavailable, Message: "Service Unavailable"}
		templErr.Execute(w, dataErr)
		return
	}

	pageData := pkg.ArtistPageData{
		Artists:   []pkg.Artists{*artists},
		Locations: *locations,
	}

	templ, err := template.ParseFiles("./frontend/template/template.gohtml")
	if err != nil {
		// If the template is not found, it will return an error 404
		log.Println("Error 404 : Not Found", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusNotFound, Message: "Not Found"}
		templErr.Execute(w, dataErr)
		return
	}

	err = templ.Execute(w, pageData)
	if err != nil {
		// If the template is not found, it will return an error 500
		log.Printf("Error 500 : Internal Server Error \n", err)
		log.Println("Error 500 : Internal Server Error", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		templErr.Execute(w, dataErr)
		return
	}

}

func HomePageSwitch(w http.ResponseWriter, r *http.Request) {
	// This function launch a function depending on the request.
	switch r.Method {
	case "GET":
		err := r.ParseForm()
		if err != nil {
			// If the form is not found, it will return an error 400
			log.Println("Error 400 : Bad Request", err)
			templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
			if err != nil {
				log.Println("Error 404 : Not Found", err)
				http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
				return
			}
			dataErr := pkg.Error{Code: http.StatusBadRequest, Message: "Bad Request"}
			templErr.Execute(w, dataErr)
			return
		}
		HandleAllFilters(w, r)
		break
	default:
		// If the method is not GET, it will return an error 405
		log.Println("Error 405 : Method Not Allowed")
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusMethodNotAllowed, Message: "Method Not Allowed"}
		templErr.Execute(w, dataErr)
		return
	}
}

func HandleAllFilters(w http.ResponseWriter, r *http.Request) {
	// This function takes all the filters, to return a template with all of them.
	artistID := strings.TrimPrefix(r.URL.Path, "/Card-Page/")
	if artistID == "0" {
		artistID = "1"
	}

	id, _ := strconv.Atoi(artistID)

	artists := FetchArtists()
	locations := FetchLocations(id)

	if artists == nil || locations == nil {
		// If the artist or the location is not found, it will return an error 503
		log.Println("Error 503 : Service Unavailable")
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusServiceUnavailable, Message: "Service Unavailable"}
		templErr.Execute(w, dataErr)
		return
	}

	artists = HandleCreationDateFilter(artists, r)
	artists = HandleMembersFilter(artists, r)

	HandleTemplate(w, artists, *locations)
}

func HandleCreationDateFilter(artists []pkg.Artists, r *http.Request) []pkg.Artists {
	// This function returns all the artists who are the same, or more years of creations than the choice of the user
	year := r.FormValue("year")
	yearInt, _ := strconv.Atoi(year)

	if year == "" {
		return artists
	}

	if yearInt < 1958 || yearInt > 2015 {
		return artists
	}

	var artistsFiltered []pkg.Artists
	for _, artist := range artists {
		if artist.CreationDate >= yearInt {
			artistsFiltered = append(artistsFiltered, artist)
		}
	}

	return artistsFiltered
}

func HandleMembersFilter(artists []pkg.Artists, r *http.Request) []pkg.Artists {
	// This function returns all the artists who are checked on the page
	urlMembers := r.Form["member"]
	if len(urlMembers) == 0 {
		return artists
	}

	var membersFiltered []pkg.Artists
	for _, artist := range artists {
		artistMember := len(artist.Members)
		for _, count := range urlMembers {
			countInt, err := strconv.Atoi(count)
			if err != nil {
				continue
			}
			if artistMember == countInt {
				membersFiltered = append(membersFiltered, artist)
				break
			}
		}
	}
	return membersFiltered
}

func HandleTemplate(w http.ResponseWriter, artists []pkg.Artists, locations pkg.Locations) {
	// This function takes all the function who are on "HandleAllFilters" and return a template with the new structure
	artistPageData := pkg.ArtistPageData{
		Artists:   artists,
		Locations: locations,
	}

	tmpl, err := template.ParseFiles("./frontend/template/homePage.gohtml")
	if err != nil {
		// If the template is not found, it will return an error 404
		log.Println("Error 404 : Not Found", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusNotFound, Message: "Not Found"}
		templErr.Execute(w, dataErr)
		return
	}

	err = tmpl.Execute(w, artistPageData)
	if err != nil {
		// If the template is not found, it will return an error 500
		log.Println("Error 500 : Internal Server Error", err)
		templErr, err := template.ParseFiles("./frontEnd/template/error.gohtml")
		if err != nil {
			log.Println("Error 404 : Not Found", err)
			http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
			return
		}
		dataErr := pkg.Error{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		templErr.Execute(w, dataErr)
	}
}
