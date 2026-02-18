package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(uri string, dbName string) *mongo.Database {
   
    if uri == "" || dbName == "" {
        log.Fatal("ERRORE: Parametri di connessione vuoti!")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Connessione fallita: %v", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatalf("Ping fallito: %v", err)
    }

    log.Printf("Connesso correttamente a: %s", dbName)
    return client.Database(dbName)
}