const clientId = "cbb509947a9d4764936944ffdca40529";
const clientSecret = "5c216a0ccd9c4d84905a95416b04179b";

async function getSpotifyToken() {
    // This function retrieves a Spotify token using the client ID and client secret.
    const response = await fetch("https://accounts.spotify.com/api/token", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "Authorization": "Basic " + btoa(`${clientId}:${clientSecret}`)
        },
        body: new URLSearchParams({
            "grant_type": "client_credentials"
        })
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data.access_token;
}

async function getArtistTopTrack(artistName) {
    //This function retrieves the top track for an artist
    const token = await getSpotifyToken();
    const searchUrl = `https://api.spotify.com/v1/search?q=${encodeURIComponent(artistName)}&type=artist`;
    const searchResponse = await fetch(searchUrl, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });
    const searchData = await searchResponse.json();
    const artistId = searchData.artists.items[0].id;

    const topTracksUrl = `https://api.spotify.com/v1/artists/${artistId}/top-tracks?market=US`;
    const topTracksResponse = await fetch(topTracksUrl, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });
    const topTracksData = await topTracksResponse.json();
    const topTrack = topTracksData.tracks[0];

    return topTrack;
}

async function searchArtist(artistName) {
    //This function searches for an artist and displays their top track
    const topTrack = await getArtistTopTrack(artistName);

    const playerDiv = document.getElementById('player');
    playerDiv.innerHTML = `
        <iframe src="https://open.spotify.com/embed/track/${topTrack.id}" class="spotify-embed"  width="1300" height="380" frameborder="0" allowtransparency="true" allow="encrypted-media"></iframe>`;
}

searchArtist(artistName);