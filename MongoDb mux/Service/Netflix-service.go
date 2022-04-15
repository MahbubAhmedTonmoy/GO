package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NetflixService interface {
	Insert(model interface{})
	//Insert(movie model.Netflix)
	Update(movieId string)
	Delete(movieId string)
	DeleteAll() int64
	GetAllMovie() []primitive.M
	GetMovie(movieId string) primitive.M
	Count(filteringQuery interface{}) (int64, error) 
}

type netflixService struct {
	collection *mongo.Collection
}

func NewNetflixService(ConnectionString string, dbName string, collectionName string) NetflixService {
	clientOption := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection_ := client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance is ready")
	return &netflixService{
		collection: collection_,
	}
}

//Most important
//var collection *mongo.Collection

//connect with mongodb only 1 time and very 1st time init

// func (s *netflixService) Insert(movie model.Netflix) {
// 	result, err := s.collection.InsertOne(context.Background(), movie)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Inserted 1 movie in db with id: ", result.InsertedID)
// }

func (s *netflixService) Insert(m interface{}) {
	result, err := s.collection.InsertOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", result.InsertedID)
	//T := reflect.TypeOf(m)
	// switch T {
	// case reflect.TypeOf(model.Netflix{}):
	// 	result, err := s.collection.InsertOne(context.Background(), m)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("Inserted 1 movie in db with id: ", result.InsertedID)
	// }
}

func (s *netflixService) Update(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}
func (s *netflixService) Delete(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := s.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.DeletedCount)
}

func (s *netflixService) DeleteAll() int64 {
	deleteresult, err := s.collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", deleteresult.DeletedCount)
	return deleteresult.DeletedCount
}

func (s *netflixService) GetAllMovie() []primitive.M {
	cursor, err := s.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M //bson.M test
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
func (s *netflixService) GetMovie(movieId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result := s.collection.FindOne(context.Background(), filter)
	var movie primitive.M
	result.Decode(&movie)
	return movie
}

func (s *netflixService) Count(filteringQuery interface{}) (int64, error) {
	result, err := s.collection.CountDocuments(context.Background(), filteringQuery)
	return result, err
}
