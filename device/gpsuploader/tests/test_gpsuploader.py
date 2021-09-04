import copy
import json
import os
import time
import subprocess

import gpsuploader.gpsuploader as gpsu


def call_command(commands):
    cmd = []
    for command in commands:
        cmd.append(command)
    subprocess.call(cmd)
    output = subprocess.Popen(cmd, stdout=subprocess.PIPE)
    return output.communicate()[0]


def test_getgps():
    json_data = r"{\"age\":0, \"latitude\":\"43.8616\", \"longitude\":\"-79.3854\", \"elevation\":\"184.0\", \"course\":\"\", \"speed\":\"N\"}"
    want_json_data = '{"age":0, "latitude":"43.8616", "longitude":"-79.3854", "elevation":"184.0", "course":"", "speed":"N"}'
    cmd = "echo {0}".format(json_data)
    assert gpsu.get_gps(cmd) == json.loads(want_json_data)


def test_add_record():
    key = "add_key"
    value = "add_value"
    d = {
        "age": 0,
        "latitude": "43.8616",
        "longitude": "-79.3854",
        "elevation": "184.0",
        "course": "",
        "speed": "N",
    }
    want_d = copy.copy(d)
    want_d[key] = value
    assert gpsu.add_record(d, key, value) == want_d


def test_check_signal():
    tests = [
        {
            "case": "signal_true",
            "input": {
                "age": 0,
                "latitude": "43.8616",
                "longitude": "-79.3854",
                "elevation": "184.0",
                "course": "",
                "speed": "N",
            },
            "want": {
                "age": 0,
                "latitude": "43.8616",
                "longitude": "-79.3854",
                "elevation": "184.0",
                "course": "",
                "speed": "N",
                "signal": "true",
            },
        },
        {
            "case": "signal_false",
            "input": {
                "signal": "false",
            },
            "want": {
                "signal": "false",
            },
        },
    ]
    for t in tests:
        print()
        print("case: {0}".format(t["case"]))
        assert gpsu.check_signal(t["input"]) == t["want"]


def test_put_item():
    print()
    stream_name = "local_gps_omega2plus"
    endpoint_url = "http://localhost:4566"
    aws_region = "ap-northeast-1"
    aws_access_key_id = "dummy"
    aws_secret_access_key = "dummy"
    setup_cmds = [
        "aws kinesis --profile local create-stream --stream-name {0} --shard-count 1 --endpoint-url {1}".format(
            stream_name, endpoint_url
        )
    ]
    for cmd in setup_cmds:
        print(cmd)
        os.system(cmd)
    # wait for localstack kinesis ACTIVE
    wait_cmd = "aws kinesis --profile local describe-stream --stream-name {0} --endpoint-url {1}".format(
        stream_name, endpoint_url
    )

    for i in range(10):
        print(wait_cmd)
        jsons = call_command(wait_cmd.split(" "))
        s = json.loads(jsons)
        if "ACTIVE" == s["StreamDescription"]["StreamStatus"]:
            break
        time.sleep(1)

    gpsu.put_item(
        stream_name,
        "test_data",
        "123",
        aws_region,
        aws_access_key_id,
        aws_secret_access_key,
        endpoint_url,
    )

    end_cmds = [
        "aws kinesis --profile local delete-stream --stream-name {0} --endpoint-url {1}".format(
            stream_name, endpoint_url, aws_region
        )
    ]
    for cmd in end_cmds:
        print(cmd)
        os.system(cmd)
