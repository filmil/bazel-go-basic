# SPDX-License-Identifier: Apache-2.0
"""A dummy bzl file to demonstrate stardoc and bzl_library."""

def hello(name):
    """Greets the user.

    Args:
        name: The name of the user to greet.
    """
    print("Hello, " + name)
