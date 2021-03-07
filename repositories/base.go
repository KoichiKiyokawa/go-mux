package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/imdario/mergo"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// BaseRepository is
type BaseRepository struct {
	name string
}

var ctx = context.Background()
var sa = option.WithCredentialsFile("./credentials.json")
var app, _ = firebase.NewApp(ctx, nil, sa)
var client, _ = app.Firestore(ctx)

// withID is
func withID(id string, data interface{}) map[string]interface{} {
	res := map[string]interface{}{"ID": id}
	mergo.Merge(&res, data)
	return res
}

func (r BaseRepository) getCollectionRef() *firestore.CollectionRef {
	return client.Collection(r.name)
}

// Find is
func (r BaseRepository) Find(id string) (interface{}, error) {
	snap, err := r.getCollectionRef().Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	return withID(id, snap.Data()), nil
}

// All is
func (r BaseRepository) All() (interface{}, error) {
	var data []map[string]interface{}
	docIter := r.getCollectionRef().Documents(ctx)
	for {
		doc, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, withID(doc.Ref.ID, doc.Data()))
	}
	return data, nil
}

// Create is
func (r BaseRepository) Create(data interface{}) (interface{}, error) {
	docRef, _, err := r.getCollectionRef().Add(ctx, data)
	if err != nil {
		return nil, err
	}
	return withID(docRef.ID, data), nil
}
