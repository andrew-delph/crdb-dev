## summary:
Sequences provide a way to automatically increment values. They can be configured with various options that control their behavior, such as the increment step and maximum value. However, there is currently an issue where sequence values may exceed the defined minimum or maximum limits. This occurs due to the current sequence cache implementation and its handling of the Batch Increment Command.

crs:
- https://github.com/andrew-delph/cockroach/pull/17/files
- https://github.com/andrew-delph/cockroach/pull/18/files

possible documentation: https://github.com/andrew-delph/cockroach/blob/master/docs/RFCS/20150806_gateway_batch.md

## desired behavior

1. **Parameters to Consider**:
   - `max_value`
   - `min_value`
   - `increment`
   - `cache`

2. **Increment Process**:
   - Increase the value to the largest possible number within the defined range, using the specified number of cache steps.

3. **Handling Boundaries**:
   - If increasing by the cache size exceeds the bounds (`max_value` or `min_value`), adjust the cache to use the remainder of the increment.

4. **Error Handling**:
   - If no valid steps can be made within the bounds, raise an error.


## Difficulties
**Sequence logic and Increment logic is seperated by an API**

This forces us to create or udpate an Api. If possible, skipping intraction with APIs directly could simply a fix. 
Logic needs to be introduced between operations of Read and Write. Increment currently only returns starting value of the operation. This means we are blindly incrementing the value. After the operation is made, the written value is validated to set the cache value range. The cache will then return accurate results but the written value is corrupted and once sequence options are then extended, the value made unreasonablity leaps.
The use of the cache and sequences is a known trade off in other databases, for example Postgres. It results in missed sequence values when the cache is invalidated.
The difference is that CRBDs cache introduces larger leaps and actually may return a value with is out of range.

pros:
- i

cons:
- b 


**Template**

pros:
- a

cons:
- b 

## Solutions
### Solution 1: Modify Api for Integer Increments

- The current implementation is designed for raw numbers and does not account for sequence options.
- It may be possible to extend the Batch Increment command to include sequence options, but this would require informing the sequence cache about the available number of values.
- However, modifying the Batch Increment in this way is likely not ideal, as it introduces a one-way change that could complicate future modifications or reversals.

pros:
- No new apis

cons:
- Increment has uses with are beyond sequences scope
- Backwards combatability

### Solution 2: Create Api Specifically for Sequences

- This approach would support sequence options directly.
- The command would return the number of increments successfully performed.

pros:
- Backwards combatability

cons:
- Creating a new api requires proper documentation
- Not sure if understanding is correct enough

### Solution 3: Skip direct usage of APIs

Locks/Transaction:

1. get value
2. calculate
3. write

**Why is this not logic used?**
> It slows down logic. Possibly the documention for Postgres has an explanation.

### Solution 4: Batch Requests
In 1 request, include multiple batch rows and operations.

## Questions:

### What is the purpose of Batch Commands?
- Are all low-level requests sent through batch commands?
> Batch provides for the parallel execution of a number of database
operations. Operations are added to the Batch and then the Batch is executed
via either DB.Run, Txn.Run or Txn.Commit.

### what is the purpose of sequences using the batch increment command?
- a

### How else is Increment API used?
- this should be documented