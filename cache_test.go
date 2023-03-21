package gocache

import (
	"errors"
	"testing"
	"time"

	"github.com/tidwall/buntdb"
)

func TestNew(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer c.Close()

	if c == nil {
		t.Fatal("expected a new cache instance, got: nil")
	}
}

func TestClose(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	err = c.Close()
	if err != nil {
		t.Fatalf("expected no error on close, got: %v", err)
	}
}

func TestSet(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer c.Close()

	err = c.Set("key", "value", 5*time.Second)
	if err != nil {
		t.Fatalf("expected no error on set, got: %v", err)
	}
}

func TestGet(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer c.Close()

	err = c.Set("key", "value", 5*time.Second)
	if err != nil {
		t.Fatalf("expected no error on set, got: %v", err)
	}

	val, err := c.Get("key")
	if err != nil {
		t.Fatalf("expected no error on get, got: %v", err)
	}

	if val != "value" {
		t.Fatalf("expected value to be 'value', got: %v", val)
	}
}

func TestGetNonExistentKey(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	defer c.Close()

	_, err = c.Get("nonexistent")
	if !errors.Is(err, buntdb.ErrNotFound) {
		t.Fatalf("expected error to be buntdb.ErrNotFound, got: %v", err)
	}
}
