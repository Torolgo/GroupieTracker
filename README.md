# Groupie Tracker Project

## Introduction
Welcome to the Groupie Tracker project! The aim of this project is to understand the API principle and its management.

### Explanations:
"***The main objective is to retrieve detailed information about an artist using an API. This information will then be displayed dynamically on one or more web pages. This includes details such as artist, members, debut album, logo and other relevant information, offering an enriched and interactive user experience.***"

## Features
- A responsive design that adapts to any screen size!
- An artist search system 
- Filters to sort artists by creation date or number of members 
- A search bar to find a specific artist 
- An interactive map to visualize the locations of artists' concerts 
- A pagination system to display artists by page 
- An error management system for cases where the artist is not found

## Dépendance:
- [Groupie Tracker API](https://groupietrackers.herokuapp.com/api)
- [Maps API](https://leafletjs.com/examples/quick-start/)
- [More Layers](https://leaflet-extras.github.io/leaflet-providers/preview/)
- [Geocoding API](https://opencagedata.com/api)

## Installation

### Prerequisites
- Go (version 1.23 or higher)

### Steps
1. Clone the repository:
    ```sh
    git clone https://ytrack.learn.ynov.com/git/rlena/Groupie_Tracker.git
    ```
2. Move to the project directory:
    ```sh
    cd ./Groupie_Tracker/cmd
    ```
3. Run the project:
    ```sh
    go run main.go
    ```
4. Open your browser and go to the following URL: ``` http://localhost:8080/ ```

## Project Structure
All the project files are organized as follows:
```
Groupie_Tracker/
├── go.mod                               # Go module configuration file
├── cmd 
│   └── main.go                          # Main execution file
├── frontEnd/
│   ├── static/                          # Contains all the static files needed for the website  
│   │   ├── css/                         # Contains all the style files for the website 
│   │   ├── fonts/                       # Contains all the font files 
│   │   ├── images/                      # Contains all the images displayed on the website 
│   │   └── js/                          # Contains all the JavaScript execution files
│   └── template/                        # Contains all the templates and web pages
│   │   ├── homePage.gohtml              # Template for the Home page
│   │   ├── errorPage.gohtml             # Template for the Error page
│   │   └── template.gohtml              # Template for the artist page, adaptive to the specific artist
├── internal/
│   ├── fetch_api.go                     # This file contains functions to fetch data from the API
│   ├── handler.go                       # This file contains functions to handle requests made to the web server
│   └── server.go                        # This file sets up the web server routes and starts the server
├── pkg/
│   └── structure.go                     # This file contains the structure used for fetching data from the API
├── README.md                            # Project documentation
└── .gitignore                           # Defines what should be ignored by git
```
Project management is done on [Trello](https://trello.com/invite/b/6760358951bc5bc44434f53b/ATTI2461d50124896c73f2389845cbd5bb70DEED55E1/groupie-tracker)<br>
The presentation of the project is done on [Canva](https://www.canva.com/design/DAGfApwZwcw/ELXT3Qu_1UqWFDYgms1TIA/edit?utm_content=DAGfApwZwcw&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)<br>
The repository of the project is on [Gitea](https://ytrack.learn.ynov.com/git/rlena/Groupie_Tracker.git) 

## Contributeurs:
* [DESSENNE Ylan](https://ytrack.learn.ynov.com/git/dylan) - Developer
* [RICARD Léna](https://ytrack.learn.ynov.com/git/rlena) - Developer