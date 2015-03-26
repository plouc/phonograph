# phonograph

## Installation

### Requirements

* neo4j
* go
* nodejs/npm
* gulp

#### Load sample data

```sh
cd api
go run data.go
```

#### Build assets

```sh
cd app
gulp
```

## API

```
GET /artists
GET /artists/12
GET /artists/12/similars
GET /artists/12/masters

GET /skills
GET /skills/13
GET /skills/13/artists

GET /masters
GET /masters/14
```