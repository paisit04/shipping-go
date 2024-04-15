package translation_test

import (
	"testing"

	"github.com/paisit04/shipping-go/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "french",
			Translation: "bonjour",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
	}

	underTest := translation.NewStaticService()
	for _, test := range tt {
		// Act
		res := underTest.Translate(test.Word, test.Language)

		// Assert
		if res != test.Translation {
			t.Errorf(`Expected "%s" to be "%s" from "%s" but got "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}
}
