## Factory Summary

- A factory function (a.k.a constructor) is a helper function for making struct instances
- A factory is any entity that can take care of object creation
- Factory functions can be high-order functions, receiving functions as parameters in order to take care of creating object instances
- You can also have structural factories, which are structs with methods dedicated to generating instances of the domain objects
- Another approach is the prototype factory, which doesn't fully initialize objects, but makes it easier to create pre-built instances