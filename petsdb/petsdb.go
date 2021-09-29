package petsdb

import (

	"context"
	"fmt"
	"log"
	"os"
	"time"
	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
)

var projectID string

// Pet model stored in firestore
type Pet struct {
	Added   time.Time `firestore:"added"`
	Caption string    `firestore:"caption"`
	Email   string    `firestore:"email"`
	Image   string    `firestore:"image"`
	Likes   int       `firestore:"likes"`
	Owner   string    `firestore:"owner"`
	Petname string    `firestore:"petname"`
	Name    string    `firestore:"pets"`
}

// GetPets Returns all pets from firestore
func GetPets() ([]Pet, error) {

	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_CLOUD_PROJECT"`)
	}

	var pets []Pet
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create firestore client: %v", err)
	}

	// Create a query to fetch all Pet entities".
	var ordered_pets = client.Collection("pets").Documents(ctx)
	fmt.Println(ordered_pets)

	defer ordered_pets.Stop() //  make sure our resources get cleaned up
	for {
		doc, err := ordered_pets.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		var pet Pet
		if err := doc.DataTo(&pet); err != nil {
			fmt.Println(err)
    	}
   	 	pets = append(pets, pet)
	}

	// Set the id field on each Task from the corresponding key.

	client.Close()
	return pets, nil
}
func AddPets() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create firestore client: %v", err)
	}

	client.Collection("pets").Add(ctx, map[string]interface{}{
        "added":    time.Now(),
        "caption":   "pretty bryan",
        "email": "bryguyluvr@xvideos.com",
		"image": "img of bryan",
		"likes": "69",
		"owner": "bryan's wife",
		"petname": "bryan",
		"name": "pet",
		})
	if err != nil {
        log.Printf("An error has occurred: %s", err)
	}
}
