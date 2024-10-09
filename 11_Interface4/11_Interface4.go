package main

import "fmt"

// ポニーは歩ける
type Pony interface {
	Walk()
	Sprint()
}

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// アースポニーも歩けるので、Ponyインタフェースに渡せる
type EarthPony struct {
}

func (ep *EarthPony) Walk() {
	fmt.Println("EarthPony 歩くよ")
}

// 走るよ！走るのは歩くのを3倍ぐらい頑張ってるよ
func (ep *EarthPony) Sprint() {
	ep.Walk()
	ep.Walk()
	ep.Walk()
}

// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// ペガサスはアースポニーの上位互換
type Pegasus struct {
	EarthPony
}

func (u *Pegasus) Fly() {
	fmt.Println("Pegasus 飛ぶよ")
}

func (u *Pegasus) Walk() {
	u.Fly() // 歩けと言われても、ペガサスは飛ぶよ！
}

// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

func main() {
	Pon := EarthPony{}
	Peg := Pegasus{}

	Pon.Walk()
	Pon.Sprint()
	Peg.Fly()
	Peg.Walk()
	Peg.Sprint()

}
