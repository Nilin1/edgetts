package edgetts

import (
	"os"
	"testing"
	"time"
)

func TestSpeech_Minimal(t *testing.T) {
	text := "种一棵树最好的时间是十年前，其次是现在。"

	opts := []Option{
		WithVoice("zh-CN-YunxiaNeural"),
	}

	speech, err := NewSpeech(opts...)
	if err != nil {
		t.Fatalf("create speech failed: %v", err)
	}

	f, err := os.Create("testdata/output.mp3")
	if err != nil {
		t.Fatalf("create file failed: %v", err)
	}
	defer f.Close()

	if err := speech.AddSingleTask(text, f); err != nil {
		t.Fatalf("add task failed: %v", err)
	}

	start := time.Now()
	if err := speech.StartTasks(); err != nil {
		t.Fatalf("start task failed: %v", err)
	}

	t.Logf("done, cost=%v", time.Since(start))
}
