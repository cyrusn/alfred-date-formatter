#!/usr/local/bin/python3
import json
import argparse
import re
from datetime import date
from math import floor


CHINESE_WEEKDAY = ["一", "二", "三", "四", "五", "六", "日"]
CHINESE_YEAR_DIGIT = ["零", "一", "二", "三", "四", "五", "六", "七", "八", "九"]
CHINESE_MONTH = ["", "一", "二", "三", "四", "五",
                 "六", "七", "八", "九", "十", "十一", "十二"]
CHINESE_DAY_TENS_DIGIT = ["", "十", "二十", "三十"]
CHINESE_DAY_ONES_DIGIT = ["", "一", "二", "三", "四", "五", "六", "七", "八", "九"]


INPUT_DATE_REG_EXP = [
    r"^(?P<year>[0-9]+)年(?P<month>[0-9]+)月(?P<day>[0-9]+)日$",
    r"^(?P<month>[0-9]+)月(?P<day>[0-9]+)日$",
    r"^(?P<day>[0-9]+)\/(?P<month>[0-9]+)\/(?P<year>[0-9]+)$",
    r"^(?P<day>[0-9]+)\/(?P<month>[0-9]+)$",
    r"^(?P<year>[0-9]+)[\.|\-](?P<month>[0-9]+)[\.|\-](?P<day>[0-9]+)$",
    r"^(?P<month>[0-9]+)[\.|\-](?P<day>[0-9]+)$",
]

OUTPUT_DATE_FORMATS = [
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


def format2ChineseYear(year):
    year_one = year % 10
    year_ten = floor(year / 10) % 10
    year_hundred = floor(year / 100) % 10
    year_thousand = floor(year / 1000)
    return (
        CHINESE_YEAR_DIGIT[year_thousand] +
        CHINESE_YEAR_DIGIT[year_hundred] +
        CHINESE_YEAR_DIGIT[year_ten] +
        CHINESE_YEAR_DIGIT[year_one]
    )


def format2ChineseDay(day):
    day_ten = floor(day / 10)
    day_one = day % 10
    return CHINESE_DAY_TENS_DIGIT[day_ten] + CHINESE_DAY_ONES_DIGIT[day_one]


def getChineseDateFormat(myDate):
    return {
        "year": format2ChineseYear(myDate.year),
        "month": CHINESE_MONTH[myDate.month],
        "day": format2ChineseDay(myDate.day),
        "weekday": CHINESE_WEEKDAY[myDate.weekday()],
    }


def parseDateString(date_string):
    today = date.today()
    myDate = today
    for reg in INPUT_DATE_REG_EXP:
        searchResult = re.search(reg, date_string)
        if searchResult:
            parsedDateDict = {k: int(v)
                              for k, v in searchResult.groupdict().items()}
            if "year" not in parsedDateDict:
                parsedDateDict["year"] = today.year

            myDate = date(**parsedDateDict)
            if myDate < today:
                myDate = myDate.replace(year=myDate.year + 1)

            break
    return myDate


def formatDate(fmt, date_string):
    d = parseDateString(date_string)
    c = getChineseDateFormat(d)
    return fmt.format(d, c=c)


def convertAlfredJSONString(date_string):
    """
    return JSON sting for script filter of alfred workflow.

    https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
    """
    myJSON = {}
    myJSON["items"] = []

    for f in OUTPUT_DATE_FORMATS:
        arg = formatDate(f, date_string)

        myJSON["items"].append({
            "title": arg,
            "arg": arg
        })
    return json.dumps(myJSON, ensure_ascii=False)


if __name__ == "__main__":

    helptext = "Chinese Date formatter:\n\nTYPE:\n"
    for i, t in enumerate(OUTPUT_DATE_FORMATS):
        helptext += "({:d}): {}\n".format(i, t)

    ap = argparse.ArgumentParser(
        formatter_class=argparse.RawDescriptionHelpFormatter,
        description=helptext
    )
    ap.add_argument(
        "input",
        help="Input date string for parser", nargs="?"
    )

    ap.add_argument(
        "-j",
        "--json",
        help="Print to JSON format for script filter of alfred workflow",
        action="store_true"
    )

    ap.add_argument(
        "-t",
        "--type",
        help="Print the formatted date with given type number",
        default=0, type=int
    )
    args = ap.parse_args()

    today = format(date.today())
    date_string = args.input if args.input is not None else today
    if args.json:
        alfred_json = convertAlfredJSONString(date_string)
        print(alfred_json)
        exit()

    if args.type:
        try:
            fmt = OUTPUT_DATE_FORMATS[args.type]
            print(formatDate(fmt, date_string))
        except IndexError:
            print('Unexpected integer for type')
        exit()

    print(parseDateString(date_string))
