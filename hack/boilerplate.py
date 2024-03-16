#!/usr/bin/env python

# Copyright 2023 The gin-apiserver Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os

# 定义要添加头部的文件扩展类型
file_extensions = ['.go']

# 读取版权信息的函数
def read_boilerplate(boilerplate_path):
    with open(boilerplate_path, 'r', encoding='utf-8') as file:
        return file.read()

# 给文件添加头部的函数
def add_boilerplate_to_file(file_path, boilerplate):
    with open(file_path, 'r+', encoding='utf-8') as file:
        contents = file.read()
        if contents.startswith(boilerplate):
            return  # 如果文件已包含头部，则不执行任何操作
        file.seek(0, 0)
        file.write(boilerplate + '\n\n' + contents)

# 遍历目录给所有文件添加头部的函数
def add_boilerplate_to_directory(directory, boilerplate):
    for subdir, _, files in os.walk(directory):
        for file_name in files:
            if any(file_name.endswith(ext) for ext in file_extensions):
                file_path = os.path.join(subdir, file_name)
                add_boilerplate_to_file(file_path, boilerplate)

# 主函数
def main():
    # 假设此脚本与boilerplate.go.txt在同一目录下
    boilerplate_path = 'hack/boilerplate.go.txt'
    boilerplate = read_boilerplate(boilerplate_path)

    # 定义需要添加头部的目录
    directories = ['../cmd', '../pkg']
    for directory in directories:
        full_directory_path = os.path.join(os.path.dirname(__file__), directory)
        add_boilerplate_to_directory(full_directory_path, boilerplate)

if __name__ == '__main__':
    main()
