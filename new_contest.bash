#!/bin/bash
set -eu

if [ $# -ne 2 ]; then
    echo "specify the contest_name and the nubmer"
    exit 1
fi

contest_name=$1
contest_number=$2

cd `dirname $0`
cwd=`pwd`

contest_dir="${cwd}/${contest_name}"
if [ ! -e ${contest_dir} ]; then
    mkdir ${contest_dir}
fi

number_dir="${cwd}/${contest_name}/${contest_number}"
if [ -e ${number_dir} ]; then
    echo "${number_dir} directory already exists."
    exit 1
fi
mkdir ${number_dir}

tasks=("a" "b" "c" "d" "e" "f")
for task in ${tasks[@]}
do
    mkdir "${number_dir}/${task}"
    cp ${cwd}/template.go ${number_dir}/${task}/main.go
done
