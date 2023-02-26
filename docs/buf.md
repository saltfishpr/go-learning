[Specifying an input](https://docs.buf.build/reference/inputs)

As of now, there are seven other options, all of which are format specific:

- The `branch` option specifies the branch to clone for git inputs.
- The `tag` option specifies the tag to clone for git inputs.
- The `ref` option specifies an explicit git reference for git inputs. Any ref that is a valid input to git checkout is accepted.
- The `depth` option optionally specifies how deep of a clone to perform. This defaults to 50 if ref is set, and 1 otherwise.
- The `recurse_submodules` option says to clone submodules recursively for git inputs.
- The `strip_components` option specifies the number of directories to strip for tar or zip inputs.
- The `subdir` option specifies a subdirectory to use within a git, tar, or zip input.

If ref is specified, branch can be further specified to clone a specific branch before checking out the ref.
