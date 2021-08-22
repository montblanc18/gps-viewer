import json

import gpsuploader.gpsuploader as gpsu


def test_gpsuploader():
    json_data = r"{\"age\":0, \"latitude\":\"43.8616\", \"longitude\":\"-79.3854\", \"elevation\":\"184.0\", \"course\":\"\", \"speed\":\"N\"}"
    want_json_data = '{"age":0, "latitude":"43.8616", "longitude":"-79.3854", "elevation":"184.0", "course":"", "speed":"N"}'
    cmd = "echo {0}".format(json_data)
    assert gpsu.get_gps(cmd) == json.loads(json.dumps(want_json_data))
