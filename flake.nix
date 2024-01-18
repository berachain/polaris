{
    description = "An EVM framework for Cosmos";

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
        flake-utils.url = "github:numtide/flake-utils";
        foundry.url = "github:shazow/foundry.nix/monthly";
    };


    outputs = { self, nixpkgs, flake-utils, foundry, ... }:
        flake-utils.lib.eachDefaultSystem (system:
            let
                pkgs = import nixpkgs {
                    inherit system;
                    overlays = [ foundry.overlay ];
                };
            in {
                devShell = with pkgs; mkShell {
                    name = "polaris-dev";
                    nativeBuildInputs = [
                        go
                        jq
                        foundry-bin
                    ];


                    shellHook = ''
                        export PS1="[dev] $PS1"
                    '';
                };
            }
        );
}