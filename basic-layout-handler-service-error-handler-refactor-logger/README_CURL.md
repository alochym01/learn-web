## Using CURL
1. curl -XGET http://127.0.0.1:8080/albums
1. curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/albums/:id
1. curl -XGET -H "Content-Type: application/json" -d @getAlbum.json http://127.0.0.1:8080/albums/:id
1. curl -XPOST -H "Content-Type: application/json" -d @createAlbum.json http://127.0.0.1:8080/albums
1. curl -XPUT -H "Content-Type: application/json" -d @updateAlbum.json http://127.0.0.1:8080/albums/:id
1. curl -XDELETE -H "Content-Type: application/json" -d @deleteAlbum.json http://127.0.0.1:8080/albums/:id
