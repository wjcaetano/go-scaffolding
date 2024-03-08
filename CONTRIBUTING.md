# Contributing to go-scaffolding

## How to contribute

We welcome contributions from the community and are pleased to have them. Please follow this guide when logging issues or making code changes.

## Logging Issues

All issues should be created using the [new issue form](https://github.com/wjcaetano/go-scaffolding/issues/new). Clearly describe the issue including steps to reproduce if there are any. Also, make sure to indicate the earliest version that has the issue being reported.

## Patching Code

Code changes are welcome and should follow the guidelines set forth below.

### Forking

When making changes, please first fork the repository and make your changes on a local branch in your fork. Once your changes are complete, open a pull request in the main repository.

### Commits

Make sure that your commits are clear and concise. Each commit should have a clear message of what was changed as well as a reference to the issue it resolves (if any).

Source code, documentation, commit and PR descriptions (except for review comment), must use english language.
All commits **MUST** follow the **seven rules of a great Git commit message** by [Chris Beams](https://cbea.ms/git-commit/)
(we highly suggest that you read the full article):
* Separate subject from body with a blank line.
* Limit the subject line to 50 characters.
* Capitalize the subject line.
* Do not end the subject line with a period.
* Use the imperative mood in the subject line.
* Wrap the body at 72 characters.
* Use the body to explain what and why vs. how.

Commits such as "fix tests", "now it's working" and many other regretable messages you may come up with **WON'T be accepted**.
Remember that if you make something wrong in a given commit you must amend it, not create a new to fix the previous one.
Please avoid taking much time to deliver code, and always rebase your code to avoid reverse merge commits.
We HATE reverse merge commits, they make git's history tree a mess.

### Testing

Any changes made should pass all existing tests, and any new features should be accompanied by new tests that cover the new functionality.

### Code Review

Once a pull request is made, a maintainer will review the code and ensure that it adheres to the standards set forth in the coding guidelines. If any issues are found, they will be reported back and need to be addressed before the changes can be merged.

## Coding Guidelines

All contributors *MUST* follow the [Coding Guidelines](/CODING_GUIDELINES.md). Contributions that do not follow these guidelines will not be taken into account until the necessary modifications are made.

## Pull Request

We use the "fork and pull" model [described here](https://help.github.com/articles/about-collaborative-development-models/), where contributors push changes to their personal fork and create pull requests to bring those changes into the source repository.

Your basic steps to get going:

* Fork this repository and create a branch from master for the issue you are working on.
* Commit as you go following our git conventions.
* Include tests that cover all non-trivial code. The existing tests should provide a template on how to test correctly.
* Make sure every test passes.
* All code changes are expected to comply with the formatting suggested by `gofmt` & the current in use `golint` configuration.
* Push your commits to GitHub and create a pull request against the corresponding component develop branch.

If taking too much time to deliver code, **always** [rebase](https://git-scm.com/docs/git-rebase) towards `master` before asking for a review, and avoid reverse merge commits. We **HATE** reverse merge commits (they make git history tree a mess) and will reject contributions that have them.
