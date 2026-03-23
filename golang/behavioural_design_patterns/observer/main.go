package observer

//  channel -   youtube, new blob post
//              a subscriber can
//              1. subsribe
//              2. unsubsribe
//              3. notify
// subscriber -  can be mail, push notifications
//              subscriber should maintain all subsribers list

func main() {
	emailSubscriber := NewEmailSubscriber("sharon@gmail.com")

	youtubeChannel := NewYoutubeChannel()
	youtubeChannel.Subsribe(emailSubscriber)
	youtubeChannel.UploadVideo("observer pattern")

	youtubeChannel.UnSubsribe(emailSubscriber)
	youtubeChannel.UploadVideo("observer pattern \n")

}
