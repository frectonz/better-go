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
        packages.default = go.overrideAttrs (final: prev: {
          src = fetchFromGitHub {
            owner = "frectonz";
            repo = "better-go";
            rev = "second-try";
            hash = "sha256-PnkYOvr8+nEv15nv6W+F1wdwzXdtgXuJgyZQTz48duU=";
          };
        });

        devShells.default = mkShell.override { stdenv = clangStdenv; } {
          buildInputs = [ go gotools ];
        };

        formatter = nixpkgs-fmt;
      });
}
