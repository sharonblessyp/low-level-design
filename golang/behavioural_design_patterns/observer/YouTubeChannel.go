package observer

import "fmt"

type YouTubeChannel struct {
	subscribers map[string]Subscriber
}

func NewYoutubeChannel() *YouTubeChannel {
	return &YouTubeChannel{
		subscribers: make(map[string]Subscriber, 0),
	}
}

// Adds subscriber to the list
func (yc *YouTubeChannel) Subsribe(subscriber Subscriber) {
	if _, exists := yc.subscribers[subscriber.GetId()]; exists {
		return
	}
	yc.subscribers[subscriber.GetId()] = subscriber
}

// UnSubscribes from the list
func (yc *YouTubeChannel) UnSubsribe(subscriber Subscriber) {
	_, exists := yc.subscribers[subscriber.GetId()]
	if !exists {
		return
	}
	delete(yc.subscribers, subscriber.GetId())
}

// Upload video and notify subscriber
func (yc *YouTubeChannel) UploadVideo(videoTitle string) {
	fmt.Printf("uploaded video: %s \n", videoTitle)
	yc.Notify()
}

func (yc *YouTubeChannel) Notify() {
	for _, subscriber := range yc.subscribers {
		subscriber.UpdateSubscriber()
	}
}
