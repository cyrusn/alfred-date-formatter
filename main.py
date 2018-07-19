#!/usr/local/bin/python3

import sys
import json
from chineseDateCharacter import getChineseDateCharacter
from parser import Parser


query = sys.argv[1]
formats = [

    "{:%Y-%m-%d (%a)}",
    "{:%-m月%-d日}",
    "{:%-m月%-d日}（{c[weekday]}）",
    "{:%Y年%-m月%-d日}",
    "{:%Y年%-m月%-d日}（星期{c[weekday]}）",
    "{:%d/%-m/%y (%a)}",
    "{:%a, %-d %b %y}",
    "{:%A, %-d %B %Y}",
    "{c[year]}年{c[month]}月{c[day]}日",
    "{c[year]}年{c[month]}月{c[day]}日（星期{c[weekday]}）",
]

d = Parser(query)

alfredJSON = {}
alfredJSON['items'] = []

for format in formats:
    c = getChineseDateCharacter(d.year, d.month, d.day)
    arg = format.format(d.date, c=c)

    alfredJSON['items'].append({
        'title': arg,
        'arg': arg
    })

print(json.dumps(alfredJSON, ensure_ascii=False))
