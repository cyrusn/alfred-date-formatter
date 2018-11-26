#!/usr/local/bin/python3

import sys
import json
import argparse
from chineseDateCharacter import getChineseDateCharacter
from parser import Parser
from datetime import date, datetime


def convert2JSON(datestring):
    d = Parser(datestring)
    formats = [
        "{:%Y-%m-%d}",
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
    alfredJSON = {}
    alfredJSON['items'] = []

    for format in formats:
        c = getChineseDateCharacter(d.year, d.month, d.day)
        arg = format.format(d.date, c=c)

        alfredJSON['items'].append({
            'title': arg,
            'arg': arg
        })
    return json.dumps(alfredJSON, ensure_ascii=False)


if __name__ == "__main__":
    today = format(date.today())
    ap = argparse.ArgumentParser()
    ap.add_argument('-i', '--input',
                    help='Input date string for parser', nargs='?', default=today)
    args = ap.parse_args()
    datestring = args.input if args.input is not None else today
    alfred_json = convert2JSON(datestring)
    print(alfred_json)
