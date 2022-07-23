## Builder Summary

- A builder is a separate component used for building an object step-by-step.
- A builder can be fluent or not, it depends on whether the receiver function receiver is returned or not. If it is returned, we have a fluent builder, which allows chain calls.
- More complex objects can have more than one builders working on it. Each builder can take care of a specific subobject of the main object.
- You can use builder parameters if you want to isolate your data structures from your client, so they only ever interact with the builders.
- A builder can be functional, having its build steps delayed to only be executed when the .Build() call is method, rather than on the go.