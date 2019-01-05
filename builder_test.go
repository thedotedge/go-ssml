package ssml

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuilder(t *testing.T) {

	t.Run("should build SSML", func(t *testing.T) {
		fixtures := []struct {
			Builder  Builder
			Expected string
		}{
			{
				Builder:  NewBuilder().Text("asdf").Space().Break(time.Duration(1) * time.Second).Space().Phoneme("test", ALPHABET_IPA, "d͡ʒð"),
				Expected: `<speak>asdf <break time="1000ms" /> <phoneme alphabet="ipa" ph="d͡ʒð">test</phoneme></speak>`,
			},
			{
				Builder:  NewBuilder().Text("asdf2").Space().Date(time.Date(2017, 02, 28, 0, 0, 0, 0, time.Local), DATE_FORMAT_DMY),
				Expected: `<speak>asdf2 <say-as interpret-as="date" format="dmy">20170228</say-as></speak>`,
			},
			{
				Builder:  NewBuilder().Lang("nb-NO", "Spørsmål?"),
				Expected: `<speak><lang xml:lang="nb-NO">Spørsmål?</lang></speak>`,
			},
		}

		for _, f := range fixtures {
			assert.Equal(t, f.Expected, f.Builder.String())
		}

	})

}
