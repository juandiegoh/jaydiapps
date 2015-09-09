package love

import "sync/atomic"

// LovesIndex map of Loves
type LovesIndex map[int32]Love

// Loves array of Loves
type Loves []Love

var lovesMap LovesIndex
var loves Loves

var count int32

func init() {
	lovesMap = make(map[int32]Love)

	CreateLove(Love{
		1,
		map[string]string{
			"es": "Vos y %s son el uno para el otro!",
			"pt": "Voce e %s foram feitos um para o outro!",
			"en": "%s and you are a match made in heaven!",
		},
		"https://s3-sa-east-1.amazonaws.com/jaydiapps/love/alto.png",
	})
	CreateLove(Love{
		2,
		map[string]string{
			"es": "Si no lo estropeas tendrás suerte con %s",
			"pt": "Se você não estragar tudo terá sorte com %s",
			"en": "If you don't mess it up, it will last with %s",
		},
		"https://s3-sa-east-1.amazonaws.com/jaydiapps/love/medioalto.png",
	})
	CreateLove(Love{
		3,
		map[string]string{
			"es": "Mejor busca a otra persona, olvidate de %s",
			"pt": "Melhor procurar outra pessoa, esquece %s",
			"en": "Better off with somebody else, forget about %s",
		},
		"https://s3-sa-east-1.amazonaws.com/jaydiapps/love/medio.png",
	})
	CreateLove(Love{
		4,
		map[string]string{
			"es": "Olvidate de %s, nunca estará contigo",
			"pt": "Esquece %s, nunca estará contigo",
			"en": "Forget it, you don't have a chance with %s",
		},
		"https://s3-sa-east-1.amazonaws.com/jaydiapps/love/bajo.png",
	})

}

// CreateLove Creates new Love in DB
func CreateLove(l Love) Love {
	l.ID = atomic.AddInt32(&count, 1)
	lovesMap[l.ID] = l
	loves = append(loves, l)
	return l
}

// FindAll Return all Love
func FindAll() Loves {
	return loves
}

// FindByID return Love with ID = id or an empty Love
func FindByID(id int32) Love {
	l, present := lovesMap[id]
	if present {
		return l
	}

	return Love{}
}
