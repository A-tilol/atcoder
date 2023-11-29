#!/bin/bash
set -eu

if [ $# -lt 3 ]; then
    echo "specify the contest name, the contest nubmer, the number of tasks and lang(option)."
    exit 1
fi

contest_name=$1
contest_number=$2
task_num=$3

lang="go"
if [ $# -eq 4 ]; then
    lang=$4
fi

cd $(dirname $0)
cwd=$(pwd)

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

tasks=("a" "b" "c" "d" "e" "f" "g" "h" "i" "j" "k" "l" "m" "n" "o" "p" "q" "r" "s" "t" "u" "v" "w" "x" "y" "z")
for i in ${!tasks[@]}; do
    if [ $i -ge $task_num ]; then
        break
    fi
    mkdir "${number_dir}/${tasks[$i]}"
    cp ${cwd}/template.${lang} ${number_dir}/${tasks[$i]}/main.${lang}
done
