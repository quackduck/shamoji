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

var shamoji256 = encoding{emoji, 32}
var shamoji1 = encoding{emoji, 20}

// var emojiJS = []rune{128169, 128123, 129302, 127875, 129313, 129413, 128034, 128020, 128039, 128065, 128077, 128078, 127797, 127758, 127752, 127823, 9917, 9730, 11088, 128640, 127829, 128663, 127993, 127911, 128273, 127929, 10084, 127794, 128121, 128069, 128293, 128132, 128125, 129419, 128029, 128044, 127819, 127820, 129472, 127846, 127856, 129409, 127936, 127941, 127922, 128690, 128641, 127774, 127822, 127821, 129373, 129361, 127813, 9200, 128176, 128142, 128222, 128163, 9986, 128274, 128269, 9999, 9992, 128652, 128659, 128757, 9973, 9875, 128678, 9924, 127805, 9940, 129412, 10067, 9888, 127359, 128350, 128083, 127801, 128375, 128013, 128012, 128010, 129416, 128011, 127789, 127925, 9851, 128276, 127937, 128212, 128218, 128206, 128204, 127818, 127827, 127828, 127839, 127871, 127849, 129371, 127850, 9749, 127852, 129365, 127826, 129374, 129370, 129360, 127838, 127824, 127815, 127817, 128043, 128024, 129423, 129421, 128014, 128017, 128056, 128175, 9728, 9760, 128060, 9889, 128172, 9884, 128526, 128561, 127755, 127958, 127771, 127918, 10052, 128045, 128053, 128722, 129348, 9729, 128208, 127804, 127800, 128167, 127786, 127812, 128170, 128027, 127930, 128155, 128154, 128153, 128156, 128075, 128296, 128161, 128679, 127921, 128298, 127913, 127955, 9971, 9975, 127938, 128225, 127853, 128165, 128037, 9752, 128031, 128241, 128096, 128374, 128188, 127868, 128030, 128028, 128376, 129422, 129408, 127942, 9878, 128301, 127891, 128300, 128247, 127873, 128081, 128064, 127916, 128048, 127912, 127908, 127928, 128138, 128137, 127777, 128701, 128676, 9978, 128025, 129414, 129417, 129415, 128052, 128041, 129345, 127923, 128515, 128519, 128569, 127795, 128539, 128642, 128066, 128745, 127890, 127798, 127834, 9000, 127814, 128202, 127869, 128267, 128452, 128203, 128478, 128250, 127909, 128191, 128234, 128004, 128019, 129362, 129372, 129366, 127874, 127851, 128084, 128087, 127803, 9968, 127881, 128230, 128063, 128110, 128119, 127183, 128091, 128026, 127809, 128035, 127859, 127914, 129354, 9881, 128197}

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
