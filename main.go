package main

import (
	"fmt"

	"github.com/naveeharn/golang101_project_cinema/movie"
	"github.com/naveeharn/golang101_project_cinema/ticket"
)

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 7.6)
	ticket.Buy(movieName)
	fmt.Println()

}
