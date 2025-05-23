# Changelog

## [1.0.0](https://github.com/hugginsio/kv2/compare/v0.11.6...v1.0.0) (2025-05-03)


### ⚠ BREAKING CHANGES

* revise docs for 1.0 ([#63](https://github.com/hugginsio/kv2/issues/63))

### Features

* add audit logs to server ([#61](https://github.com/hugginsio/kv2/issues/61)) ([e763ee9](https://github.com/hugginsio/kv2/commit/e763ee910404df2b31764638ae4b0abaee3e8cac)), closes [#39](https://github.com/hugginsio/kv2/issues/39)
* revise docs for 1.0 ([#63](https://github.com/hugginsio/kv2/issues/63)) ([3893b93](https://github.com/hugginsio/kv2/commit/3893b93658ec6b73c8abfc7f8382b796179f8830))

## [0.11.6](https://github.com/hugginsio/kv2/compare/v0.11.5...v0.11.6) (2025-04-20)


### Bug Fixes

* goreleaser tap, add missing CLI docs ([#59](https://github.com/hugginsio/kv2/issues/59)) ([da9c2ee](https://github.com/hugginsio/kv2/commit/da9c2eefe3a12a93d5849859ece8e650871ee9ab))

## [0.11.5](https://github.com/hugginsio/kv2/compare/v0.11.4...v0.11.5) (2025-04-20)


### Bug Fixes

* **release:** allow app token to access tap repo ([#57](https://github.com/hugginsio/kv2/issues/57)) ([8fcd416](https://github.com/hugginsio/kv2/commit/8fcd416d17f55d9d2618dc0a75e2d0d23ac5465f))

## [0.11.4](https://github.com/hugginsio/kv2/compare/v0.11.3...v0.11.4) (2025-04-20)


### Bug Fixes

* do not use app token for GHCR ([#55](https://github.com/hugginsio/kv2/issues/55)) ([503a2c5](https://github.com/hugginsio/kv2/commit/503a2c5da219870c130ff82a6176efcd0738ab4f))

## [0.11.3](https://github.com/hugginsio/kv2/compare/v0.11.2...v0.11.3) (2025-04-20)


### Bug Fixes

* use app token for release automation ([#53](https://github.com/hugginsio/kv2/issues/53)) ([8b9dc22](https://github.com/hugginsio/kv2/commit/8b9dc22010107aed0644740e19e147c2205de448))

## [0.11.2](https://github.com/hugginsio/kv2/compare/v0.11.1...v0.11.2) (2025-04-20)


### Bug Fixes

* **cli:** mount token for goreleaser ([#51](https://github.com/hugginsio/kv2/issues/51)) ([36c70ee](https://github.com/hugginsio/kv2/commit/36c70eefb23ff48876b1de4a67d270284051df0c))

## [0.11.1](https://github.com/hugginsio/kv2/compare/v0.11.0...v0.11.1) (2025-04-20)


### Bug Fixes

* **cli:** release behavior ([#49](https://github.com/hugginsio/kv2/issues/49)) ([48a44d5](https://github.com/hugginsio/kv2/commit/48a44d5b1090dfde70eb3209ca118b5298bde0af))

## [0.11.0](https://github.com/hugginsio/kv2/compare/v0.10.0...v0.11.0) (2025-04-20)


### Features

* add endpoint for better version command ([#46](https://github.com/hugginsio/kv2/issues/46)) ([df60d1d](https://github.com/hugginsio/kv2/commit/df60d1d65e894bc52551bcd697169ebb0910ffd1))
* **cli:** prompt user if no flags provided to create ([#41](https://github.com/hugginsio/kv2/issues/41)) ([09cd196](https://github.com/hugginsio/kv2/commit/09cd196310fe6901ff81dc4c9d579d91c8711723))
* configure hostname, TLS ([#44](https://github.com/hugginsio/kv2/issues/44)) ([b9420a0](https://github.com/hugginsio/kv2/commit/b9420a0104674257d7221545581b109276f261e4))
* fix version information, add goreleaser ([#45](https://github.com/hugginsio/kv2/issues/45)) ([8dc2c96](https://github.com/hugginsio/kv2/commit/8dc2c968fed48719976f53cafc0c9ecf99ab0ab7))
* invert TLS setting ([#47](https://github.com/hugginsio/kv2/issues/47)) ([14e5fe5](https://github.com/hugginsio/kv2/commit/14e5fe5ef0b73ced76e1179cd520141e3323b223))

## [0.10.0](https://github.com/hugginsio/kv2/compare/v0.8.0...v0.10.0) (2025-04-06)


### ⚠ BREAKING CHANGES

* re-implement server with Connect RPC ([#35](https://github.com/hugginsio/kv2/issues/35))

### Features

* add kv2 CLI ([#36](https://github.com/hugginsio/kv2/issues/36)) ([e1bf8cf](https://github.com/hugginsio/kv2/commit/e1bf8cfec04458dfce2c5262e07612922dba5ea2))
* create FUNDING.yaml ([#33](https://github.com/hugginsio/kv2/issues/33)) ([ecac217](https://github.com/hugginsio/kv2/commit/ecac21702bd5dded43c3cab7cda91cdfdaeff428))
* re-implement server with Connect RPC ([#35](https://github.com/hugginsio/kv2/issues/35)) ([206c96f](https://github.com/hugginsio/kv2/commit/206c96fc84e27958e93e2aeba9de89e5718f3dc1))

## [0.8.0](https://github.com/hugginsio/kv2/compare/v0.7.0...v0.8.0) (2025-03-29)


### Features

* reduce image size ([#32](https://github.com/hugginsio/kv2/issues/32)) ([f06affa](https://github.com/hugginsio/kv2/commit/f06affa48ff4afddc6a761f1d63cae59ab79e001))
* update log behavior ([#30](https://github.com/hugginsio/kv2/issues/30)) ([b8bcc94](https://github.com/hugginsio/kv2/commit/b8bcc94467cfccc2d9643b8d782c732c80a79016))

## [0.7.0](https://github.com/hugginsio/kv2/compare/v0.6.0...v0.7.0) (2025-03-27)


### Features

* add external KMS support and overhaul docs ([#25](https://github.com/hugginsio/kv2/issues/25)) ([a8b6a6c](https://github.com/hugginsio/kv2/commit/a8b6a6ccaf4247963ff5678a29166759e69df116))
* cloud storage support ([#27](https://github.com/hugginsio/kv2/issues/27)) ([fe5ef6c](https://github.com/hugginsio/kv2/commit/fe5ef6cd80ca38b841251fdfe4587f228982d6b1)), closes [#3](https://github.com/hugginsio/kv2/issues/3)
* update docs & API responses ([#28](https://github.com/hugginsio/kv2/issues/28)) ([212a2b7](https://github.com/hugginsio/kv2/commit/212a2b70ccae506c85d98cd079a4e0709c43e95e))

## [0.6.0](https://github.com/hugginsio/kv2/compare/v0.5.1...v0.6.0) (2025-03-18)


### Features

* remove CGO, use noop crypto in dev mode ([#22](https://github.com/hugginsio/kv2/issues/22)) ([a2d934b](https://github.com/hugginsio/kv2/commit/a2d934b1ffe6279c0d226ce8bea382e7acb7ea38))
* update docs, cleanup server ([#20](https://github.com/hugginsio/kv2/issues/20)) ([a161888](https://github.com/hugginsio/kv2/commit/a16188896091a8680e7c78eff4d0e6a15b53f522))

## [0.5.1](https://github.com/hugginsio/kv2/compare/v0.5.0...v0.5.1) (2025-03-18)


### Bug Fixes

* **actions:** update password reference ([#18](https://github.com/hugginsio/kv2/issues/18)) ([3db0965](https://github.com/hugginsio/kv2/commit/3db0965d9a4556fe3c19ebe395539cd991766b15)), closes [#2](https://github.com/hugginsio/kv2/issues/2)

## [0.5.0](https://github.com/hugginsio/kv2/compare/v0.4.0...v0.5.0) (2025-03-17)


### Bug Fixes

* update the release jobs, again ([#16](https://github.com/hugginsio/kv2/issues/16)) ([90b4ed0](https://github.com/hugginsio/kv2/commit/90b4ed03cc3ef73c7bd71f207a3b3f8c0c4249c3))

## [0.4.0](https://github.com/hugginsio/kv2/compare/v0.3.0...v0.4.0) (2025-03-17)


### Bug Fixes

* error in forge return ([#14](https://github.com/hugginsio/kv2/issues/14)) ([1bed30b](https://github.com/hugginsio/kv2/commit/1bed30b04eec0eb85cd4f29c6f202de870c2029a))

## [0.3.0](https://github.com/hugginsio/kv2/compare/v0.2.0...v0.3.0) (2025-03-17)


### Bug Fixes

* release with tags ([#12](https://github.com/hugginsio/kv2/issues/12)) ([6db0546](https://github.com/hugginsio/kv2/commit/6db05468112a93a2f639ebdb4e4a9ab29657a8e1))

## [0.2.0](https://github.com/hugginsio/kv2/compare/v0.1.0...v0.2.0) (2025-03-17)


### Bug Fixes

* update release job ([#10](https://github.com/hugginsio/kv2/issues/10)) ([d68406c](https://github.com/hugginsio/kv2/commit/d68406c4b4b3db1ffeb38270e43aa065290f010a))

## 0.1.0 (2025-03-17)


### Features

* add cd workflow ([#6](https://github.com/hugginsio/kv2/issues/6)) ([ba1cca7](https://github.com/hugginsio/kv2/commit/ba1cca7ec98ad506fad7bc575e70cd459f0d4a60)), closes [#2](https://github.com/hugginsio/kv2/issues/2)
* build image on release ([#9](https://github.com/hugginsio/kv2/issues/9)) ([7422cb9](https://github.com/hugginsio/kv2/commit/7422cb9ba6ed7dca2f13632441bc28802c983a85))
* daggerize PR checks ([#5](https://github.com/hugginsio/kv2/issues/5)) ([b2653d5](https://github.com/hugginsio/kv2/commit/b2653d5c6d3b70b127526fe7a6016450cb1e3809)), closes [#2](https://github.com/hugginsio/kv2/issues/2)
* import from backup ([#1](https://github.com/hugginsio/kv2/issues/1)) ([b55bd20](https://github.com/hugginsio/kv2/commit/b55bd208c2e4e229d5622bc5aaef33bcb12f23fd))
* recreate repository ([f8f65fe](https://github.com/hugginsio/kv2/commit/f8f65fef7bd0fc57d4ce1297f319ff48ab67c83c))


### Bug Fixes

* update release-please token reference ([#7](https://github.com/hugginsio/kv2/issues/7)) ([7b9e47b](https://github.com/hugginsio/kv2/commit/7b9e47b83edb0970c819479dcd2222549a0c89c6)), closes [#2](https://github.com/hugginsio/kv2/issues/2)
