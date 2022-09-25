package fortune_test

import (
	"net/http"
	"net/http/httptest"
	"omikuji/fortune"
	"testing"
	"time"
)

func TestIsOshougatsu(t *testing.T) {
	// TimeNow のバックアップと defer でリカバー
	oldTimeNow := fortune.TimeNow
	defer func() { fortune.TimeNow = oldTimeNow }()

	cases := []struct {
		month    int
		day      int
		expected bool
	}{
		{month: 1, day: 1, expected: true},
		{month: 1, day: 2, expected: true},
		{month: 1, day: 3, expected: true},
		{month: 1, day: 4, expected: false},
	}

	for _, c := range cases {
		fortune.TimeNow = time.Date(2022, time.Month(c.month), c.day, 23, 59, 59, 0, time.Local)
		actual := fortune.IsOshougatsu()
		if actual != c.expected {
			t.Errorf("Fail assert equal. Expect: %v Actual: %v", c.expected, actual)
		}
	}
}

func TestOmikuji(t *testing.T) {
	if rdmNum := fortune.Omikuji(); !(0 <= rdmNum && rdmNum <= 6) {
		t.Errorf("Over ranges from 1 to 150")
	}
}

func TestHandler(t *testing.T) {
	// TimeNow のバックアップと defer でリカバー
	oldTimeNow := fortune.TimeNow
	defer func() { fortune.TimeNow = oldTimeNow }()
	fortune.TimeNow = time.Date(2022, time.Month(1), 1, 23, 59, 59, 0, time.Local)

	w := httptest.NewRecorder()                        // 擬似的なhttp.ResponseWriterを返す
	r := httptest.NewRequest(http.MethodGet, "/", nil) // 擬似的なリクエストを作る
	fortune.Handler(w, r)                              // ここで実行
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
}
