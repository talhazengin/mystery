package main

import (
	"mystery/astrarium"
	. "mystery/astrarium"
)

func main() {

	link1 := Link{Id: 1}
	link2 := Link{Id: 2}
	link3 := Link{Id: 3}
	link4 := Link{Id: 4}
	link5 := Link{Id: 5}
	link6 := Link{Id: 6}

	link1.NextLinks = []*Link{&link4, &link6, &link3, &link5}
	link2.NextLinks = []*Link{&link4, &link6}
	link3.NextLinks = []*Link{&link5, &link1, &link6, &link4}
	link4.NextLinks = []*Link{&link6, &link3, &link1, &link2}
	link5.NextLinks = []*Link{&link3, &link1}
	link6.NextLinks = []*Link{&link2, &link4, &link3, &link1}

	links := []*Link{&link1, &link2, &link3, &link4, &link5, &link6}

	path := astrarium.NewAstrarium(links).Resolve()

	for _, l := range path {
		print(l.Id, " ")
	}
}
