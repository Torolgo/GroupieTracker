async function getCoordinates(cityName) {
    // This function retrieves the coordinates of a city
    const apiKey = '2cfc4e1bd5e4427b948d550ca541cf01';
    const url = `https://api.opencagedata.com/geocode/v1/json?q=${encodeURIComponent(cityName)}&key=${apiKey}`;
    //recovery of the coordinates of the city
    try {
        //answer recovery and storage
        const response = await fetch(url);
        const data = await response.json();

        //return of the coordinates of the city
        if (data.results.length > 0) {
            const location = data.results[0].geometry;
            return {latitude: location.lat, longitude: location.lng};
        } else {
            throw new Error("Geocoding error");
        }
    } catch (error) {
        console.error("Error when retrieving coordinates : ", error);
    }
}

async function cityLocation(citiesInGo) {
    //This function displays the location of the cities on the map
    let citie = citiesInGo.split("[")[1].split("]")[0];
    let cities = citie.split(" ");
    let cityWithoutCountry  = []
    for(let i = 0; i < cities.length; i++) {
        cityWithoutCountry.push(cities[i].split("-")[0]);
    }

    // Initialize the map
    let map = L.map('map').setView([51.5074456, -0.1277653], 1);

    //Importing the layers
    let MainLayer = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    let Stadia_AlidadeSatellite = L.tileLayer('https://tiles.stadiamaps.com/tiles/alidade_satellite/{z}/{x}/{y}{r}.{ext}', {
        minZoom: 0,
        maxZoom: 20,
        attribution: '&copy; CNES, Distribution Airbus DS, © Airbus DS, © PlanetObserver (Contains Copernicus Data) | ' +
            '&copy; <a href="https://www.stadiamaps.com/" target="_blank">Stadia Maps</a> ' +
            '&copy; <a href="https://openmaptiles.org/" target="_blank">OpenMapTiles</a> ' +
            '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
        ext: 'jpg'
    });

    let OPNVKarte = L.tileLayer('https://tileserver.memomaps.de/tilegen/{z}/{x}/{y}.png', {
        maxZoom: 18,
        attribution: 'Map <a href="https://memomaps.de/">memomaps.de</a> <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    });

    // Add a layer control
    L.control.layers({
        "Default view" : MainLayer,
        "Satellite view" : Stadia_AlidadeSatellite,
        "Public transport view" : OPNVKarte
    }).addTo(map)

    // Retrieve the coordinates of the cities with the getCoordinates function and display them on the map
    for (let city of cityWithoutCountry) {
        const coordinates = await getCoordinates(city);
        if (coordinates) {
            // Setting up markers
            let marker = L.marker([coordinates.latitude, coordinates.longitude]);
            marker.bindPopup(city).openPopup();
            marker.addTo(map);
            map.setView([coordinates.latitude, coordinates.longitude], 1);
        }
    }
}

cityLocation(cities);