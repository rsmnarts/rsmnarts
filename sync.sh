#! /bin/bash

cp -rf ~/.config/nvim .
cp -rf ~/.zshrc .
git add -A
git commit -m 'sync files'
git push origin master
