#!/bin/bash

echo 'let courseNameMap: [String: String] = ['
echo $(
    curl -s https://www.gran-turismo.com/jp/gtsport/module/community/localize/ \
        | jq . \
        | gsed -rn 's/.*CourseName.(.*)": "(.*)".*/"\1": "\2",/p'
)
echo ']'
