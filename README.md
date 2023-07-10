# stream

* [SICP 3章のストリーム](https://sicp.iijlab.net/fulltext/x350.html)の一部をgoで実装したものです
  * goである必然性はありません
  * Lispが読めない人でも比較的読みやすいと思ったからというのが理由の一つです

## 見どころ

自然数全体を、定数ゼロとsuccのみで構成しているところとか

```go
var Natural = IntegerStartingFrom(0)

func IntegerStartingFrom(n int) Stream[int] {
	return Cons(
		n,
		func() Stream[int] {
			return IntegerStartingFrom(n + 1)
		},
	)
}
```

フィボナッチ数全体を、以下の2つで構成しているところとか

* 無限に長いかもしれないストリーム2つを足したストリームを足したストリームを得るAddStream
* fibsとfibsをずらしたfibs.Cdr()


```go
func GenerateFib2() Stream[int] {
	return Cons(
		0,
		func() Stream[int] {
			return Cons(
				1,
				func() Stream[int] {
					fibs := GenerateFib2()
					return AddStream(fibs.Cdr(), fibs)
				},
			)
		},
	)
}
```
