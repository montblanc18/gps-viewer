from boto import kinesis
import json
import os
import subprocess


def get_gps(gps_cmd: str):
    gps_info = subprocess.check_output(
        gps_cmd, stdin=subprocess.DEVNULL, stderr=subprocess.DEVNULL, shell=True
    )
    t = gps_info.strip().decode("utf-8")
    return json.loads(json.dumps(t))


def put_item(stream_name: str, gps_data, partition_key, aws_region):
    conn = kinesis.connect_to_regison(region_name=aws_region)
    conn.put_record(stream_name, data=gps_data, partition_key=partition_key)


def main():
    aws_region = os.environ["AWS_REGION"]
    if aws_region == "":
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

    gps_data = get_gps(gps_cmd)
    put_item(stream_name, gps_data, partition_key, aws_region)


if __name__ == "__main__":
    main()
