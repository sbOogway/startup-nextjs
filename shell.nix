{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.nodejs
    pkgs.typescript
    pkgs.yarn
  ];

  shellHook = ''
    echo "Welcome to the TypeScript development environment!"
    echo "You can run 'tsc' to compile TypeScript files."
    echo "You can run 'node <your-js-file.js>' to execute JavaScript files."
  '';
}
