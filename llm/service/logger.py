import logging
from logging import Logger


def get_logger(name: str) -> Logger:
    fmt = "%(asctime)s [%(levelname)s] [%(name)s] %(message)s"
    handler = logging.StreamHandler()
    handler.setFormatter(logging.Formatter(fmt))
    logger = logging.getLogger(name)
    if not logger.handlers:
        logger.addHandler(handler)
    logger.setLevel(logging.DEBUG)
    logger.propagate = False
    return logger
