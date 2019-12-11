#!/usr/bin/env bash

git add .

echo -n "enter git commit message:"
read name
git commit -m "$name"
git push origin master