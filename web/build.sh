#!/usr/bin/env bash
npm run build
rm -rf ../website
mv dist ../website