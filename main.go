package main

type Store struct {
	A int
	B int
}

type Calculate interface {
	Add()
	Sub()
}

func (s Store) Add() {

}

func (s Store) Sub() {

}

func main() {
	var c Calculate
	c = &Store{}
	c.Add()
	c.Sub()
}
