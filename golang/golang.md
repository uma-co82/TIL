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