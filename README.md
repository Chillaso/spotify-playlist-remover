# Spotify playlist remover
Remove all tracks added by an user in a playlist

# Requirements
* Go
* (Spotify developer account)[https://developer.spotify.com/dashboard/login]
* Spotify project created

# Usage

```git
git clone https://github.com/Chillaso/spotify-playlist-remover.git
cd spotify-playlist-remover
```

Export spotify environment variables 
```bash
export SPOTIFY_ID=your_spotify_id
export SPOTIFY_SECRET=your_spotify_secret
```

```bash
go run . --user userID --playlist playlistID
```
