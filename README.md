# Introduction

These instructions are intended to make it easier for developers to create new
Cluster API provider repositories. For more information on the Cluster API project
please see the base repository [here](https://github.com/kubernetes-sigs/cluster-api).

There is no intent to maintain these instructions. Better automation or more
sophisticated code sharing should be preferred over this duct tape and bubble gum.

¯\_(ツ)_/¯

# Quickstart

- The name of provider repos [should](
https://github.com/kubernetes-sigs/cluster-api/issues/383) be of the form
`cluster-api-provider-$(cloud)`. For example, `cluster-api-provider-openshift`
is the name of the repo which implements the Cluster API provider for OpenShift.

- [Create](https://help.github.com/articles/creating-a-new-repository/) a new
empty GitHub repo under your org using the GitHub GUI, for example
https://github.com/samsung-cnct/cluster-api-provider-ssh.

- [Duplicate](https://help.github.com/articles/duplicating-a-repository/)
this repo (https://github.com/davidewatson/cluster-api-provider-skeleton) and
push it to the `cluster-api-provider-ssh` repo you created in the previous
step. Note the arguments to clone and push.

```
git clone --bare https://github.com/davidewatson/cluster-api-provider-skeleton.git
cd cluster-api-provider-skeleton.git
git push --mirror https://github.com/davidewatson/cluster-api-provider-ssh.git
cd ..
rm -rf cluster-api-provider-skeleton.git
```

- Clone the new repository.
  For Go dependencies to be built correctly with dep, place this repository in your $GOPATH as follows:

```
mkdir -p $GOPATH/src/sigs.k8s.io/
cd $GOPATH/src/sigs.k8s.io/
git clone https://github.com/davidewatson/cluster-api-provider-ssh.git
cd cluster-api-provider-ssh
```

- Customize the new repository. A simple search and replace may suffice for
some name changes, e.g. on OS X, something like this might work:

```
find ./* -type f  -path ./vendor -prune -o -path ./.git -prune -o -exec sed -i '' -e 's/skeleton/ssh/g' {} \;
find ./* -type f  -path ./vendor -prune -o -path ./.git -prune -o -exec sed -i '' -e 's/Skeleton/SSH/g' {} \;
```

The following directory must be renamed/moved:

```
git mv pkg/cloud/skeleton/ pkg/cloud/ssh/
```

For other changes, like the README.md, OWNERS_ALIASES, etc., you'll have to
think more.

- Get all the dependencies in your vendor directory
```
make depend
```
