cmake_minimum_required(VERSION 2.8)

project(go-sqlite3-extension-functions)

add_library(go-sqlite3-extension-functions SHARED extension-functions.c)

install(TARGETS go-sqlite3-extension-functions DESTINATION lib)
