package main

import (
	"mystery/astrarium"
	. "mystery/astrarium"
)

func main() {

	// // mystery 1
	// links := []Link{
	// 	{Node1: &Node{Id: 1}, Node2: &Node{Id: 2}},
	// 	{Node1: &Node{Id: 1}, Node2: &Node{Id: 3}},
	// 	{Node1: &Node{Id: 2}, Node2: &Node{Id: 4}},
	// 	{Node1: &Node{Id: 3}, Node2: &Node{Id: 4}},
	// 	{Node1: &Node{Id: 4}, Node2: &Node{Id: 5}},
	// 	{Node1: &Node{Id: 7}, Node2: &Node{Id: 5}},
	// 	{Node1: &Node{Id: 5}, Node2: &Node{Id: 6}},
	// 	{Node1: &Node{Id: 7}, Node2: &Node{Id: 8}},
	// 	{Node1: &Node{Id: 8}, Node2: &Node{Id: 6}},
	// }

	links := []Link{
		{Node1: &Node{Id: 1}, Node2: &Node{Id: 2}},
		{Node1: &Node{Id: 2}, Node2: &Node{Id: 3}},
		{Node1: &Node{Id: 3}, Node2: &Node{Id: 1}},
	}

	astrarium.NewAstrarium(links)
}
