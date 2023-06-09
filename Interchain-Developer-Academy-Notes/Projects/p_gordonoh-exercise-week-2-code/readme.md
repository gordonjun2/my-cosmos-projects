# Exercise

This is an exercise. It is graded, but your grade carries no consequences.

## Preparation

We created a Docker container with the included `Dockerfile`:

```sh
$ docker build . -t exercise-w2
```

Create this Docker image right now, before you start working on the exercise. It takes time and it makes a copy of the original files.

## What we set up

We created the rest of this repository with Ignite CLI v0.22.1 and the following command:

```sh
$ docker run --rm -it -v $(pwd):/exercise -w /exercise exercise-w2 ignite scaffold chain github.com/b9lab/other-world
```

The idea behind it is that this is the blockchain backing a future metaverse. The blockchain will account for anything of value in it.

We have added:

* One file that duplicates what accessing the storage looks like: `x/otherworld/keeper/world_params_duplicate.go`.
* And 2 test files:

  * `x/otherworld/types/world_params_student_test.go`
  * `x/otherworld/keeper/world_params_student_test.go`

All 3 files cannot compile before you have done as per below.

## What you need to do

You need to run a Ignite command to create a new Protobuf type and its associated keeper functions that follows this description:

1. The name is `WorldParams`.
2. It has 3 params exactly, no more, no less:
   1. A `string` for the world's `name`.
   2. A `uint` for the world's `gravity`.
   3. A `uint` for the world's `landSupply`.
3. It is unique in the blockchain store. At what key it is stored does not matter, but you should keep the naming conventions of keeper functions chosen by Ignite.
4. It should not have any associated messages.

## Definition of success

Make these tests pass:

* [`x/otherworld/types/world_params_student_test.go`](x/otherworld/types/world_params_student_test.go): It needs to pass with:

    ```sh
    $ docker run --rm -it -v $(pwd):/exercise -w /exercise exercise-w2 go test github.com/b9lab/other-world/x/otherworld/types
    ```

    Or:

    ```sh
    $ go test github.com/b9lab/other-world/x/otherworld/types
    ```

* [`x/otherworld/keeper/world_params_student_test.go`](x/otherworld/keeper/world_params_student_test.go): It needs to pass with:

    ```sh
    $ docker run --rm -it -v $(pwd):/exercise -w /exercise exercise-w2 go test github.com/b9lab/other-world/x/otherworld/keeper
    ```

    Or:

    ```sh
    $ go test github.com/b9lab/other-world/x/otherworld/keeper
    ```

To run them both at the same time, run:

```sh
$ docker run --rm -it -v $(pwd):/exercise -w /exercise exercise-w2 /exercise/score.sh
```

Or:

```sh
$ ./score.sh
```

It returns you a score that reads like:

```txt
FS_SCORE:100%
```

### Read-only files

In [`fileconfig.yml`](./fileconfig.yml), we have defined a list of _read-only_ files. They are not technically read-only, but we will overwrite your changes if you modify them.

## CI Environment

There is a CI environment on the Academy Gitlab server. Every time you push to your repository, it will run and score your work.

The script that our CI runs is `score-ci.sh`. You will notice that it overwrites the _read-only_ files with their original ones. So if the CI gives you a score different from the one you have on your machine, you may have made changes to read-only files.

The score that appears on Gitlab's CI is the _correct_ one.