package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("Expected new stack to be non-nil")
	}
	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}
}

func TestPushAndItems(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	items := s.Items()
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}
	if items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Errorf("Unexpected items: %v", items)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(20)
	item, ok := s.Pop()
	if !ok || item != 20 {
		t.Errorf("Expected to pop 20, got %v", item)
	}
	item, ok = s.Pop()
	if !ok || item != 10 {
		t.Errorf("Expected to pop 10, got %v", item)
	}
	_, ok = s.Pop()
	if ok {
		t.Error("Expected pop on empty stack to return false")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack()
	_, ok := s.Peek()
	if ok {
		t.Error("Expected peek on empty stack to return false")
	}
	s.Push("a")
	item, ok := s.Peek()
	if !ok || item != "a" {
		t.Errorf("Expected to peek 'a', got %v", item)
	}
	if s.Size() != 1 {
		t.Error("Peek should not remove item from stack")
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack()
	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Error("Expected stack to not be empty after push")
	}
}

func TestSize(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Error("Expected size 0 for new stack")
	}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
}

func TestItemsAfterPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Pop()
	items := s.Items()
	if len(items) != 1 || items[0] != 1 {
		t.Errorf("Expected [1], got %v", items)
	}
}
