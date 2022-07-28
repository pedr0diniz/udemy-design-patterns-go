## Adapter Summary

- Implementing an Adapter is easy
- You determine the API you have and the API you need
- Create a component which aggregates (has a pointer to, ...) the adaptee
- Intermediate representations can pile up: use caching and other optimizations to avoid unnecessary use of resources