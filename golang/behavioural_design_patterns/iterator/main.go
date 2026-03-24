package iterator

import "fmt"

/*
1. iterator iterface - define necessary methods supported by iterator
2. concrete implementation of iterator
3. collection interface
4. concrete implementation of collection(youtube playlist)
5. how do we acheive contract with iterator interface?
   - collection interface must depend on iterator interface
*/

func main() {
	// create a video
	video := NewVideo("iterator pattern\n")

	// creta a playlist and add video to it
	yoytubePlayList := NewYouTubePlayList()
	yoytubePlayList.AddVideoToPlaylist(video)

	// to iterate the youtube playlist, create an iterator and loop through
	iterator := yoytubePlayList.CreateIterator()
	if iterator.HasNext() {
		video := iterator.Next()
		if video != nil {
			fmt.Println(video.GetTitle())
		}
	}

}
