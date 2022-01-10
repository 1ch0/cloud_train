#!/bin/bash
#
#***********************************************************************
#         Author:            1ch0
#         Date:              2022-01-07
#         FileName:          docker_run_test.sh
#         Description:       1ch0 script
#         Blog:              https://1ch0.github.io/
#         Copyright (C):     2022 All rights reserved
#***********************************************************************
image=$(grep -m 1 1ch0 build_image.sh | awk '{print $1}')
# 删除字符中的首尾单引号
eval image=$image

echo -e "running docker image: $image\n"

docker run -it --rm -p8080:8080 "$image" bash
