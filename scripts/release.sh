#!/bin/bash

set -e

export TAG=$(svu next)

read -p "Creating new release for $TAG. Do you want to continue? [Y/n] " prompt

if [[ $prompt == "y" || $prompt == "Y" || $prompt == "yes" || $prompt == "Yes" ]]; then
	python3 scripts/prepare_changelog.py
	git add CHANGELOG.md
	git commit -m "chore: bump version to $TAG for release" || true && git push
	echo "Creating new git tag $TAG"
	git tag "$TAG" -m "$TAG"
	git push --tags
else
	echo "Cancelled"
	exit 1
fi
