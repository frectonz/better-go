{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=master";
    utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { self
    , nixpkgs
    , utils
    }:
    utils.lib.eachDefaultSystem
      (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      with pkgs; {
        devShells.default = mkShell.override { stdenv = clangStdenv; } {
          buildInputs = [ go gotools ];
        };

        formatter = nixpkgs-fmt;
      });
}
