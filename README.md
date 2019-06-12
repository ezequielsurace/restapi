## Endpoints

### Get All Books
``` bash
GET api/books
```
### Get Single Book
``` bash
GET api/books/{id}
```

### Delete Book
``` bash
DELETE api/books/{id}
```

### Create Book
``` bash
POST api/books

# Request sample
# {
#   "isbn":" 9780316234498",
#   "title":"The 100",
#   "author":{"firstname":"Kass",  "lastname":"Morgan"}
# }
```

### Update Book
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"9780316234498",
#   "title":"The 100",
#   "author":{"firstname":"Kass",  "lastname":"Morgan"}
# }

```
