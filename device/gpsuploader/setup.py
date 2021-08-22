from setuptools import setup
from setuptools import find_packages

__version__ = "0.0.1"

setup(
    name="gpsuploader",
    version=__version__,
    description="test package",
    long_description="""
    Get GPS Info on GPS Expansion module attached Omega2Plus.
    """,
    author="Shin Kurita",
    url="https://github.com/montblanc18/gps-viewer/gpsuploader",
    license="MIT",
    package_dir={"": "src", "resource": "resource"},
    packages=find_packages(where="src", exclude=("tests")),
)
