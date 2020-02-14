package buildcsv

import (
	"context"
	"encoding/csv"
	"io"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func buildcsv() {
	var lignes = loadLignes()
	var arrets = loadArrets()
	var courses = loadCourses()

	var horaire horaire

	/*
		foreach horaires as horaire {
			horaire = horaire{

			}

			horaire.course = courses[horaire.RS_SV_COURS_A]
			horaire.arret = arrets[horaire.RS_SV_ARRET_P]
		}
	*/
}

func initializeBucket(opt option.ClientOption) *storage.BucketHandle {

	config := &firebase.Config{ProjectID: "bmtimetable", StorageBucket: "bmtimetable.appspot.com"}
	app, err := firebase.NewApp(context.Background(), config, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()
	storage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("error initializing storage: %v\n", err)
	}

	bkt, err := storage.DefaultBucket()
	if err != nil {
		log.Fatalf("error initializing bucket: %v\n", err)
	}

	return bkt
}

func loadDataIntoMemory(file io.Reader) map[string][]string {
	memory := make(map[string][]string)

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		memory[record[0]] = record
	}

	return memory
}
