package main

import (
	"fmt"
)

func main() {
	// TODO: Create a NotificationBuilder and use it to set properities
	var bldr = newNotificationBuilder()

	// TODO: Use the builder to set some properities
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(10)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetNotType("alert")
	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Println("Notification: %+v", notif)
	}
}
