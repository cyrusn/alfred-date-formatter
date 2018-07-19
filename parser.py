import re
from datetime import datetime, date


class Parser():
    """docstring for Parser"""

    def __init__(self, datestring):
        super(Parser, self).__init__()
        self.year = self.parseDate(datestring)['year']
        self.month = self.parseDate(datestring)['month']
        self.day = self.parseDate(datestring)['day']
        self.date = datetime(self.year, self.month, self.day)

    def parseDate(self, datestring):
        today = date.today()
        regs = [
            r"^(?P<year>[0-9]+)年(?P<month>[0-9]+)月(?P<day>[0-9]+)日$",
            r"^(?P<month>[0-9]+)月(?P<day>[0-9]+)日$",
            r"^(?P<day>[0-9]+)\/(?P<month>[0-9]+)\/(?P<year>[0-9]+)$",
            r"^(?P<day>[0-9]+)\/(?P<month>[0-9]+)$",
            r"^(?P<year>[0-9]+)[\.|\-](?P<month>[0-9]+)[\.|\-](?P<day>[0-9]+)$",
            r"^(?P<month>[0-9]+)[\.|\-](?P<day>[0-9]+)$",
        ]
        for reg in regs:
            searchResult = re.search(reg, datestring)
            if searchResult:
                dateDict = {k: int(v)
                            for k, v in searchResult.groupdict().items()}
                if 'year' in dateDict:
                    return dateDict
                else:
                    dateDict.update({'year': today.year})
                    return dateDict

        return {
            "year": today.year,
            "month": today.month,
            "day": today.day,
        }
