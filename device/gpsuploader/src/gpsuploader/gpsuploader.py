import boto3
import copy
import datetime
import json
import os
import pytz
import subprocess


def get_gps(gps_cmd: str):
    gps_info = subprocess.check_output(
        gps_cmd, stdin=subprocess.DEVNULL, stderr=subprocess.DEVNULL, shell=True
    )
    t = gps_info.strip().decode("utf-8")
    # return json.loads(json.dumps(t))
    return json.loads(t)


def add_record(gps_data: dict, add_key: str, add_value: str):
    retval = copy.copy(gps_data)
    retval[add_key] = add_value
    return retval


def check_signal(gps_data: dict):
    retval = copy.copy(gps_data)
    try:
        if retval["signal"] == "false":
            pass
    except (KeyError, TypeError):
        retval["signal"] = "true"
    return retval


def put_item(
    stream_name: str,
    gps_data,
    partition_key,
    aws_region,
    aws_access_key_id,
    aws_secret_access_key,
    endpoint_url="",
):
    client = boto3.Session(
        aws_access_key_id=aws_access_key_id,
        aws_secret_access_key=aws_secret_access_key,
        region_name=aws_region,
    ).client("kinesis")
    if endpoint_url != "":
        client = boto3.Session(
            region_name=aws_region,
            aws_access_key_id="dummy",
            aws_secret_access_key="dummy",
        ).client("kinesis", endpoint_url=endpoint_url)
    client.put_record(
        StreamName=stream_name, Data=json.dumps(gps_data), PartitionKey=partition_key
    )


def handler():
    aws_region = os.environ["AWS_REGION"]
    if aws_region == "":
        raise ValueError

    device_id = os.environ["DEVICE_ID"]
    if device_id == "":
        raise ValueError

    stream_name = os.environ["STREAM_NAME"]
    if stream_name == "":
        raise ValueError

    partition_key = os.environ["PARTITION_KEY"]
    if partition_key == "":
        raise ValueError

    gps_cmd = os.environ["GPS_CMD"]
    if gps_cmd == "":
        raise ValueError

    aws_access_key_id = os.environ["AWS_ACCESS_KEY_ID"]
    if aws_access_key_id == "":
        raise ValueError

    aws_secret_access_key = os.environ["AWS_SECRET_ACCESS_KEY"]
    if aws_secret_access_key == "":
        raise ValueError

    endpoint_url = os.environ["ENDPOINT_URL"]

    gps_data = get_gps(gps_cmd)
    s = check_signal(gps_data)
    t = add_record(
        s,
        "recorded_at",
        datetime.datetime.now().astimezone(pytz.timezone("Asia/Tokyo")).isoformat(),
    )
    tt = add_record(t, "device_id", device_id)
    put_item(
        stream_name,
        tt,
        partition_key,
        aws_region,
        aws_access_key_id,
        aws_secret_access_key,
        endpoint_url,
    )


if __name__ == "__main__":
    handler()
