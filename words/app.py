import random

from flask import Flask
import nltk
from nltk.corpus import wordnet as wn

app = Flask(__name__)
WORDS = []


@app.route('/')
def word_of_the_day():
    return "The word of the day is: {}".format(pick_word())


def pick_word():
    return random.choice(WORDS)


def setup():
    try:
        all_words = wn.words()
        print("have wordnet, all is well")
    except LookupError:
        print("don't have wordnet, downloading")
        nltk.download('wordnet')
        all_words = wn.words()
    global WORDS
    WORDS = [w for w in all_words if w[0].isalpha() and "_" not in w]


if __name__ == '__main__':
    setup()
    app.run(debug=True, host='0.0.0.0')
