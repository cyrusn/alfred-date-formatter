from datetime import datetime
from math import floor


def getChineseDateCharacter(year, month, day):
    cWeekDays = ["一", "二", "三", "四", "五", "六", "日"]
    cYears = ["零", "一", "二", "三", "四", "五", "六", "七", "八", "九"]
    cMonths = ["", "一", "二", "三", "四", "五",
               "六", "七", "八", "九", "十", "十一", "十二"]
    cDayTens = ["", "十", "二十", "三十"]
    cDayDigits = ["", "一", "二", "三", "四", "五", "六", "七", "八", "九"]

    dten = floor(day / 10)
    ddigit = day % 10

    ydigit = year % 10
    yten = floor(year / 10) % 10
    yhundred = floor(year / 100) % 10
    ythousand = floor(year / 1000)

    weekday = datetime(year, month, day).weekday()

    return {
        "year": cYears[ythousand] + cYears[yhundred] + cYears[yten] + cYears[ydigit],
        "month": cMonths[month],
        "day": cDayTens[dten] + cDayDigits[ddigit],
        "weekday": cWeekDays[weekday],
    }
