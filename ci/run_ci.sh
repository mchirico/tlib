#!/bin/bash

set -eu

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export fly_target=${fly_target:-mce}
echo "Concourse API target ${fly_target}"
echo "Tutorial $(basename $DIR)"

pushd $DIR
  fly -t ${fly_target} set-pipeline -p tlib-pipeline -c build-golang-pipeline.yml -n
  fly -t ${fly_target} unpause-pipeline -p tlib-pipeline
#  fly -t ${fly_target} trigger-job -w -j tutorial-pipeline/job-hello-world
popd

echo -e "\n\n                  Common commands:"
echo -e "**************************************\n\n"
echo -e "\n"
echo -e "                           fly -t mce watch --job tlib-pipeline/unit"
echo -e "                           fly -t mce builds|grep 'tlib-pipeline'"
echo -e "                           fly -t mce destroy-pipeline -p tlib-pipeline -n"
echo -e "                           fly -t mce workers -d "
echo -e "                            "
echo -e "                           To login to a running container: "
echo -e "                           fly -t mce intercept --job tlib-pipeline/unit "
echo -e "\n"
echo -e "\n"

