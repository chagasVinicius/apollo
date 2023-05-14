## Add Category(V1)

### Endpoint
```json
//POST
//// Endpoint
"/v1/categories"

//body
{
    "name" : string,
    "short_desc": string
}

// response body
{
    "id" : UUID,
    "name": string,
    "short_desc": string,
    "created_at": timestamp,
    "updated_at": timestamp
}
```

### Fluxo
- Check request body
- Add new category to database
- generate event to search category musics
- return category
