# Nothing Search

## A play on Everything

## 0.0.1

### Added

- Basic image support
- Thumbnails

### Changed

### Fixed

### To be done

- Add a SQLite database to speed up the app
- Always serve thumbnails as jpeg (?)
- Keep some thumbnails cached on the disk, probably the most recent ones. The rest can be just served on the fly. If a 256x256 image weighs under 50KB and the cache limit is 512MB there could be 10k images cached at any given time
- Some sort of pagination
