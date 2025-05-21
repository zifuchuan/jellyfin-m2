# _*_ coding:utf-8 _*_
# !/usr/bin/env python

import html, io, json, os, pickle, random, re, sqlite3, string, sys, subprocess, time, traceback, urllib.parse, \
    urllib.request, xml.etree.ElementTree
import lxml.etree


def func1():
    # 检测出
    fd = open('test.txt')
    results = fd.read()
    return results


def func2():
    # 未检测
    fd = open('test.txt')
    results = fd.read()

    fd.close()
    return results
