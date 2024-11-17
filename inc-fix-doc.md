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


## solutions
### Solution 1: Modify Batch Increment

- The current implementation is designed for raw numbers and does not account for sequence options.
- It may be possible to extend the Batch Increment command to include sequence options, but this would require informing the sequence cache about the available number of values.
- However, modifying the Batch Increment in this way is likely not ideal, as it introduces a one-way change that could complicate future modifications or reversals.

pros:
- No new apis

cons:
- blots existing apis
- backward combatabile

### Solution 2: Create a New Batch Command

- This approach would support sequence options directly.
- The command would return the number of increments successfully performed.

pros:
- backward combatabile

cons:
- Creating a new api requires proper documentation
- Not sure if understanding is correct enough

## Questions:

### What is the purpose of Batch Commands?
- Are all low-level requests sent through batch commands?
> Batch provides for the parallel execution of a number of database
operations. Operations are added to the Batch and then the Batch is executed
via either DB.Run, Txn.Run or Txn.Commit.

### what is the purpose of sequences using the batch increment command?
- a