# GPS UPLOADER

## Setup

```bash
$ python -m venv venv
$ pip install -r requirements.txt
```

## Development

```bath
## Start Development
$ source venv/bin/activate

## End Development
$ source venv/bin/deactivate
```


## Omega2+ GPS Expansion

```bash
$ ubus call gps info
{
    "age": 0,
    "latitude": "43.8616",
    "longitude": "-79.3854",
    "elevation": "184.0",
    "course": "",
    "speed": "N"
}
```

## Upload Data Format

based on omega2+ gps expansion format.

```bash
$ ubus call gps info
{
    "age": 0,
    "latitude": "43.8616",
    "longitude": "-79.3854",
    "elevation": "184.0",
    "course": "",
    "speed": "N"
    "recorded_at": "2021-08-06T12:34:56+09:00"
    "device_id": "012345"
}
```