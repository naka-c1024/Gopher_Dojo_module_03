package fortune_test

import (
	"omikuji/fortune"
	"testing"
)

func TestRtnHealth(t *testing.T) {
	cases := []struct {
		name     string // サブテストの名前
		input    int    // 関数に渡すもの
		expected string // 期待するもの
	}{
		{name: "0", input: 0, expected: "Very good. It will improve with exercise."},
		{name: "1", input: 1, expected: "Good."},
		{name: "2", input: 2, expected: "Good."},
		{name: "3", input: 3, expected: "No problems."},
		{name: "4", input: 4, expected: "No problems. However, do not binge eat or drink."},
		{name: "5", input: 5, expected: "Try to make sure your body does not become cold."},
		{name: "6", input: 6, expected: "Your worries will continue. Sometimes, you just need to rest your mind."},
		{name: "7", input: 7, expected: "??"},
	}
	for _, c := range cases {
		c := c // クロージャが順番に呼ばれない可能性もあるため、このスコープ内でのローカル変数をちゃんと定義する
		t.Run(c.name,
			func(t *testing.T) {
				if actual := fortune.RtnHealth(c.input); actual != c.expected {
					t.Errorf("want IsEven(%d) = %v, but actual = %v", c.input, c.expected, actual)
				}
			})
	}
}

func TestRtnWish(t *testing.T) {
	cases := []struct {
		name     string // サブテストの名前
		input    int    // 関数に渡すもの
		expected string // 期待するもの
	}{
		{name: "0", input: 0, expected: "It will definitely come true."},
		{name: "1", input: 1, expected: "It shall come true without a hitch."},
		{name: "2", input: 2, expected: "Teamwork is the key to success. If you work together, it shall come true."},
		{name: "3", input: 3, expected: "If you take action, it will come true."},
		{name: "4", input: 4, expected: "It will take time, but it will happen."},
		{name: "5", input: 5, expected: "If you desire it, it will come true. However, this will come with much difficulty."},
		{name: "6", input: 6, expected: "It will be difficult for it to come true."},
		{name: "7", input: 7, expected: "??"},
	}
	for _, c := range cases {
		c := c // クロージャが順番に呼ばれない可能性もあるため、このスコープ内でのローカル変数をちゃんと定義する
		t.Run(c.name,
			func(t *testing.T) {
				if actual := fortune.RtnWish(c.input); actual != c.expected {
					t.Errorf("want IsEven(%d) = %v, but actual = %v", c.input, c.expected, actual)
				}
			})
	}
}

func TestRtnFortune(t *testing.T) {
	cases := []struct {
		name     string // サブテストの名前
		input    int    // 関数に渡すもの
		expected string // 期待するもの
	}{
		{name: "0", input: 0, expected: "Dai-kichi"},
		{name: "1", input: 1, expected: "Kichi"},
		{name: "2", input: 2, expected: "Chuu-kichi"},
		{name: "3", input: 3, expected: "Sho-kichi"},
		{name: "4", input: 4, expected: "Sue-kichi"},
		{name: "5", input: 5, expected: "Kyo"},
		{name: "6", input: 6, expected: "Dai-kyo"},
		{name: "7", input: 7, expected: "??"},
	}
	for _, c := range cases {
		c := c // クロージャが順番に呼ばれない可能性もあるため、このスコープ内でのローカル変数をちゃんと定義する
		t.Run(c.name,
			func(t *testing.T) {
				if actual := fortune.RtnFortune(c.input); actual != c.expected {
					t.Errorf("want IsEven(%d) = %v, but actual = %v", c.input, c.expected, actual)
				}
			})
	}
}
