package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

/*MongoCN objeto de conexión a la BD*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://"+os.Getenv("USERNAME_MONGO")+":"+os.Getenv("PASS_MONGO")+"@"+os.Getenv("NAME_DATABASE")+".jtpcy.mongodb.net/"+os.Getenv("NAME_DATABASE")+"?retryWrites=true&w=majority")

/*conectarBD Conecta base de datos*/
func conectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

/*CheckConnection verifica la conexion a la BD*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
