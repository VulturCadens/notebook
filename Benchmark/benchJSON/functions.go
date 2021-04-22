package benchJSON

import (
	"encoding/json"
)

type characters struct {
	Major []string `json:"major"`
	Minor []string `json:"minor"`
}

type volume struct {
	Number int32  `json:"volume"`
	Title  string `json:"title"`
}

type book struct {
	Title        string      `json:"title"`
	Author       string      `json:"author"`
	Country      string      `json:"country"`
	Year         int32       `json:"year"`
	Volumes      []*volume   `json:"volumes"`
	Characters   *characters `json:"characters"`
	Translations []int32     `json:"translations"`
}

var structBook = book{
	Title:   "Les Misérables",
	Author:  "Victor Hugo",
	Country: "France",
	Year:    1862,
	Volumes: []*volume{
		{
			Number: 1,
			Title:  "Fantine",
		},
		{
			Number: 2,
			Title:  "Cosette",
		},
		{
			Number: 3,
			Title:  "Marius",
		},
		{
			Number: 4,
			Title:  "The Idyll in the Rue Plumet and the Epic in the Rue St. Denis",
		},
		{
			Number: 5,
			Title:  "Jean Valjean",
		},
	},
	Characters: &characters{
		Major: []string{"Jean Valjean", "Javert", "Fantine", "Cosette", "Marius Pontmercy", "Éponine"},
		Minor: []string{"Monsieur Thénardier", "Madame Thénardier", "Enjolras", "Gavroche", "Bishop Myriel"},
	},
	Translations: []int32{1862, 1887, 1976, 1987, 2008, 2013},
}

var jsonBook = []byte(`
{
    "title": "Les Misérables",
    "author": "Victor Hugo",
    "country": "France",
    "year": 1862,
    "volumes": [
        {
            "volume": 1,
            "title": "Fantine"
        },
        {
            "volume": 2,
            "title": "Cosette"
        },
        {
            "volume": 3,
            "title": "Marius"
        },
        {
            "volume": 4,
            "title": "The Idyll in the Rue Plumet and the Epic in the Rue St. Denis"
        },
        {
            "volume": 5,
            "title": "Jean Valjean"
        }
    ],
    "characters": {
        "major": ["Jean Valjean", "Javert", "Fantine", "Cosette", "Marius Pontmercy", "Éponine"],
        "minor": ["Monsieur Thénardier", "Madame Thénardier", "Enjolras", "Gavroche", "Bishop Myriel"]
    },
    "translations": [1862, 1887, 1976, 1987, 2008, 2013]
}
`)

func JsonToStruct() *book {
	if err := json.Unmarshal(jsonBook, &structBook); err != nil {
		panic(err)
	}
	return &structBook
}

func StructToJson() []byte {
	jsonBook, err := json.Marshal(&structBook)
	if err != nil {
		panic(err)
	}
	return jsonBook
}
