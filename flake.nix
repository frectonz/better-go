{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=master";
    utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
  };

  outputs =
    { self
    , utils
    , nixpkgs
    , nix-filter
    }:
    utils.lib.eachDefaultSystem
      (system:
      let
        pkgs = import nixpkgs { inherit system; };
        filter = nix-filter.lib;
      in
      with pkgs; {
        packages.default = go.overrideAttrs (final: prev: {
          src = filter {
            root = ./.;
            exclude = [
              ./flake.nix
              ./flake.lock
            ];
          };
        });

        devShells.default = mkShell.override { stdenv = clangStdenv; } {
          buildInputs = [ go gotools ];
        };

        formatter = nixpkgs-fmt;
      });
}
