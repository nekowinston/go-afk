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
          go-afk = pkgs.buildGo120Module rec {
            name = "go-afk";
            version = "0.2.0";
            src = ./.;
            vendorSha256 = "sha256-UCqIn5+p2Zm7B8In5OvCVo2PvWaWLpcLa+Ya0sTiSaw=";
            ldflags = [
              "-s -w -X github.com/nekowinston/go-afk.version=${version}"
            ];
            CGO_ENABLED = 0;

            meta = with pkgs.lib; {
              description = "go-afk";
              homepage = "https://github.com/nekowinston/afk-go";
              license = licenses.mit;
            };
          };
          default = go-afk;
        };
        devShells.default = pkgs.mkShell {buildInputs = with pkgs; [go];};
        apps = rec {
          go-afk = flake-utils.lib.mkApp {drv = self.packages.${system}.go-afk;};
          default = go-afk;
        };
      }
    );
}
