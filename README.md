# Shamoji
#### - All the security of SHA256 in 26 emoji 

## Install
```shell
git clone https://github.com/quackduck/shamoji && cd shamoji
go install
```

## Use

### Hash things
```shell
$ echo "hello" | shamoji
🦆🥭🪗🚝🐉🟣😑🥞🖇🥚🆗🥜🥪🙀🙎🧏💊🦂💁🎪💑👖🐦🏸🤺👠
$ echo "ishan" | shamoji
🦆🎻😊🆒😽🍤👹🛩🧘🐉📻🕟🛰🚐🏕🅱🍜🕠🟥⏲😴🦩🧞🦻📯🕣
$ cat main.go | shamoji
🦆🍵🧺🎒😀👑🛌🏬🏣😋🧟🍋🎮🔝🏋🥅🪱🚛💗🔴🔉🐷🎹🚋⛸🍒
```

### See shamojis instead of regular hashes
```shell
$ git log | shamoji --replace
commit 🥼🆕🌿📥📛⚡⏫🥎🍘💚🎆🎠💿🤜🪲🧍
Author: Ishan Goel <igoel.mail@gmail.com>
Date:   Mon Jun 13 08:32:38 2022 -0700

    some stuff

commit 🤿💃🦼🔜🚎👊🌩📅🍎💡🛺🛂👘🤧🙁🌤
Author: Ishan Goel <igoel.mail@gmail.com>
Date:   Sat Mar 26 15:13:26 2022 -0700

    Init
```

