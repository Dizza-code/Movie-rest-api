package models

import (
	"context"
	"fmt"
	"log"

	"example.com/movies-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie  string             `json:"movie"`
	Actors []string           `json:"actors"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}

// inserting a movie into the collection
func InsertMovie(movie Movie) error {
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	//passing the context and the movie object
	inserted, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return err
	}
	fmt.Println("inserted a record with id:", inserted.InsertedID)
	return nil
}

// insert many movies
// This function accepts a slice of movies
func InsertMany(movies []Movie) error {
	//convert to a slice of interface
	newMovies := make([]interface{}, len(movies))
	for i, movie := range movies {
		newMovies[i] = movie
	}
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	//passing the context and the movie object
	result, err := collection.InsertMany(context.TODO(), newMovies)
	if err != nil {
		panic(err)
	}
	log.Println(result)
	return err
}

// update a movie record, the  functions takes the moveID and the update movieRecords
func UpdateMovie(movieId string, movie Movie) error {
	// first we covert the ID into a primitive.ID this is the bson id that mongo understands
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return err
	}
	//create a filter to find the document with the matching underscore id
	//the filter specifies which document to update
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"movie": movie.Movie, "actors": movie.Actors}}

	//we get the collection then call the updateOne method on it
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Print("New record: ", result)
	return nil
}

// deleting a movie
// we covert the movie ID to an objectID and create a filter
func DeleteMovie(movieId string) error {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Delete result: ", result)
	return nil
}

// find a movie by name
func Find(movieName string) Movie {
	//we define a vairiable to hold the response
	var result Movie

	//we define a filter using bson.d which is a slice of key value pairs for queries
	filter := bson.D{{"movie", movieName}}
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	//using FindOne we fetch the first document that matches the query
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
func GetMovieByID(movieId string) (*Movie, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return nil, err
	}
	var movie Movie
	//filter to look for ID
	filter := bson.M{"_id": id}
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	err = collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return &movie, err
}

// fetch multiple records matching an attribute
// this function returns a slice of movies
func FindAll(movieName string) []Movie {
	var results []Movie
	filter := bson.D{{"movie", movieName}}
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	//we can cursor.all to decode all results into a slice of movie structs
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

// Now to list all movies, we just replace th filter with an empty filter
func ListAll() []Movie {
	var results []Movie
	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//we can cursor.all to decode all results into a slice of movie structs
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

// Delete All movies
func DeleteAll() error {

	collection := db.MongoClient.Database(db.Db).Collection(db.CollName)
	delResult, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Records deleted: ", delResult.DeletedCount)
	return err
}
