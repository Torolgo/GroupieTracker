package pkg

// This file contains the structure of the data that will be used in the application
type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ArtistPageData struct {
	Artists   []Artists `json:"artists"`
	Locations Locations `json:"locations"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
