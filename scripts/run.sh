#!/usr/bin
/bin/repopulate
/bin/rat
/bin/fs-snapshot snap > snapTwo.json
/bin/python /bin/compareSnapshots.py snapOne.json snapTwo.json
