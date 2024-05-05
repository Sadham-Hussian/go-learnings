package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

var meilisearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   "http://localhost:7700",
	APIKey: "aSampleMasterKey",
})

type IMeilisearchIndex interface {
	GetIndexName() string
}

type MovieIndex struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Cast   []string `json:"cast"`
	Genres []string `json:"genres"`
}

func (i *MovieIndex) GetIndexName() string {
	return "movies"
}

type MusicIndex struct {
	ID               int     `json:"id"`
	FirstAirDate     string  `json:"first_air_date"`
	Name             string  `json:"name"`
	OriginalLanguage string  `json:"original_language"`
	Popularity       float32 `json:"popularity"`
	VoteAverage      float32 `json:"voteAverage"`
	VoteCount        int     `json:"vote_count"`
}

func (i *MusicIndex) GetIndexName() string {
	return "musics"
}

func CreateIndex(indexName string) error {
	_, err := meilisearchClient.CreateIndex(&meilisearch.IndexConfig{
		Uid:        indexName,
		PrimaryKey: "id",
	})
	if err != nil {
		return err
	}
	log.Println("Index created : ", indexName)
	return nil
}

func DeleteIndex(indexName string) error {
	_, err := meilisearchClient.DeleteIndex(indexName)
	if err != nil {
		return err
	}
	log.Println("Index deleted : ", indexName)
	return nil
}

func AddDocuments[T any](indexName string, data []T) error {
	_, err := meilisearchClient.Index(indexName).AddDocuments(&data, "id")
	if err != nil {
		return err
	}
	return nil
}

func Search(indexName string) error {
	resp, err := meilisearchClient.Index("movies").Search("Water", &meilisearch.SearchRequest{
		Limit:  5,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	ToJson("HITS : ", resp.Hits)

	resp, err = meilisearchClient.Index("musics").Search("love", &meilisearch.SearchRequest{
		Limit:  5,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	ToJson("HITS : ", resp.Hits)

	_, err = meilisearchClient.Index("movies").UpdateFilterableAttributes(&[]string{"cast", "genres"})
	if err != nil {
		panic(err)
	}

	resp, err = meilisearchClient.Index("movies").Search("", &meilisearch.SearchRequest{
		Limit:  10,
		Offset: 0,
		Filter: `genres = Adventure AND cast = "Robert Downey Jr."`,
	})
	if err != nil {
		panic(err)
	}
	ToJson("HITS : ", resp.Hits)

	return nil
}

func main() {

	// Reading from json
	var movies []MovieIndex
	file, err := os.Open("movies.json")
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&movies); err != nil {
		panic(err)
	}
	for i := 0; i < len(movies); i++ {
		movies[i].ID = i + 1
	}

	var musics []MusicIndex
	file, err = os.Open("music.json")
	if err != nil {
		panic(err)
	}
	jsonParser = json.NewDecoder(file)
	if err = jsonParser.Decode(&musics); err != nil {
		panic(err)
	}
	for i := 0; i < len(musics); i++ {
		musics[i].ID = i + 1
	}

	// Create index
	err = CreateIndex("movies")
	if err != nil {
		panic(err)
	}

	err = CreateIndex("musics")
	if err != nil {
		panic(err)
	}

	// Add Documents
	err = AddDocuments("movies", movies)
	if err != nil {
		panic(err)
	}

	err = AddDocuments("musics", musics)
	if err != nil {
		panic(err)
	}

	// Search Documents
	Search("movies")

	// Delete index
	err = DeleteIndex("movies")
	if err != nil {
		panic(err)
	}

	err = DeleteIndex("musics")
	if err != nil {
		panic(err)
	}

}

func ToJson(message string, obj any) {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println("error converting to json : ", err)
	}
	log.Println(message, ": ", string(b))
}
