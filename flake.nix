{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    devenv.url = "github:cachix/devenv/9ba9e3b908a12ddc6c43f88c52f2bf3c1d1e82c1";
    flake-utils.url = "github:numtide/flake-utils";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    self,
    nixpkgs,
    devenv,
    flake-utils,
    ...
  } @ inputs:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
          config = {
          };
        };
      in {
        packages.default = pkgs.callPackage (
          {
            lib,
            fetchFromGitHub,
            buildGoModule,
          }:
            buildGoModule rec {
              pname = "ionscale";
              version = "0.16.0";

              src = ./.;
              vendorHash = "sha256-UzxfIaZ2tbCt0g4WtH0gnSw8HGFI+07JOe4HUdPQmqs=";

              ldflags = [
                "-s"
                "-w"
                "-X github.com/jsiebens/ionscale/internal/version.Version=${version}"
              ];

              doCheck = false;

              meta = {
                description = "A lightweight implementation of a Tailscale control server";
                homepage = "https://jsiebens.github.io/ionscale/";
                license = lib.licenses.bsd3;
                mainProgram = "ionscale";
              };
            }
        ) {};

        devShell = devenv.lib.mkShell {
          inherit inputs pkgs;
          modules = [
            ({pkgs, ...}: let
              old_pkgs = import (builtins.fetchTarball {
                url = "https://github.com/NixOS/nixpkgs/archive/336eda0d07dc5e2be1f923990ad9fdb6bc8e28e3.tar.gz";
              }) {};

              connect-pkgs = import (builtins.fetchTarball {
                url = "https://github.com/NixOS/nixpkgs/archive/976fa3369d722e76f37c77493d99829540d43845.tar.gz";
              }) {};
            in {
              packages = with pkgs; [
                go
                gnumake

                old_pkgs.buf
                old_pkgs.protoc-gen-go
                connect-pkgs.protoc-gen-connect-go
                old_pkgs.templ
                old_pkgs.grpcurl
              ];

              env = {
              };
            })
          ];
        };
      }
    );
}
