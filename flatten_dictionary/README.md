
Write a function to flatten a nested dictionary. Namespace the keys with a period.


{
    "key": 3,
    "foo": {
        "a": 5,
        "bar": {
            "baz": 8
        }
    }
}

Becomes: 

{
    "key": 3,
    "foo.a": 5,
    "foo.bar.baz": 8
}