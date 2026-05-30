package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGitRootFromSubdirectory(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "internal", "config")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := FindGitRoot(sub)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := filepath.Abs(root)
	if got != want {
		t.Fatalf("FindGitRoot(%q) = %q, want %q", sub, got, want)
	}
}

func TestFindGitRootFallbackNoGit(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "nested")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := FindGitRoot(sub)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := filepath.Abs(sub)
	if got != want {
		t.Fatalf("FindGitRoot(%q) = %q, want %q", sub, got, want)
	}
}

func TestFindGitRootGitfile(t *testing.T) {
	root := t.TempDir()
	sub := filepath.Join(root, "pkg")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".git"), []byte("gitdir: /fake/worktree"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := FindGitRoot(sub)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := filepath.Abs(root)
	if got != want {
		t.Fatalf("FindGitRoot(%q) = %q, want %q", sub, got, want)
	}
}
