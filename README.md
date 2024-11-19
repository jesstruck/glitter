# Glitter

And overlay to git to work on multiple repos for a single organisations

---

## Overview
Glitter is a CLI-tool written in [Go(Golang)](https://go.dev/), it's intended users are DevOps people that might not necessarily work on all the repos in an org, needs quick access to all the code, in a clean state.

## Usage

### Custom SSH agent

If you store your git ssh keys in a custom SSH agent, like [1Password's agent](https://developer.1password.com/docs/ssh/get-started/#step-4-configure-your-ssh-or-git-client)

You can still use this tool by setting, in the same shell as you intend to run glitter

         export SSH_AUTH_SOCK=<PATH_TO_YOUR_SSH_SOCKET>

## Installation

        git clone git@github.com:jesstruck/glitter.git
        cd glitter
        make install
