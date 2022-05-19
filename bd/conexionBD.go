package bd

import (
	"context"
	"log"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

//mongo.
//MongoCN es la coneccion y esta en mayuscula para que se pueda acceder desde todo el proyecto
var MongoCN = conectatBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://Edison128:rNVLaM4lZvNUCe3f@cluster0.r0xht.mongodb.net/test")

/*conectatBD es la base de datos*/
func conectatBD() *mongo.Client {
	//el contexto es un entoron que esta disponible durante la coneccion de GO con la base de datos

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

	log.Println("Conexion exitosa con la Bd")
	return client
}

func CheckConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
