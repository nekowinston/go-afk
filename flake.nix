{
  description = "afk-go flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages = rec {
          afk-go = pkgs.buildGoModule rec {
            name = "afk-go";
            pname = "afk";
            version = "0.1.0";
            src = ./.;
            vendorSha256 = "sha256-EYBABMCGflg34QAZhm2n1ZYXBqSE7/skb0rPA/xxGN8=";
            ldflags = [
              "-s -w -X github.com/nekowinston/afk-go.version=${version}"
            ];
            CGO_ENABLED = 0;

            meta = with pkgs.lib; {
              description = "afk-go";
              homepage = "https://github.com/nekowinston/afk-go";
              license = licenses.mit;
            };
          };
          default = afk-go;
        };
        devShells.default = pkgs.mkShell {buildInputs = with pkgs; [go];};
        apps = rec {
          afk-go = flake-utils.lib.mkApp {drv = self.packages.${system}.afk-go;};
          default = afk-go;
        };
      }
    );
}
