# Shamoji
#### - All the security of SHA256 in 26 emoji 

## Install

```shell
brew install quackduck/tap/shamoji
```
or
```shell
git clone https://github.com/quackduck/shamoji && cd shamoji
go install
```

## Use

### Hash things
```shell
$ echo "hello" | shamoji
๐ฆ๐ฅญ๐ช๐๐๐ฃ๐๐ฅ๐๐ฅ๐๐ฅ๐ฅช๐๐๐ง๐๐ฆ๐๐ช๐๐๐ฆ๐ธ๐คบ๐ 
$ echo "ishan" | shamoji
๐ฆ๐ป๐๐๐ฝ๐ค๐น๐ฉ๐ง๐๐ป๐๐ฐ๐๐๐ฑ๐๐ ๐ฅโฒ๐ด๐ฆฉ๐ง๐ฆป๐ฏ๐ฃ
$ cat main.go | shamoji
๐ฆ๐ต๐งบ๐๐๐๐๐ฌ๐ฃ๐๐ง๐๐ฎ๐๐๐ฅ๐ชฑ๐๐๐ด๐๐ท๐น๐โธ๐
```

### See shamojis instead of regular hashes
```shell
$ git log | shamoji --replace
commit ๐ฅผ๐๐ฟ๐ฅ๐โกโซ๐ฅ๐๐๐๐ ๐ฟ๐ค๐ชฒ๐ง
Author: Ishan Goel <igoel.mail@gmail.com>
Date:   Mon Jun 13 08:32:38 2022 -0700

    some stuff

commit ๐คฟ๐๐ฆผ๐๐๐๐ฉ๐๐๐ก๐บ๐๐๐คง๐๐ค
Author: Ishan Goel <igoel.mail@gmail.com>
Date:   Sat Mar 26 15:13:26 2022 -0700

    Init
```

