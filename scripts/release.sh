#!/bin/bash

set -e

export CURRENT_TAG=$(svu)
export NEXT_TAG=$(svu next)

read -p "Creating new release for $NEXT_TAG (Last version: $CURRENT_TAG). Do you want to continue? [Y/n] " prompt

if [[ $prompt == "y" || $prompt == "Y" || $prompt == "yes" || $prompt == "Yes" ]]; then
	python3 scripts/prepare_changelog.py
	git add CHANGELOG.md
	git commit -m "chore: bump version to $NEXT_TAG for release" || true && git push
	echo "Creating new git tag $NEXT_TAG"
	git tag "$NEXT_TAG" -m "$NEXT_TAG"
	git push --tags
	goreleaser release
else
	echo "Cancelled"
	exit 1
fi
