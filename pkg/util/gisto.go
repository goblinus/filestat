package util

import (
	"fmt"
	"io"
	"os"
)

const (
	GistSpace      = '\u0010'
	GistReturn     = '\u000D'
	GistNewLine    = '\u000A'
	GistElement    = '\u2591'
	GistSpaceDispl = '\u035f'
	Single         = 1
)

var (
	SpaceSymbols = map[rune]rune{
		GistSpace:   GistSpaceDispl,
		GistReturn:  GistSpaceDispl,
		GistNewLine: GistSpaceDispl,
	}
	GistFullLength = []rune{
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
		GistElement,
	}
)

func DrawHistogramm(data map[rune]int) {
	sumVal := sum(data)
	for k, v := range data {

		//Избегаем перевода строк и прочего сбоя форматирования гистограммы
		symbolToDisplay := k
		if _, ok := SpaceSymbols[k]; ok {
			symbolToDisplay = SpaceSymbols[k]
		}

		symbolInterest := (float64(sumVal) / 100) * float64(v)
		symbolGistLine := getGistLine(symbolInterest)
		symbolMetaData := fmt.Sprintf("[%c <%U>, стат: %d/%d/%.f%%]", symbolToDisplay, k, sumVal, v, symbolInterest)
		io.WriteString(os.Stdout, fmt.Sprintf("%s %s\n", symbolMetaData, symbolGistLine))
	}
}

func sum(data map[rune]int) int {
	var sumVal int
	for _, v := range data {
		sumVal += v
	}
	return sumVal
}

//gistLine масштабирует процент для вывода длины линии в масштабе 10/1
func getGistLine(interest float64) string {
	idx := int(interest / 10)
	if idx < Single {
		idx = Single
	}
	return string(GistFullLength[0:idx])
}
