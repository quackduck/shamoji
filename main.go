// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var emoji = []rune(`๐ฆ๐ฃ๐ค๐ฅ๐๐ช๐ค๐ฅ๐ฏ๐ฆ๐ช๐ท๐ธ๐คพ๐๐๐๐ชฆ๐ฆขโ๐ฅโง๐ฒ๐ช๐๐๐ก๐๐จ๐๐โ๐จ๐๐งจ๐๐ฆ๐จ๐ฆ๐ง๐๐๐๐ธ๐คฟ๐ถ๐๐๐ฆจ๐ชฒ๐ต๐คถ๐ฅ๐๐ฅท๐ฆ๐๐ณ๐ฅฆ๐๐๐๐๐๐คซ๐ง๐๐ซ๐คฃ๐๐ฅฑ๐ ๐ฃ๐๐๐ฑ๐๐ก๐๐๐ง๐ง๐น๐ฐ๐ค๐ฅผ๐๐ซ๐ฆฉ๐จ๐บ๐๐ญ๐ข๐๐๐ฃ๐๐คฝ๐ฆป๐คฆ๐งฑ๐ต๐ฆ๐ณ๐โพ๐ก๐ป๐ค ๐ฆโฐ๐๐๐ซ๐ฟ๐๐คบโฐ๐ฐ๐ฌ๐๐คดโฐ๐ฐ๐๐๐น๐ท๐ฑ๐ฆ๐ฌ๐ฃ๐โซ๐ฅ๐๐ฐ๐ฆถ๐๐๐ฉ๐น๐๐ต๐ฅฃ๐ดโฑ๐๐ค๐ฆ๐ฆ๐ค๐๐ง๐๐๐๐ฉฐ๐๐น๐ฉ๐บ๐๐ฝ๐ซ๐๐ง๐๐ชฅ๐ค๐๐ฅ๐ป๐คโญ๐บ๐ปโช๐ฆ๐จโ๐ฝ๐๐ก๐ช๐๐พ๐ฆ๐ฒ๐ฌ๐ก๐ฏโก๐ง๐๐ช๐๐๐ฆซ๐๐งญ๐ฌ๐๐งฆ๐ฉ๐๐โฏ๐ถโ๐ฅณโฝโบ๐๐ฅ๐ฆญ๐๐คฒ๐ฉ๐บโช๐๐งถ๐ป๐ชโ๐๐ค๐ฃ๐ช๐คณ๐ซ๐๐ง๐ช๐๐ฎ๐น๐พ๐๐๐จ๐๐ง๐ฒ๐ธโ๐๐โธ๐ฃ๐ค๐๐๐ช๐งณ๐ชฐ๐พ๐ช๐ธ๐ฉโฟ๐๐ฆฎ๐๐๐๐ฑ๐ฌ๐ฆ๐๐๐ฉธ๐ช๐ฅ๐๐๐๐ฐ๐ฉฒ๐ฉ๐ฐ๐ช๐๐๐ง๐งฏ๐๐๐ฆ๐๐คน๐ฅฝ๐ค๐ฏ๐๐ฅ๐๐๐ด๐ชถ๐๐งค๐ด๐จ๐๐ฅฎ๐๐๐๐โฑ๐ฎ๐ช๐๐๐ฐ๐ชด๐ฒ๐๐ถ๐ฐ๐ด๐๐ช๐โณ๐ฆฝ๐๐ซ๐โ๐ฏ๐๐ด๐ก๐ฑ๐ฝ๐๐น๐๐ฅ๐ชโ๐ค๐คงโ๐๐๐ฅฅ๐๐๐ถ๐ฅ๐๐๐๐ฟ๐ฟ๐ฐ๐๐๐ช๐๐๐๐ญ๐งซ๐ต๐๐๐ฝ๐งฒโซ๐ณ๐๐ฅ๐คท๐ชฑโฌ๐ค๐งท๐พ๐๐ก๐งผ๐งง๐ง๐๐ช๐๐๐๐๐ฝ๐ค๐๐ฆ๐๐๐๐ฆ๐ข๐ฅ ๐ ๐ชก๐๐ท๐คฑ๐๐งฅ๐ฐ๐๐ชฃ๐ฅ๐๐ต๐ฅ๐ง๐ฆ ๐๐ฅ๐ฟ๐ฃ๐ณ๐ง๐ช ๐๐ฉ๐ฅ๐ฆ๐๐จ๐๐งข๐โ๐๐ฆ๐ฆฌ๐๐๐ฆน๐ง๐ฆ๐๐ธ๐๐งฟ๐ฟ๐๐ฅซโจ๐๐ง๐๐ผ๐ก๐๐ฉโ๐ชโ๐ฆ๐ป๐๐ ๐ง๐๐ ๐๐ฃ๐ฆ๐ฑ๐ท๐ฆ๐ซ๐ฌ๐ง๐ซ๐๐ผ๐ฆ๐ง๐๐ณ๐๐งฉ๐ฟ๐ซ๐บ๐คญ๐ฟ๐๐๐ค๐ค๐ค๐๐๐๐๐๐๐๐๐๐๐ถ๐ฃ๐ผ๐พ๐ฅง๐งโ๐ผ๐๐ผ๐ชโช๐ข๐๐บ๐๐ถ๐ ๐๐ฅ๐๐ฆ๐ค๐ช๐ฅ๐ง๐ซ๐ฆ๐ฅ๐ฉด๐๐ซ๐ฅฐ๐ฆพ๐ผ๐ฅฏโน๐๐ณ๐๐ฆฏ๐๐๐ธ๐คฅ๐๐๐๐ฒ๐ฅด๐ฌ๐๐งฃ๐ฆ๐จ๐ชจ๐ฑ๐ฉ๐๐น๐ฉนโธ๐๐ชโด๐ฆ๐๐พ๐ฆ๐๐ง๐ฅ๐ฆ๐งฎ๐คค๐ด๐๐๐งช๐คต๐๐ฏ๐ฏ๐ซ๐งฌ๐ฎ๐ฃ๐๐๐จ๐ฆบ๐ค๐๐๐ค๐๐บ๐ฆ๐๐ฌ๐ฆ๐พ๐ช๐ฅ๐ซ๐ฆต๐ท๐ฅญ๐ฝ๐๐๐ค๐๐ช๐๐บ๐๐ผ๐ฆค๐๐ฆ๐ธ๐โ๐ด๐โ๐ฅป๐๐ฅก๐ข๐ฝ๐๐ฅ๐บ๐โฌ๐ข๐๐ป๐ฑโ๐ฅบ๐๐ฏ๐ง๐โฝ๐๐๐ช๐ฎ๐ช๐ช๐ฟ๐ต๐ฅ๐ฐ๐บ๐งบโฟ๐๐ก๐งพ๐ซ๐ฑ๐๐๐๐ฅ๐คฉ๐จ๐ฆ๐ง๐ฅถ๐๐ฎ๐ง๐๐ ๐ง๐ฏ๐ค๐ญ๐๐๐ฆก๐ฒ๐ฒ๐๐๐ต๐ป๐ชต๐ฒ๐ข๐ง๐๐ฉ๐ป๐๐ง๐จ๐๐๐ฏโ๐ฅ๐๐ฆฆ๐ฑ๐๐ผโ๐ช๐ช๐ฐ๐๐๐๐ค๐ผโ๐ฅฌ๐๐๐๐๐บ๐ฅ๐๐ฅค๐ธ๐คผ๐โญ๐คฐ๐ฏ๐ธ๐โฉ๐ง๐๐ฆ๐คฎ๐๐๐ฉณ๐๐ฅฟ๐๐ด๐งน๐ฅข๐ก๐๐ณ๐ญโ๐๐๐๐๐ฆช๐๐ฅจ๐ฎ๐๐ฅ๐๐๐ฟ๐ญ๐๐๐๐๐ฑ๐ข๐๐ถ๐๐ฆฅ๐๐๐ฅ๐๐ฑ๐ ๐ฅ๐จ๐๐๐ณ๐ป๐งธ๐ฆ๐ฅพ๐๐๐ข๐๐๐ท๐ฆ๐ฃ๐ช๐ถ๐๐ญ๐๐ฅ๐ฅ๐๐ญ๐๐ณ๐๐ค๐๐๐๐ด๐๐ฆฃ๐ท๐๐๐๐๐ฃ๐๐ ๐๐ญ๐ค๐๐๐๐ฟ๐คช๐๐ซ๐๐๐ฅ๐ฅ๐๐๐คฌ๐ช๐ฝ๐บโ๐พ๐ฆ๐ข๐คก๐ท๐๐๐๐๐๐๐พ๐๐ฉ๐๐๐ฆด๐งต๐๐ซ๐ฆ๐๐ฆธ๐๐ฅ๐๐โ๐ฏ๐ง ๐ค๐๐๐งก๐ฌ๐๐๐ฆ๐๐ง๐๐ค๐๐ด๐โฒ๐ฅ๐ก๐ฃ๐๐ป๐ฎ๐น๐๐๐ชค๐ฆ๐ด๐ง๐๐ฎ๐๐ฅธ๐ก๐ต๐โท๐๐๐๐นโฎ๐น๐๐๐๐ก๐ซ๐๐๐ณ๐ฎโฌ๐ฆ๐๐พโ๐ชง๐๐๐ฅฉ๐๐ฐ๐ฅ๐ข๐๐๐๐๐งด๐๐ข๐งโฝ๐๐งฝ๐๐ถ๐๐๐โน๐ ๐ง๐ซ๐๐โณ๐ฉ๐ โญ๐ฅฒ๐ต๐น๐ต๐ท๐ซ๐๐ชข๐ฅ๐ฉบ๐ง๐๐ญ๐๐๐๐ท๐ข๐๐๐๐๐ฆ๐ง๐๐ฝโต๐ฒ๐๐๐๐๐๐๐ ๐ฆง๐ซ๐ฅ๐๐ง๐ผ๐ผ๐ณ๐๐ถ๐๐๐๐๐ก๐๐๐งป๐ฒ๐๐ค๐ฌ๐ท๐ง๐ค๐ฏ๐โฒ๐ฅ๐๐งฐ๐ผ๐๐๐ต๐ฎ๐๐ด๐ท๐จโ๐ค๐ฌ๐๐๐ซ๐๐ฆผ๐ง๐โ๐ฉฑ๐ง๐ฅ๐ชโพ๐๐ฅ๐ฑ๐ฆ๐๐ฅ๐บ๐ฅ๐ฆ๐คจ๐๐ญ๐๐๐โฉ๐ฝ๐ ๐ฅ๐ฐ๐๐๐ท๐ป๐คข๐๐๐ธ๐ช๐๐ฆท๐ซ๐๐๐๐ต๐๐๐ฆฟ๐ฒ๐๐ณ๐๐ ๐ฎ๐๐ค๐ฟโบ๐ฉ๐ฌ๐ณ๐๐๐๐๐ญ๐ซ๐บ๐๐ง๐๐ช๐๐น๐๐๐๐ช๐๐๐ฅ๐ผ๐ช๐๐๐๐ฅต๐ค๐๐ค๐ป๐๐ถ๐ฅ๐ง๐คฏ๐๐ค๐ต๐ธ๐น๐ธ๐๐ฅช๐๐๐๐๐๐ค๐ธ๐ง๐ถ๐ง๐ฆ๐ง๐๐๐ด๐คธ๐ณ๐๐ง๐ชณ`)

var (
	shamoji256 = encoding{emoji, 32}
	shamoji1   = encoding{emoji, 20}
)

func main() {
	if len(os.Args) > 1 { // there's an arg
		switch os.Args[1] {
		case "-r", "--replace":
			replace()
			return
		}
	}

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	hash := sha256.Sum256(in)

	fmt.Println(shamoji256.Encode(hash[:]))
}

func replace() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)

	currWord := ""
	//spaceToPrint := ""
	for in.Scan() {
		d := in.Bytes()
		r, _ := utf8.DecodeRune(d)
		if unicode.IsSpace(r) {
			fmt.Print(currWord + string(d))
			currWord = ""
			//spaceToPrint += string(d)
		} else {
			currWord += string(r)
		}
		// check if the word is a hex encoded hash
		if len(currWord) == 64 || len(currWord) == 40 {
			d, err := hex.DecodeString(currWord)
			if err == nil {
				if len(currWord) == 64 {
					fmt.Print(shamoji256.Encode(d))
				} else {
					fmt.Print(shamoji1.Encode(d))
				}
				//fmt.Print(shamoji.Encode(d))
				currWord = ""
			}
		}
	}
}

type encoding struct {
	set []rune
	// dataLen is the length of the byte data used. In duckcoin, this is always 24.
	dataLen int
}

func (e *encoding) Encode(data []byte) string {
	convertedBase := toBase(new(big.Int).SetBytes(data), "", e.set)
	// repeat emoji[0] is to normalize result length. 0 because that char has zero value in the set
	return strings.Repeat(string(e.set[0]), e.EncodedLen()-len([]rune(convertedBase))) + convertedBase
}

func (e *encoding) Decode(data string) ([]byte, error) {
	if len([]rune(data)) != e.EncodedLen() {
		return nil, errors.New("could not decode: invalid length of data: " + data)
	}
	num, err := fromBase(data, e.set)
	if err != nil {
		return nil, err
	}
	return num.FillBytes(make([]byte, e.dataLen)), nil
}

func (e *encoding) EncodedLen() int {
	return int(math.Ceil(
		float64(e.dataLen) * math.Log2(256) / math.Log2(float64(len(e.set))),
	))
}

func toBase(num *big.Int, buf string, set []rune) string {
	base := int64(len(set))
	div, rem := new(big.Int), new(big.Int)
	div.QuoRem(num, big.NewInt(base), rem)
	if div.Cmp(big.NewInt(0)) != 0 {
		buf += toBase(div, buf, set)
	}
	return buf + string(set[rem.Uint64()])
}

func fromBase(enc string, set []rune) (*big.Int, error) {
	result := new(big.Int)
	setlen := len(set)
	encRune := []rune(enc)
	numOfDigits := len(encRune)
	for i := 0; i < numOfDigits; i++ {
		mult := new(big.Int).Exp( // setlen ^ numOfDigits-i-1 = the "place value"
			big.NewInt(int64(setlen)),
			big.NewInt(int64(numOfDigits-i-1)),
			nil,
		)
		idx := findRuneIndex(encRune[i], set)
		if idx == -1 {
			return nil, errors.New("could not decode " + enc + ": rune " + string(encRune[i]) + " is not in charset")
		}
		mult.Mul(mult, big.NewInt(idx)) // multiply "place value" with the digit at spot i
		result.Add(result, mult)
	}
	return result, nil
}

// findRuneIndex returns the index of rune r in slice a, or -1 if not found
func findRuneIndex(r rune, a []rune) int64 {
	for i := range a {
		if a[i] == r {
			return int64(i)
		}
	}
	return -1
}
