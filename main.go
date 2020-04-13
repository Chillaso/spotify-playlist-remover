package main

import (
	"flag"
	"log"

	"github.com/zmb3/spotify"
)

var (
	userToRemove string
	playlist     spotify.ID
)

func main() {
	setArgs()
	client := login()
	tracks := getTracksToRemoveByUser(client, playlist)
	removeTracksFromPlaylist(client, tracks)
}

func setArgs() {
	userFlag := flag.String("user", "", "ID of user who added tracks into the playlist and want to be removed")
	playlistFlag := flag.String("playlist", "", "playlist ID we want to use")

	flag.Parse()

	if *userFlag == "" || *playlistFlag == "" {
		log.Fatalln("Bad arguments, please use, --user to specify userID you want to search and remove all his added tracks" +
			" and --playlist to specify playlistID you want to manage")
	} else {
		userToRemove = *userFlag
		playlist = spotify.ID(*playlistFlag)
	}
}

func getTracksToRemoveByUser(client *spotify.Client, playlistID spotify.ID) []spotify.ID {

	tracks, err := client.GetPlaylistTracks(playlistID)
	if err != nil {
		log.Fatal(err)
	}

	tracksToRemove := make([]spotify.ID, 0, 600)

	for page := 1; ; page++ {
		for _, track := range tracks.Tracks {
			if track.AddedBy.ID == userToRemove {
				tracksToRemove = append(tracksToRemove, track.Track.ID)
			}
		}
		err = client.NextPage(tracks)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("There are %d tracks to remove", len(tracksToRemove))

	return tracksToRemove
}

func removeTracksFromPlaylist(client *spotify.Client, tracks []spotify.ID) {
	//Spotify can't handle more than 100 per request
	for i := 0; i < len(tracks); i += 100 {
		var track100 []spotify.ID
		if i+100 < len(tracks) {
			track100 = tracks[i : i+100]
		} else {
			track100 = tracks[i:]
		}
		snapshotID, err := client.RemoveTracksFromPlaylist(playlist, track100...)
		if err != nil {
			log.Fatalf(err.Error())
		} else {
			log.Printf("Success, deleted tracks, new snapshotID is: %s", snapshotID)
		}
	}

}
