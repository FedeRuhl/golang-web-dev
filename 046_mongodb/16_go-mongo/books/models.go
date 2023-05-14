package books

import (
	"context"
	"errors"
	"golang-web-dev/046_mongodb/16_go-mongo/config"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type Book struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float64 // `json:"price" bson:"price,truncate"`
}

func AllBooks() ([]Book, error) {
	q, err := config.Books.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	bks := []Book{}
	err = q.All(context.Background(), &bks)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	isbn := r.FormValue("isbn")
	bk := Book{}
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}
	q := config.Books.FindOne(context.Background(), bson.M{"isbn": isbn})
	err := q.Decode(&bk)

	if err != nil {
		return bk, err
	}

	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = f64

	// insert values
	_, err = config.Books.InsertOne(context.Background(), bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		log.Fatal(err)
		return bk, errors.New("406. Not Acceptable. Enter number for price")
	}
	bk.Price = f64

	// update values
	_, err = config.Books.UpdateOne(context.Background(), bson.M{"isbn": bk.Isbn}, bson.M{"$set": bk})
	if err != nil {
		log.Fatal(err)
		return bk, err
	}
	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	_, err := config.Books.DeleteOne(context.Background(), bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
