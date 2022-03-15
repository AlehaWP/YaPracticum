package defoptions

import (
	"testing"
)

func TestNewDefOptions(t *testing.T) {
	t.Run("Тест DefOptions", func(t *testing.T) {
		got := NewDefOptions()
		if len(got.ServAddr()) == 0 {
			t.Errorf("NewDefOptions().ServAddr())= %v, want len > 0", got.ServAddr())
		}
		if len(got.DBConnString()) == 0 {
			t.Errorf("NewDefOptions().DBConnString())= %v, want len > 0", got.DBConnString())
		}
		if len(got.RepoFileName()) == 0 {
			t.Errorf("NewDefOptions().RepoFileName())= %v, want len > 0", got.RepoFileName())
		}
		if len(got.RespBaseURL()) == 0 {
			t.Errorf("NewDefOptions().RespBaseURL())= %v, want len > 0", got.RespBaseURL())
		}

	})
}
