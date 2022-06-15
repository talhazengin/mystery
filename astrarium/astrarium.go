package astrarium

import "golang.org/x/exp/slices"

type Astrarium struct {
	Links []*Link
}

func NewAstrarium(links []*Link) *Astrarium {
	return &Astrarium{Links: links}
}

func (a Astrarium) Resolve() []*Link {
	if len(a.Links) == 0 {
		return nil
	}

	for _, link := range a.Links {
		path := []*Link{link}
		if a.XYZ(&path) {
			return path
		}
	}

	return nil
}

func (a Astrarium) XYZ(path *[]*Link) bool {
	if len(*path) == len(a.Links) {
		return true
	}

	lastLink := (*path)[len(*path)-1]

	for _, link := range getNewLinks(path, lastLink.NextLinks) {
		*path = append(*path, link)
		if a.XYZ(path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func getNewLinks(path *[]*Link, nextLinks []*Link) []*Link {
	newLinks := []*Link{}

	for _, link := range nextLinks {
		if !slices.Contains(*path, link) {
			newLinks = append(newLinks, link)
		}
	}

	return newLinks
}
