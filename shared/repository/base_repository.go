package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepository[T any] struct {
	Collection *mongo.Collection
}

func RegisterBaseRepository[T any](db *mongo.Database, collectionName string) *BaseRepository[T] {
	return &BaseRepository[T]{
		Collection: db.Collection(collectionName),
	}
}

func (r *BaseRepository[T]) Create(ctx context.Context, doc *T) error {
	if setter, ok := any(doc).(interface{ SetDocTypeFrom(v interface{}) }); ok {
		setter.SetDocTypeFrom(doc)
	}
	_, err := r.Collection.InsertOne(ctx, doc)
	return err
}

func (r *BaseRepository[T]) Fetch(ctx context.Context, filter bson.M, findOptions options.FindOptions) ([]T, error) {
	cursor, err := r.Collection.Find(ctx, filter, &findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	err = cursor.All(ctx, &results)
	return results, err
}

func (r *BaseRepository[T]) GetOne(ctx context.Context, filter bson.M) (T, error) {
	var result T
	err := r.Collection.FindOne(ctx, filter).Decode(&result)
	return result, err
}
