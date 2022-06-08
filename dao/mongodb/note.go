package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"wechat/model"
)

//func InsertNote(noteId int) error {
//	collection := mongoDB.Database("wechat").Collection("note")
//	doc := model.NoteContent{
//		NoteId: noteId,
//		Delta:  "",
//	}
//	insert, err := collection.InsertOne(context.TO。DO), doc)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	fmt.Println("Inserted a Single Document: ", insert.InsertedID)
//	return nil
//}

func UpdateNote(noteId int, delta string) error {
	collection := mongoDB.Database("wechat").Collection("note")
	filter := bson.D{{"noteid", noteId}}

	update := bson.D{
		{"$set", bson.D{
			{"delta", delta},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("update note error:", err)
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func DeleteNote(noteId int) error {
	collection := mongoDB.Database("wechat").Collection("note")
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"noteid", noteId}})
	if err != nil {
		log.Println("delete note error:", err)
		return err
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	return nil
}

func SelectNote(noteId int) (string, error) {
	collection := mongoDB.Database("wechat").Collection("note")
	filter := bson.D{{"noteid", noteId}}
	var result model.NoteContent
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println("select note error:", err)
		return "", err
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result.Delta, nil
}
