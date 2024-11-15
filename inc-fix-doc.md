summary:
Sequences provide a way to automicatically increment a value. Sequesence are provided options which affect its behavior such as increment, or max options.
Currently there is an issue where the sequence values maybe go beyound the min/max values set in the options.
This happens due to current sequence cache implementation and its usage of the Batch Increment Command.

crs:
- https://github.com/andrew-delph/cockroach/pull/17/files
- https://github.com/andrew-delph/cockroach/pull/18/files

possible documentation: https://github.com/andrew-delph/cockroach/blob/master/docs/RFCS/20150806_gateway_batch.md

desired behavior:
- considering max_value, min_value, increment, cache.
- increase to the largest possible number with range with the bnumber of cache steps
- if increasing the cache reaches the bounds, set the cache to the ramainder of the increment
- if no steps can be made, the increase in an error

sol1: modify batch increase
- it implmented for raw numbers and does not consider sequence options
- it may be possible to extend the batch inrement command to include sequence options but this would require informating the sequence cache the number of values available
- it is likely not a good idea to modify increment batch in this way as it is considered a 1 way change

sol2: create new batch command
- this will provide sequence options
- it will return the number of increases made

sol3: in batch, 