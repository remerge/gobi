# Gob Indirect

Gob encoding with indirect (pointer) support

Objects which are referenced by pointers are only writen once. Before the object
is written an ID is written. If the object was already written once, just the id
is stored. This breaks several things (we might fix that in the future).
Interfaces are not supported atm.

Note: this is not compatible with gob at the moment.
