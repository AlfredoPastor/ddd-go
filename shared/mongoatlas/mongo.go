package mongoatlas

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config interface {
	GetMongoAtlasEnv() *MongoAtlasEnv
}

//MongodbEnv obtiene la configuracion desde el entorno
type MongoAtlasEnv struct {
	MongoAtlasURL         string        `env:"MONGOATLAS_URL"`
	MongoAtlasDb          string        `env:"MONGOATLAS_DB"`
	MongoAtlasConnTimeout time.Duration `env:"MONGOATLAS_CONN_TIMEOUT" envDefault:"10s"`
	CredentialsFile       string        `env:"MONGOATLAS_CREDENTIALS_FILE"`
	User                  string        `env:"MONGOATLAS_USER" envDefault:"root"`
	Password              string        `env:"MONGOATLAS_PASSWORD" envDefault:""`
}

type MongoClient struct {
	*mongo.Client
	database string
}

func NewMongoClient(cfg Config) (*MongoClient, error) {
	cli, err := connect(cfg.GetMongoAtlasEnv())
	if err != nil {
		return &MongoClient{}, err
	}

	return &MongoClient{
		Client:   cli,
		database: cfg.GetMongoAtlasEnv().MongoAtlasDb,
	}, nil
}

func (m *MongoClient) CreateCollection(name string) *mongo.Collection {
	return m.Client.Database(m.database).Collection(name)
}

func connect(env *MongoAtlasEnv) (client *mongo.Client, err error) {
	if env.MongoAtlasConnTimeout.Seconds() < 2.0 && env.MongoAtlasConnTimeout.Seconds() > 30.0 {
		return nil, fmt.Errorf("mongoAtlasconfig: timeout invalido: %s", env.MongoAtlasConnTimeout)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", env.User, env.Password, env.MongoAtlasURL, env.MongoAtlasDb)))
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear el cliente de Mongo Atlas: %s", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("el ping hacia Mongo Atlas falló: %s", err)
	}
	return client, err
}
