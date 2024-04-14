# Nothing Search

## A play on Everything

## 0.0.2

- A bit of thumbnail generation. It should probably NOT create all files at once, but rather just generate it on the fly and save, as in the previous solution.
- A bit of database

## 0.0.1

- Basic image support
- Thumbnails

## To be done

- Add a SQLite database to speed up the app
- Always serve thumbnails as jpeg (?)
- Keep some thumbnails cached on the disk, probably the most recent ones. The rest can be just served on the fly. If a 256x256 image weighs under 50KB and the cache limit is 512MB there could be 10k images cached at any given time
- Some sort of pagination
