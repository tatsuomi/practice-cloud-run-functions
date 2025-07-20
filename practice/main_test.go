package main

import "testing"

// Hello: Hello World 実行用関数
func TestHello(t *testing.T) {
	type testConfig struct {
		name string
		arg  string
		want string
	}
	testCase := []testConfig{
		{
			name: "Hello World Test",
			arg:  "World",
			want: "Hello World",
		},
		{
			name: "Hello Taro Test",
			arg:  "Taro",
			want: "Hello Taro",
		},
	}
	// テストの実行
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// 実行ケースの表示
			t.Logf("Running test: %s\n", tc.name)
			// 関数の実行と結果の検証
			if got := Hello(tc.arg); got != tc.want {
				t.Errorf("Hello(%s) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}
