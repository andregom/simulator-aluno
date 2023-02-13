#!/bin/bash

if [! -f ".env" ]; then
    cp .env.example .env
fi

npm install --save-dev

npm run start:dev
