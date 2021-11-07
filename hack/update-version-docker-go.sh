#!/bin/bash

# Find the version of the component in the docker folder and update it.
# It requiers an argument which is the component.go file of one of the directories in the docker folder
# It uses SetDockerBuildInfo function in go to find the version

SOURCE_FILE=$1

echo "Finding version in $SOURCE_FILE"
ORIGINAL_VERSION=$(cat $SOURCE_FILE | grep SetDockerBuildInfo | awk -F"," '{print substr($2, 2)}' | awk '{print substr($0, 2, length($0) - 2)}')

echo "Original version $ORIGINAL_VERSION"

# Split version into numbers
IFS='.' read -ra ver <<< "$ORIGINAL_VERSION"
#Take the minor and update
minor=${ver[1]}
minor=$((minor+1))

NEW_VERSION="${ver[0]}.$minor.${ver[2]}"

echo "New version $NEW_VERSION"
sed -i "s/$ORIGINAL_VERSION/$NEW_VERSION/" $SOURCE_FILE
