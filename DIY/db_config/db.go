package db_config

import (
	"context"
	"fmt"
	"log"

	"github.com/dunzoit/diy/lalit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://lalit:dunzo527@cluster0.apnwt.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "catalouge"
const colName = "products"

var collection *mongo.Collection

func conn() (collection *mongo.Collection) {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
	return
}

func AddProduct(product models.Product) {

	inserted, err := collection.InsertOne(context.Background(), product)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product added in db with id: ", inserted.InsertedID)
}

// delete one product
func DeleteOneProduct(productId string) {
	id, err := primitive.ObjectIDFromHex(productId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}

//delete all products

func DeleteAllProducts() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func GetAllProducts() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var products []primitive.M

	for cur.Next(context.Background()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	defer cur.Close(context.Background())
	return products
}
