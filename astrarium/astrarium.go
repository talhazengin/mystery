package astrarium

type Astrarium struct {
	Links []Link
}

func NewAstrarium(links []Link) *Astrarium {
	return &Astrarium{Links: links}
}

func (a *Astrarium) AddNewLink(link Link) {
	a.Links = append(a.Links, link)
}

func (a Astrarium) Resolve() *Path {
	if len(a.Links) == 0 {
		return nil
	}

	for i, link := range a.Links {
		link.Node1.Id
	}
}

func (a Astrarium) getLinks(nodeId int) []Link {
	var links []Link

	for _, link := range a.Links {
		if link.Node1.Id == nodeId || link.Node2.Id == nodeId {
			links = append(links, link)
		}
	}

	return links
}
