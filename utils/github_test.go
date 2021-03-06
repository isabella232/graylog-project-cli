package utils

import "testing"

func TestParseGitHubURL(t *testing.T) {
	var url GitHubURL
	var expected string
	var err error

	// Custom URL
	url, _ = ParseGitHubURL("github://Graylog2/graylog2-server.git")
	expected = "git@github.com:Graylog2/graylog2-server.git"
	if url.SSH() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.SSH())
	}

	url, _ = ParseGitHubURL("github://Graylog2/graylog2-server.git")
	expected = "https://github.com/Graylog2/graylog2-server.git"
	if url.HTTPS() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.HTTPS())
	}

	// Git URL
	url, _ = ParseGitHubURL("git@github.com:Graylog2/graylog2-server.git")
	expected = "git@github.com:Graylog2/graylog2-server.git"
	if url.SSH() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.SSH())
	}

	url, _ = ParseGitHubURL("git@github.com:Graylog2/graylog2-server.git")
	expected = "https://github.com/Graylog2/graylog2-server.git"
	if url.HTTPS() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.HTTPS())
	}

	// HTTPS URL
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	expected = "git@github.com:Graylog2/graylog2-server.git"
	if url.SSH() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.SSH())
	}

	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	expected = "https://github.com/Graylog2/graylog2-server.git"
	if url.HTTPS() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.HTTPS())
	}

	// Directory
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	expected = "graylog2-server"
	if url.Directory() != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url.HTTPS())
	}

	// Matches ok
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	match := "Graylog2/graylog2-server"
	if !url.Matches(match) {
		t.Errorf("expected <%s> to match <%s>", url, match)
	}

	// Matches case
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	match = "GraYloG2/GraYlog2-SErver"
	if !url.Matches(match) {
		t.Errorf("expected <%s> to match <%s>", url, match)
	}

	// Match fails
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	match = "Graylog2/graylog2-server-does-not-work"
	if url.Matches(match) {
		t.Errorf("expected <%s> to not match <%s>", url, match)
	}

	// Missing .git suffix
	_, err = ParseGitHubURL("https://github.com/Graylog2/graylog2-server")
	if err == nil {
		t.Error("expected URL without .git suffix to fail")
	}

	// Unknown URL format
	_, err = ParseGitHubURL("https://example.com/Graylog2/graylog2-server")
	if err == nil {
		t.Error("expected unknown URL to fail")
	}

	url, _ = ParseGitHubURL("github://Graylog2/graylog2-server.git")
	if url.IsHTTPS() || url.IsHTTPS() {
		t.Error("expected URL to not be SSH or HTTPS")
	}
	url, _ = ParseGitHubURL("https://github.com/Graylog2/graylog2-server.git")
	if !url.IsHTTPS() || url.IsSSH() {
		t.Error("expected URL to be HTTPS and not SSH")
	}
	url, _ = ParseGitHubURL("git@github.com:Graylog2/graylog2-server.git")
	if url.IsHTTPS() || !url.IsSSH() {
		t.Error("expected URL to be SSH and not HTTPS")
	}
}

func TestReplaceGitHubURL(t *testing.T) {
	url, _ := ReplaceGitHubURL("github://Graylog2/graylog2-server.git", "foo/graylog2-server")
	expected := "github://foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}
	url, _ = ReplaceGitHubURL("github://Graylog2/graylog2-server.git", "foo/graylog2-server.git")
	expected = "github://foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}

	url, _ = ReplaceGitHubURL("https://github.com/Graylog2/graylog2-server.git", "foo/graylog2-server")
	expected = "https://github.com/foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}
	url, _ = ReplaceGitHubURL("https://github.com/Graylog2/graylog2-server.git", "foo/graylog2-server.git")
	expected = "https://github.com/foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}

	url, _ = ReplaceGitHubURL("https://github.com/Graylog2/graylog2-server.git", "foo/graylog2-server")
	expected = "https://github.com/foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}
	url, _ = ReplaceGitHubURL("https://github.com/Graylog2/graylog2-server.git", "foo/graylog2-server.git")
	expected = "https://github.com/foo/graylog2-server.git"
	if url != expected {
		t.Errorf("expected <%s> but got <%s>", expected, url)
	}
}

func TestParseGitHubPRString(t *testing.T) {
	var cases = []struct {
		input    string
		prRepo   string
		prNumber int
		err      bool
	}{
		{"Graylog2/graylog2-server#123", "Graylog2/graylog2-server", 123, false},
		{"https://github.com/Graylog2/graylog-plugin-collector/pull/9692", "Graylog2/graylog-plugin-collector", 9692, false},
		{"https://github.com/Graylog2/graylog-plugin-collector/pull/", "", 0, true},
		{"https://github.com/9692", "", 0, true},
		{"https://example.com/Graylog2/graylog-plugin-collector/pull/9692", "", 0, true},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			repo, num, err := ParseGitHubPRString(c.input)
			if repo != c.prRepo {
				t.Errorf("expected <%s>, got <%s>", c.prRepo, repo)
			}
			if num != c.prNumber {
				t.Errorf("expected <%d>, got <%d>", c.prNumber, num)
			}
			if err == nil && c.err {
				t.Errorf("expected an error but got none")
			}
			if err != nil && !c.err {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}
