## Singleton Summary

- Lazy one-time initialization using sync.Once
- Adhere to DIP: depend on interfaces, not concrete types. This way, you can replace your singletons when testing
- Singleton is not scary :)