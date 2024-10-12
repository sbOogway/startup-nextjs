{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go               # Go programming language
    pkgs.git              # Git for version control
    pkgs.golangci-lint    # Linter for Go code
    pkgs.gopls            # Language server for Go (optional)
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH
    mkdir -p $GOPATH

    if [ ! -f go.mod ]; then
      go mod init myproject  # Replace 'myproject' with your module name
    fi

    # Add Gin as a dependency
    go get -u github.com/gin-gonic/gin
    go get -u github.com/go-sql-driver/mysql
  '';
}
