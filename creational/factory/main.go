package main

import (
	"./data"
)

func main(
	s, _ := data.NewStore(data.MemoryStorage)
	f, _ := s.Open("file")
	
	n, _ := f.Write([]byte("data"))
	defer f.Close()
)