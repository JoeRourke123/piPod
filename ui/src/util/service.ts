export const fetchAuthStatus = async () => {
    const response = await fetch(
        `http://localhost:9091/isAuth`
    );

    return await response.json()
}

export const fetchAlbums = async (albumId?: string, offset?: number) => {
    const response = await fetch(
        `http://localhost:9091/albums/${albumId}?next=${offset}`
    );

    return await response.json();
}

export const fetchPlaylist = async (playlistId?: string, offset?: number) => {
    const response = await fetch(
        `http://localhost:9091/playlists/${playlistId}?next=${offset}`
    );

    return await response.json();
}

export const fetchSpotifyToken = async () => {
    const responseJson = await fetchAuthStatus();

    return responseJson["access_token"];
}