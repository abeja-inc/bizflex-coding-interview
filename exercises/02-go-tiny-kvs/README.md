# Technical Interview Exercises - Tiny KVS (Go)

これは、[株式会社 ABEJA](https://abejainc.com/ja/) の開発者採用面談で利用するコーディングテストの演習問題のひとつです。

## 制限事項

以下の制限事項を守ってください。

1. Go 1.16 以降をお使いください
2. Go の標準ライブラリ以外のパッケージ/モジュールを利用しないでください（テストは除く）

## 問題 1

このディレクトリにある `kvs.go` を編集して、簡易的な Key Value Store (以下「KVS」) を実装してください。この KVS に格納できる値は `int` 型の値のみとし、以下のインターフェースを満たす実装とします。

```go
type KVS interface {
	Insert(key string, value int)
	Count() int
	Search(key string) []int
	PrefixSearch(prefix string) []int
}
```

- **Insert** - 文字列のキー `key` に対応する値として `value` を登録します。すでに対応する値が存在する場合も**上書きはせず**、すべての値を格納します。
- **Count** - 格納されている値の個数を返します。キーの個数ではないことに注意してください。
- **Search** - 文字列のキー `key` に対応するすべての値へのスライスを返します。値がひとつも存在しない場合は、空のスライスを返します。スライス中の値の順番に決まりはありません。
- **PrefixSearch** - 文字列 `prefix` と前方一致するすべてのキーに対応するすべての値へのスライスを返します。値がひとつも存在しない場合は、空のスライスを返します。スライス中の値の順番に決まりはありません。

`kvs.go` にダミーの実装がありますので、こちらを編集して実装を完成させてください。

## 問題 2

問題 1 を素直に実装すると、KVS に登録した件数に比例して、`PrefixSearch()` 関数の処理に時間がかかるようになってしまいます。たとえば、次のようなコードがあった場合に、

```go
kvs := NewKVS()

// (1) ここで数千〜数百万件のレコード追加
// ...

// (2) 共通の接頭辞を持つ単語をいくつか登録
kvs.Insert("honey", 10)
kvs.Insert("honeycomb", 20)
kvs.Insert("honeybee", 30)
kvs.Insert("honeymoon", 40)
kvs.Insert("honesty", 50)

// (3) 共通の接頭辞で検索
kvs.PrefixSearch("honey")
kvs.PrefixSearch("hone")
kvs.PrefixSearch("ho")
kvs.PrefixSearch("h")
kvs.PrefixSearch("bee")
```

- (1) で登録されるレコードの件数が数千〜数百万、あるいはそれ以上のオーダーで増えても (3) で実行する `PrefixSearch()` の速度が影響を受けないように改良してください。
- ただし、(1) で登録するレコードのキーは、(2) で登録されるレコードのキーと共通の接頭辞を持たないものとします。

なお、上記シナリオを模倣したベンチマークが `kvs_test.go` にあり、`make bench` で実行できます。

## テスト

`kvs_test.go` にテストが記述されており、同じディレクトリで `make` を実行するとテストが実行されるので、このテストがすべてパスするようにしてください。

```bash
$ make test
go test -v ./...
=== RUN   TestCount
=== RUN   TestCount/empty
=== RUN   TestCount/one
...
PASS
ok  	abeja-inc/tiny-kvs	9.783s
```
