package iterator

import "fmt"

type YouTubePlayList struct {
	playlist []*Video
}

// create an empty playlist
func NewYouTubePlayList() *YouTubePlayList {
	return &YouTubePlayList{playlist: []*Video{}}
}

func (yp *YouTubePlayList) AddVideoToPlaylist(video *Video) {
	fmt.Printf("Added video to the playlist %s", video.title)
	yp.playlist = append(yp.playlist, video)
	return
}

// return iterator to the playlist
func (yp *YouTubePlayList) CreateIterator() Iterator {
	return NewYoutubePlayListIterator(yp.playlist)
}
