# Shred
Go function to overwrite a file, then delete it.

The function first gets the file's size.

If the size is 0, there is no data to overwrite. The file is deleted and the function returns.

Otherwise, the function loops over the file 3 times, generating new blocks of random data and writing them over the file's contents. Finally, the file is deleted from the filesystem, and the function returns.

Test cases:
All test cases expect to have the test files prepared by running "make" in the project directory.
If you do not wnat these files any longer, run "make clean".

Current tests:
- Typical file a user has access to
- An empty file a user has access to
- A file the user owns, but permission bits deny access
- A file the user down't own and doesn't give "other" write access

Possible tests:
- Testing files that someone else already has open
- Testing targets that aren't regular files, such as
    directories
    named pipes
    block and character special files
- Files that are larger than an int64 in length