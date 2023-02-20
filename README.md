# pig-latin

A small webserver that translates a word or a sentence to Pig Latin,
then saves it to a sqlite database.

```
  $ go run main.go
```
Launch http://localhost:3000

`POST` : /
```
{
    "data": "Hello world"
}
```
output
```
 {
    "ID": 1,
    "CreatedAt": "2023-02-19T13:47:50.3077495+08:00",
    "UpdatedAt": "2023-02-19T13:47:50.3077495+08:00",
    "DeletedAt": null,
    "Input": "Hello world",
    "Output": "ellohay orldway"
 }
```

You can also query all of the saved translated word or sentence. Using

`GET` : /
```
[
    {
        "ID": 1,
        "CreatedAt": "2023-02-19T13:47:50.3077495+08:00",
        "UpdatedAt": "2023-02-19T13:47:50.3077495+08:00",
        "DeletedAt": null,
        "Input": "Hello world",
        "Output": "ellohay orldway"
    },
]
