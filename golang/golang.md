## Column

## marshal, unmarshalの継承

```go
func (レシーバー) MarshalJSON() ([]byte, error) {}
func (レシーバー) UnmarshalJSON(data []byte) error {}
```

- cobra コマンドラインツール生成
- パッケージ構成 https://qiita.com/rema424/items/9ffbdf584b705cae6a19
- https://ponde-m.hatenablog.com/entry/2019/03/24/230333
- wire (DI実現)
- internalはpkgからは読めない
- marshal, unmarshal継承

## logパッケージのFatalについて
- Fatalを使うのはメイン関数だけに絞るほうが良い
- deferが実行されなかったりする
- ァイルのクローズとかは問題ないですが、deferでredisのロックの解除とかやってるとハマります

## 豆
- echo $? でプログラムが正常に終わったか調べれる
- 文字列の連続を調べる時はprefixでも良さげ

## sqlで取りづらい時
- map[int]entityでid:entityのmapを作ることを考えてみる

## 適当な連番

- https://golang.org/pkg/crypto/rand/#Read

## いかつい型アサーション?

```go
type Date time.Time

date := &Date{}
time := (*time.Time)(date)
```
型キャストですね。ポインタじゃなかったら、  
time.Time(charge.Birthdate)  
と書けるところを  
*time.Time(charge.Birthdate)  
と書くと  
*(time.Time(charge.Birthdate))  
と解釈されてしまうので、カッコを付けて*も型の一部であることを表記してるんだと思います。

## interface型アサーション

```go
type Date interface {
  Now() error
}

type Date1 struct {
}

type Date2 struct {
}

func (d *Date1) Now() error {
  return nil
}

func (d *Date2) Now() error {
  return nil
}

var date Date

date1 := date.(Date1)
date2 := date.(Date2)
```