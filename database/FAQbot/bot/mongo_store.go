package bot

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type botMongoStore struct {
	collection *mongo.Collection
}

func NewBotMongoStore(config MongoConfig) (BotStore, error) {
	clientOptions := options.Client().ApplyURI(config.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	collection := db.Collection(config.Collection)
	return &botMongoStore{collection: collection}, nil
}

func (b *botMongoStore) Create(bot *Bot) (*Bot, error) {
	bot.Id = uuid.New().String()
	_, err := b.collection.InsertOne(context.TODO(), bot)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (b *botMongoStore) GetById(id string) (*Bot, error) {
	filter := bson.D{{"id", id}}
	bot := &Bot{}
	err := b.collection.FindOne(context.TODO(), filter).Decode(&bot)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (b *botMongoStore) List() ([]Bot, error) {
	filter := bson.D{}
	bots := []Bot{}
	cursor, err := b.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &bots)
	if err != nil {
		return nil, err
	}
	return bots, nil
}
