package chrono

import (
	"testing"
	"time"

	"github.com/zeroberto/go-ms-template/chrono"
	"github.com/zeroberto/go-ms-template/chrono/provider"
)

func TestGetCurrentTime(t *testing.T) {
	expectedLessOrEqual := time.Now()

	var ts chrono.TimeStamp = &provider.TimeStampImpl{}

	got := ts.GetCurrentTime()

	if expectedLessOrEqual.After(got) {
		t.Errorf("GetCurrentTime() failed, expected bigger or equal %v, got %v", expectedLessOrEqual, got)
	}
}
