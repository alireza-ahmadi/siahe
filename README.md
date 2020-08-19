# Siahe
Siahe provides a simple API for indexing terms that are related to a specific ID
and retrieving those IDs by doing efficient prefix search on their indexed related
terms. Current implementation of Siahe uses Radix Tree under the hood.

## FAQ

### What is the use case of Siahe?
I believe that searching over a set of terms which has low-to-zero change frequency
is the best use case for this library. In fact, I've implemented this library to
use it in a service which is responsible for providing instant(ish) search functionality
over a few thousand terms.

### Can Siahe be thread safe?
Unfortunately no. In my use case, data wouldn't change after the initial insertion
so I've gone with a simpler single writer thread approach; However, you can use
multiple threads for searching among indexed terms.

### What does "Siahe" mean?
Siahe means list in Persian.