package love

import "sync/atomic"

// Loves array of Loves
type Loves []Love

var loves Loves
var count int32

func init() {
	CreateLove(Love{
		1,
		map[string]string{
			"es": "Vos y %s son el uno para el otro!",
			"pt": "Voce e %s foram feitos um para o outro!",
			"en": "%s and you are a match made in heaven!",
		},
		"http://image1.png",
	})
	CreateLove(Love{
		2,
		map[string]string{
			"es": "Texto 2",
		},
		"http://image2.png",
	})
}

// CreateLove Creates new Love in DB
func CreateLove(l Love) Love {
	l.ID = atomic.AddInt32(&count, 1)
	loves = append(loves, l)
	return l
}

// FindAll Return all Love
func FindAll() Loves {
	return loves
}

// FindByID return Love with ID = id or an empty Love
func FindByID(id int32) Love {
	for _, l := range loves {
		if l.ID == id {
			return l
		}
	}
	return Love{}
}
