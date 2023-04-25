Implement a simple text editor. The editor initially contains an empty string, . Perform  operations of the following  types:

1) append - Append string to the end
2) delete - Delete the last x characters
3) print - Print the  character of x
4) undo - Undo the last (not previously undone) operation, reverting  to the state it was in prior to that operation.
Example

operation
index   S       ops[index]  explanation
-----   ------  ----------  -----------
0       abcde   1 fg        append fg
1       abcdefg 3 6         print the 6th letter - f
2       abcdefg 2 5         delete the last 5 letters
3       ab      4           undo the last operation, index 2
4       abcdefg 3 7         print the 7th characgter - g
5       abcdefg 4           undo the last operation, index 0
6       abcde   3 4         print the 4th character - d