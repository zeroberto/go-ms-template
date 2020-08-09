#! /bin/bash

# terms
oldRepositoryOwnerName='zeroberto'
oldAppName='go-ms-template'

files=`find -type f ! -path "./.git*"`

echo 'Welcome to the initial configuration of the application based on the go-ms-template. Please enter the following parameters:'
echo -e '\n'

# app name
echo -n '1. Enter the application name and press [ENTER]: '
read newAppName

# repository owner name
echo -n '2. Enter the repository owner name and press [ENTER]: '
read newRepositoryOwnerName

echo -e '\n'
echo 'Do you confirm the new configuration?'
echo "1. Application name: $newAppName"
echo "2. Repository owner name: $newAppName"
echo -e '\n'

echo -n 'Press [y/n]: '
read confirmation

if [[ $confirmation == 'y' ]]; then
  for file in $files
  do
    sed -i "s/$oldRepositoryOwnerName/$newRepositoryOwnerName/g" $file
    sed -i "s/$oldAppName/$newAppName/g" $file
  done
  echo 'Setup completed. If you prefer, the file setup.sh can be removed.'
fi
