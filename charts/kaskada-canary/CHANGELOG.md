# Kaskada Helm Chart Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

<!-- ## [vX.Y.Z] - UNRELEASED
### Highlights
### All Changes
- Added
- Updated
- Changed
- Fixed
- Deprecated
- Removed -->

## [v1.8.0] - 2022-04-13

### Added

- Add annotations to Deployment. ([#2477](https://github.com/kubernetes-sigs/external-dns/pull/2477)) [@beastob](https://github.com/beastob)

### Changed

- Fix RBAC for `istio-virtualservice` source when `istio-gateway` isn't also added. ([#2564](https://github.com/kubernetes-sigs/external-dns/pull/2564)) [@mcwarman](https://github.com/mcwarman)
