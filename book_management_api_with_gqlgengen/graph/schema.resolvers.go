package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/zaahidali/book_management_api_with_gqlgengen/graph/model"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string, publicationYear int, genre string) (*model.Book, error) {
	// randNum, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999))
	randNum, err := rand.Int(rand.Reader, big.NewInt(9999))

	if err != nil {
		log.Printf("Failed to generate random ID: %v", err)
		return nil, err
	}

	book := &model.Book{
		ID:              randNum.String(),
		Title:           title,
		Author:          author,
		PublicationYear: publicationYear,
		Genre:           genre,
	}

	_, err = r.DB.NewInsert().Model(book).Exec(ctx)
	if err != nil {
		log.Printf("Failed to add book: %v", err)
		return nil, err
	}

	return book, nil

}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, title *string, author *string, publicationYear *int, genre *string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: UpdateBook - updateBook"))
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - deleteBook"))
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: Book - book"))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Books - books"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
