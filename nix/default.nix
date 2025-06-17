{
  sources ? import ./sources.nix,
  system ? builtins.currentSystem,
  ...
}:
import sources.nixpkgs {
  overlays = [
    (_: pkgs: {
      flake-compat = import sources.flake-compat;
    })
    (import "${sources.gomod2nix}/overlay.nix")
  ];
  config = { };
  inherit system;
}
