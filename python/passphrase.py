#!/usr/bin/env python3

import argparse
from random import randint


def generate_passphrase(length):
    with open('/usr/share/dict/words', 'r') as f:
        wordlist = [w.strip() for w in f.readlines() if "'" not in w]
        choices = " ".join([wordlist[randint(1, len(wordlist))] for _ in range(length)])
        print(choices)


if __name__ == '__main__':
    argparser = argparse.ArgumentParser("Generate a passphrase")
    argparser.add_argument('length', nargs='?', type=int, default=4,
                           help="How many words to include")
    args = argparser.parse_args()
    generate_passphrase(args.length)
