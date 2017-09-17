#!/bin/bash

while getopts ":u:r:" o; do
	case "${o}" in
		u)
			user=${OPTARG}
			;;
		r)
			repo=${OPTARG}
			;;
	esac
done

rootappfolder="/var/app/deploy/$user/$repo"

mkdir -p $rootappfolder

git clone https://github.com/$user/$repo $rootappfolder
