package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}
