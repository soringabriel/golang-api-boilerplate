package databases

import (
	"api/helpers"
	"api/logger"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongodbDatabase *MongodbClient

type MongodbClient struct {
	Client         *mongo.Client
	DB             *mongo.Database
	ContextTimeout time.Duration
}

func SetupMongodbDatabase() {
	mongodbUrl := helpers.EnvVariable("MONGODB_URL")
	mongodbDbName := helpers.EnvVariable("MONGODB_DB_NAME")
	mongodbCtxTimeoutString := helpers.EnvVariable("MONGODB_CTX_TIMEOUT")
	mongodbCtxTimeoutInt, err := strconv.Atoi(mongodbCtxTimeoutString)
	if err != nil {
		logger.Instance.Fatal("Failed to convert MONGODB_CTX_TIMEOUT env variable to int", err)
	}
	mongodbCtxTimeout := time.Duration(mongodbCtxTimeoutInt) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), mongodbCtxTimeout)
	defer cancel()
	opts := options.Client().ApplyURI(mongodbUrl)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Instance.Fatal("Failed to connect to mongodb ", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Instance.Fatal("Failed to ping mongodb ", err)
	}

	db := client.Database(mongodbDbName)

	// Setup mongodb client
	MongodbDatabase = &MongodbClient{}
	MongodbDatabase.Client = client
	MongodbDatabase.DB = db
	MongodbDatabase.ContextTimeout = mongodbCtxTimeout
}

func (c *MongodbClient) InsertOne(collection string, newDocument bson.D) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	insertResult, err := c.DB.Collection(collection).InsertOne(ctx, newDocument)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (c *MongodbClient) InsertMany(collection string, newDocuments []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	insertResult, err := c.DB.Collection(collection).InsertMany(ctx, newDocuments)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (c *MongodbClient) Read(collection string, filter bson.M, options *options.FindOptions, result interface{}) (bool, error) {
	// Initial declarations
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	cursor, err := c.DB.Collection(collection).Find(ctx, filter, options)
	if err != nil {
		return false, err
	}
	defer cursor.Close(ctx)

	// Parse results
	if err = cursor.All(ctx, result); err != nil {
		return false, err
	}

	return true, nil
}

func (c *MongodbClient) Count(collection string, filter bson.M) (int64, error) {
	// Initial declarations
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	count, err := c.DB.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *MongodbClient) UpdateOne(collection string, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	updateResult, err := c.DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (c *MongodbClient) UpdateMany(collection string, filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	updateResult, err := c.DB.Collection(collection).UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (c *MongodbClient) DeleteOne(collection string, filter bson.D) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	deleteResult, err := c.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (c *MongodbClient) DeleteMany(collection string, filter bson.D) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ContextTimeout)
	defer cancel()

	// Execute secure query
	deleteResult, err := c.DB.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
