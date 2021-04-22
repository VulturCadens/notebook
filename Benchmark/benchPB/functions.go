package benchPB

import (
	proto "github.com/golang/protobuf/proto"
)

var structBook = Book{
	Title:   "Les Misérables",
	Author:  "Victor Hugo",
	Country: "France",
	Year:    1862,
	Volumes: []*Volume{
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
	Characters: &Characters{
		Major: []string{"Jean Valjean", "Javert", "Fantine", "Cosette", "Marius Pontmercy", "Éponine"},
		Minor: []string{"Monsieur Thénardier", "Madame Thénardier", "Enjolras", "Gavroche", "Bishop Myriel"},
	},
	Translations: []int32{1862, 1887, 1976, 1987, 2008, 2013},
}

var protobufBook, _ = proto.Marshal(&structBook)

func ProtobufToStruct() *Book {
	if err := proto.Unmarshal(protobufBook, &structBook); err != nil {
		panic(err)
	}
	return &structBook
}

func StructToProtobuf() []byte {
	protobufBook, err := proto.Marshal(&structBook)
	if err != nil {
		panic(err)
	}
	return protobufBook
}
