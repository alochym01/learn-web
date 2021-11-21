curl -XPOST -H "Content-Type: application/json" -d @createAlbum1.json http://127.0.0.1:8080/albums
curl -XPOST -H "Content-Type: application/json" -d @createAlbum2.json http://127.0.0.1:8080/albums
curl -XPOST -H "Content-Type: application/json" -d @createAlbum3.json http://127.0.0.1:8080/albums
{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
