# clean-architecture-idea

アイデアとか言いながら、Clean ArchitectureのFlow of control を実装してみただけという・・・
気が向いたらテスト追加したり適当なGoのAPIフレームワーク(EchoとかGin)でいい感じにいけそうか試したりする。

2021/04/20 追記
--
Echoで試してみたが、Flow of controlの流れを守るのであればだいぶ厳しそう。
Echoに慣れてる人は普通`c.JSON`を使ってレスポンスを返すと思うが、
Presenterにリクエストスコープの`echo.Context`を渡す場合に微妙な方法しかない気がする。
