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

var emoji = []rune(`🦆🐣🐤🐥🎟🪐🤖🌥🎯🦋🟪🛷📸🤾🔞🎀🔈🪦🦢➖🥊⚧🌲🪗😚🎌🌡🐁👨🎗🌝⛄💨🖕🧨📏🦒🎨🔦🧆🚊🗑🏅🛸🤿🐶🚂🚀🦨🪲🍵🤶🎥💟🥷🦟🐞😳🥦📎🏕🐅🏇🕚🤫🕧🌀😫🤣🕘🥱🐠📣🌙🏒💱🎏🟡🚅🏝🔧🧃😹🅰🕤🥼🚍🫂🦩🖨😺🌗🌭📢🖐💎🚣🌉🤽🦻🤦🧱🌵💦📳🔚⚾😡😻🤠🐦⛰🕕🐂💫📿🌚🤺⏰😰📬🛄🤴➰🍰🙋🍇🔹🏷👱🦀🚬🏣💇⚫🥑💃🚰🦶🐟😒📩🐹💋🐵🥣📴⛱💅🤑🦎🦜🤐👌🚧🏈💗🎋🩰🛌🏹😩🏺🏏🗽🫀🎅🧘🎁🪥😤📘🥇🛻🤌⏭🕺💻⛪🚦🟨❌💽🙆💡🪒🔛🍾🦉🖲💬🚡👯⚡🧁🏉🌪💁🌎🦫💔🧭🏬🍒🧦🍩🐖😞⏯🛶⛓🥳◽⏺🚜🥈🦭👆🤲🏩🛺⚪🌇🧶👻🪙➗🚆🍤👣📪🤳🔫🏟🧗🪄🐎🚮🚹🎾🕜👒🔨🌊🧎💲🌸⌚🍘🖋⏸🍣🌤📖💐🪓🧳🪰😾👪👸💩♿🚝🦮🍈🔘🚋🔱🛬🦄🏑👘🩸💪👥👋🛏🎚🏰🩲🔩🔰🎪🐔🌒🧓🧯🔗👟🦍🎍🤹🥽🟤🏯📇🟥🛀🕗🔴🪶🔄🧤🌴🚨🌑🥮👗🐍🔖🕞⏱📮🪟🐘🍌🕰🪴📲🐌🔶📰🍴🏜😪🍔⛳🦽🏂👫🔉✊📯👈🐴📡🎱🐽🏌🕹😄🥌🪆✋🤞🤧⛏🎐🚓🥥🐃🚉😶🥖🛑🚛🔇💿🍿🛰🐊📋🪁🛂🚚🚙🏭🧫👵🆖🏍🌽🧲⏫🔳😔🥒🤷🪱⬜👤🧷🅾🏊🎡🧼🧧🧄🍉🍪👚🔎🐛🍅🎽🤍🆚🌦🚔🛅😓🦘👢🥠🎠🪡👇🍷🤱📕🧥🎰👉🪣🥟🐐🔵🥐💧🦠👛🥗🎿🛣🕳🧔🪠🍆🚩💥🎦🚐🏨💌🧢🙂✅📚🏦🦬🕑👖🦹🧑🦖😌🐸🍐🧿😿🎇🥫✨👝🧖🍗🎼🕡🚈🟩⚓🪚❔🦇🔻🚃💠🟧🐒🚠😛🗣🦊🐱💷🦂🟫🔬👧🫖📟🍼🟦🧋🏁🚳🆗🧩🚿🫑🚺🤭🐿🚎📁💤🏤🤎📄👜🐜🌅🔁🔃😀💘🔐🚘🕶🕣🚼💾🥧📧❕🔼🗄🌼🪞⏪🕢📔👺🕒👶🏠💕🥓🎉🦙🛤🪘📥🧏🫁🦑🥁🩴🔕🌫🥰🦾📼🥯⛹🐕🛳🕟🦯🙀📜🍸🤥💆🕓🔓🍲🥴🐬🛃🧣👦🐨🪨😱🛩👓👹🩹⛸🛋🪜⛴🦁🏛🚾🦚👎🧙🥉🦛🧮🤤🎴🕊🍜🧪🤵🏓💯🍯🚫🧬💮🟣🛁😉🗨🦺🤏🎆😊🤗🚑🌺🦈🕔😬🦏🗾🚪😥🫒🦵😷🥭🚽🔌😝🎤🐚🪂🆘🍺🍀🛼🦤💖🍦🕸🆙⛑💴🐙❎🥻🍕🥡🚢👽😑🥜🔺🐉⬛🐢🛎🍻🅱⌛🥺🏚🐯🧞🔋⚽📀🍂🪝🍮🪖🐪👿🏵🥙💰🗺🧺➿🌓🗡🧾🐫🌱📉👙🌕🚥🤩🍨🦓🌧🥶📆🏮🧜🐗🌠🧟🚯🤙🐭🚏🚕🦡🔲🚲📂🌟📵📻🪵👲🎢🧚🔟👩🚻😎🧐😨📍🎞😯❗🥎🐄🦦🍱🎃🖼⛎🪔🪃🌰🗞🆕👍🤒💼❓🥬💈🙊🛗🎄📺🥚🍓🥤💸🤼🌜⭕🤰🗯🏸🎊⛩🧒🗜🦅🤮🔂📝🩳😈🥿🍎🚴🧹🥢🍡🎖👳😭☔🆓🎙📑🍑🦪🏗🥨🐮🐝🥘🌋🔆🌿👭🚗🔜🍟👐🚱🍢🏐🌶🖊🦥🃏🚒🥀🖇🖱🕠🥅📨🏞📅🗳🐻🧸🦗🥾🙉🌐🛢😐😁🎷🦌😣🏪🚶🌘🚭💄🥍🥕😏🍭📒💳🗓🤓🚞💂🙁🏴🍛🦣🔷🙅👔🆎💝🎣🌏👠😃🎭🤝📗🏙📈🗿🤪🍍🍫🏆👊🕥🔥😟🐆🤬🪅📽🐺⛔🌾🦔💢🤡🌷🌆🙏📌💉😆🗝👾🖍🐩👃🀄🦴🧵🗂🛫🕦🚁🦸🍝🥃🆒🕝⛈🌯🧠🤚🛕🆔🧡👬😅🔙🦞🎈🧕🏖📤🕙😴💓⛲🛥🐡💣🕛🌻🎮🎹😂👏🪤🦐🛴🧀🚌🔮💍🥸🏡🚵🐑⛷🛍🌛🍊🍹⏮🛹🔊📙🏎👡🫓🍋🐏🎳👮⏬🦃🐋🐾⛅🪧😖💊🥩😘👰🏥🏢🆑📊🍚👑🧴🌁🟢🧊⛽👂🧽🏃💶🍄🙎👞⏹🛠🧝🫐🏔👕⏳🎩🍠⭐🥲😵💹🕵🚷📫📓🪢🥞🩺🧛🙌🔭🚄🙍📛🕷😢🔝🎓💚💞😦🧍🔍😽⛵🐲🗒🎎🐓💀👀📐📠🦧🎫🍥🔒😧🗼👼🌳🏘📶🗃🏄💛😇🛡🙇🌄🧻😲😗🚤🌬👷🏧🤟🕯🎂⏲🥂💙🧰🐼🌌📃💵😮🔏👴🐷🌨➕🤘🍬🛖👅🫔🌂🦼🧂🍁☕🩱🧈🥄🔪◾🚇🥔📱🦝🕖🥋💺🥏📦🤨🕐💭💑🍞🌖⏩🍽🟠🥛🐰🔀🍃📷🗻🤢😕🐇🔸🪕🏋🦷🏫👁🌔🔔🎵👄🍏🦿🎲🔅🐳🚟😠🌮🖖🤛🅿⛺🌩🎬🏳😍🐀😜🙈📭🫕🎺📞🧅🖌🪀🎑🌹🌞💜🌈🪑🙃😋🖥😼🪛🍙💏🚖🥵🤔🙄🤜🎻🌃🍶🥝🎧🤯🔑🖤🛵😸📹🚸🛒🥪🎛🌍🐈🍖🏀🤕🎸🍧🎶🐧🦕🧇💒🎒🕴🤸🍳😙🧉🪳`)

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
