#!/bin/sh
APP_NAME="key_generator"
go install .
go build -o $APP_NAME

if [[ $? != 0 ]]
then
    echo "Build failed"
    exit
fi

foo=${INPUT:?Input file directory is mandatory}
foo=${OUTPUT:?Output file directory is mandatory}

echo "Running..."
./$APP_NAME -input=$INPUT -output=$OUTPUT