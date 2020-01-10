class Mytool < Formula
  desc "Some description"
  homepage "https://github.com/myrepo/mytool"
  url ".../mytool-darwin-amd64.tar.gz", :using => GitHubPrivateRepositoryReleaseDownloadStrategy
  sha256 "<SHA256 CHECKSUM>"
  head "https://github.com/myrepo/mytool.git"

  def install
    bin.install "mytool"
  end

  # Homebrew requires tests.
  test do
    assert_match "mytool version 1.0.0", shell_output("#{bin}/mytool -v", 2)
  end
end
